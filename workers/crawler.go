package workers

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"

	ether "github.com/ethereum/go-ethereum"
	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/kv-base-hack/base_crawler/common"
	"github.com/kv-base-hack/base_crawler/lib/blockchain"
	"github.com/kv-base-hack/base_crawler/lib/tokenrate"
	"github.com/kv-base-hack/base_crawler/lib/topic"
	"github.com/kv-base-hack/base_crawler/storage"
	"github.com/kv-base-hack/base_crawler/storage/db"
	"github.com/kv-base-hack/common/utils"
)

var maxDiff = big.NewFloat(1.3)
var minDiff = big.NewFloat(0.7)

// EvmCrawler gets trade logs, adding the information about USD equivalent on each trade.
type EvmCrawler struct {
	chain             common.Chain
	enable            bool
	log               *zap.SugaredLogger
	fromBlock         int64
	db                db.DB
	storage           *storage.Storage
	tokenFormatter    blockchain.TokenAmountFormatterInterface
	maxWorkers        int
	maxBlockForWorker int64
	timeout           time.Duration
	blockTimeResolver blockchain.BlockTimeResolverInterface
	tokenRate         tokenrate.TokenUsdtRateProvider
	topic             topic.Topic
	nodePool          *blockchain.NodePool
	tradeTable        string
	transferTable     string
}

// NewEvmCrawler create a new Crawler instance.
func NewEvmCrawler(log *zap.SugaredLogger,
	chain common.Chain,
	enable bool,
	storage *storage.Storage,
	fromBlock int64,
	db db.DB,
	tokenFormatter blockchain.TokenAmountFormatterInterface,
	maxWorkers int,
	maxBlockForWorker int64,
	timeout time.Duration,
	blockTimeResolver blockchain.BlockTimeResolverInterface,
	tokenRate tokenrate.TokenUsdtRateProvider,
	topic topic.Topic,
	nodePool *blockchain.NodePool,
	tradeTable string,
	transferTable string,
) *EvmCrawler {
	return &EvmCrawler{
		log:               log.With("chain", chain.String()),
		chain:             chain,
		fromBlock:         fromBlock,
		db:                db,
		tokenFormatter:    tokenFormatter,
		storage:           storage,
		maxWorkers:        maxWorkers,
		maxBlockForWorker: maxBlockForWorker,
		timeout:           timeout,
		blockTimeResolver: blockTimeResolver,
		tokenRate:         tokenRate,
		topic:             topic,
		nodePool:          nodePool,
		tradeTable:        tradeTable,
		transferTable:     transferTable,
		enable:            enable,
	}
}

func (c *EvmCrawler) Init() error {
	if !c.enable {
		return nil
	}
	log := c.log.With("runID", utils.RandomString(29))

	lastStoredBlock, err := c.db.GetLastStoredBlock(c.tradeTable)
	if err != nil {
		log.Errorw("error when get last stored from db", "err", err)
	}
	c.storage.SetLastStoredBlock(c.chain, lastStoredBlock)
	fromBlock := c.fromBlock
	if fromBlock < lastStoredBlock+1 {
		fromBlock = lastStoredBlock + 1
	}
	start := time.Now()
	latestBlock := c.storage.GetLatestBlock(c.chain)
	// get fromBlock to latest block first
	log.Infow("init trade", "fromBlock", fromBlock, "latestBlock", latestBlock, "lastStoredBlock", lastStoredBlock)
	err = c.getAndAddToStorageMissingTrades(log, fromBlock, latestBlock, c.timeout)
	if err != nil {
		return err
	}
	log.Infow("finish missing trade(ms)", "duration", time.Since(start).Milliseconds())

	return nil
}

func (c *EvmCrawler) Run() error {
	if !c.enable {
		return nil
	}
	c.log.Infow("start run periodically to get newer block")
	ticker := time.NewTicker(time.Second * 30)
	for ; ; <-ticker.C {
		log := c.log.With("runID", utils.RandomString(29))
		lastStoredBlock := c.storage.GetLastStoredBlock(c.chain)
		latestBlock := c.storage.GetLatestBlock(c.chain)
		if lastStoredBlock+1 < c.fromBlock {
			lastStoredBlock = c.fromBlock
		}
		log.Debugw("check periodically", "lastStoredBlock", lastStoredBlock, "latestBlock", latestBlock)
		if latestBlock > lastStoredBlock {
			err := c.getAndAddToStorageMissingTrades(log, lastStoredBlock+1, latestBlock, c.timeout)
			if err != nil {
				return err
			}
			log.Infow("update periodically", "lastStoredBlock", lastStoredBlock, "latestBlock", latestBlock)
		}
	}

	return nil
}

func (c *EvmCrawler) addLogsToStorage(trades []common.Tradelog, transfer []common.Transferlog) error {
	sort.Slice(trades, func(i, j int) bool {
		return (trades[i].BlockNumber < trades[j].BlockNumber) ||
			(trades[i].BlockNumber == trades[j].BlockNumber && trades[i].TxIndex < trades[j].TxIndex)
	})

	sort.Slice(transfer, func(i, j int) bool {
		return (transfer[i].BlockNumber < transfer[j].BlockNumber) ||
			(transfer[i].BlockNumber == transfer[j].BlockNumber && transfer[i].TxIndex < transfer[j].TxIndex)
	})

	c.log.Debugw("sort logs", "trades", trades, "transfer", transfer)

	err := c.db.SaveLogs(trades, transfer, c.tradeTable, c.transferTable)
	if err != nil {
		return err
	}

	var lastBlockNumber int64
	if len(trades) > 0 {
		lastBlockNumber = int64(trades[len(trades)-1].BlockNumber)
	} else {
		if len(transfer) > 0 {
			lastBlockNumber = int64(transfer[len(transfer)-1].BlockNumber)
		}
	}
	if lastBlockNumber == 0 {
		c.log.Errorw("error when set value for lastBlockNumber", "lastBlockNumber", lastBlockNumber)
		return nil
	}
	c.storage.SetLastStoredBlock(c.chain, lastBlockNumber)
	c.log.Infow("set last stored block", "block_number", lastBlockNumber)

	return nil
}

func (c *EvmCrawler) getAndAddToStorageMissingTrades(log *zap.SugaredLogger, fromBlock, toBlock int64, timeout time.Duration) error {
	if fromBlock > toBlock {
		return nil
	}
	var g errgroup.Group
	startBlock := fromBlock
	var mutex sync.Mutex
	for {
		finished := false
		newTradeLogs := []common.Tradelog{}
		newTransferLogs := []common.Transferlog{}
		for i := 0; i < c.maxWorkers; i++ {
			endBlock := startBlock + c.maxBlockForWorker
			if endBlock >= toBlock {
				endBlock = toBlock
				finished = true
			}
			log.Infow("worker running",
				"i", i,
				"startBlock", startBlock,
				"endBlock", endBlock,
				"toBlock", toBlock,
				"finished", finished)
			func(start, end int64) {
				g.Go(func() error {
					err := c.blockTimeResolver.Sync(c.chain, uint64(start), uint64(end))
					if err != nil {
						log.Errorw("can not sync block timestamp", "start", start, "end", end, "err", err)
						return err
					}
					trades, transfer, err := c.fetchLogs(log, big.NewInt(start), big.NewInt(end), timeout)
					if err != nil {
						log.Errorw("error when fetch logs", "start", start, "end", end, "err", err)
						return err
					}
					mutex.Lock()
					defer mutex.Unlock()
					newTradeLogs = append(newTradeLogs, trades...)
					newTransferLogs = append(newTransferLogs, transfer...)
					return nil
				})
			}(startBlock, endBlock)

			startBlock = endBlock + 1
			if startBlock > toBlock {
				finished = true
			}
			if finished {
				break
			}
		}
		err := g.Wait()
		if err != nil {
			log.Errorw("error when run concurrent", "err", err)
			return err
		}

		err = c.addLogsToStorage(newTradeLogs, newTransferLogs)
		if err != nil {
			log.Errorw("error when add trade and transfer logs to storage", "err", err)
			return err
		}

		if finished {
			break
		}
	}
	return nil
}

func (c *EvmCrawler) fetchLogsWithTopics(fromBlock, toBlock *big.Int, timeout time.Duration, topics [][]ethereum.Hash) ([]types.Log, error) {
	query := ether.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topics:    topics,
	}

	return c.nodePool.FilterLogs(c.chain, query)
}

func (c *EvmCrawler) fetchLogs(log *zap.SugaredLogger, fromBlock, toBlock *big.Int, timeout time.Duration) ([]common.Tradelog, []common.Transferlog, error) {
	topics := c.topic.GetTopics()

	typeLogs, err := c.fetchLogsWithTopics(fromBlock, toBlock, timeout, topics)
	if err != nil {
		log.Errorw("error when fetch logs by topics", "fromBlock", fromBlock, "toBlock", toBlock, "err", err)
		return nil, nil, errors.Wrap(err, "failed to fetch log by topic")
	}
	return c.assembleTradeLogs(typeLogs, timeout)
}

// we check equal value first
// sometime transferValue changed, so we should check amountIn in range of transferValue, not exceed +-0.3 of transferValue
func (c *EvmCrawler) isEqual(transferValue *big.Int, amountIn *big.Int) bool {
	// check is transferValue and amountIn
	if transferValue.Cmp(amountIn) == 0 {
		return true
	}

	transferFloat := big.NewFloat(0).SetInt(transferValue)
	maxTransferValue := big.NewFloat(0).Mul(transferFloat, maxDiff)
	minTransferValue := big.NewFloat(0).Mul(transferFloat, minDiff)
	amountInFloat := big.NewFloat(0).SetInt(amountIn)

	if amountInFloat.Cmp(maxTransferValue) < 0 && amountInFloat.Cmp(minTransferValue) > 0 {
		return true
	}
	return false
}

func (c *EvmCrawler) handleSwapTopic(txHash ethereum.Hash, eventLogs []types.Log, timeout time.Duration) (common.Tradelog, error) {
	if len(eventLogs) <= 0 {
		return common.Tradelog{}, fmt.Errorf("invalid event logs for swap topic 1")
	}

	log := c.log.With("txhash", txHash)
	now := time.Now()
	log.Debugw("handle swap topic", "eventLogs", eventLogs)

	txReceipt, txSender, err := c.nodePool.GetTxData(c.chain, txHash, c.timeout)
	if err != nil {
		log.Errorw("error when get txdata", "err", err)
		return common.Tradelog{}, err
	}
	blockTimestamp, err := c.blockTimeResolver.Resolve(c.chain, eventLogs[0].BlockNumber)
	if err != nil {
		log.Errorw("error when resolve block timestamp", "err", err)
		return common.Tradelog{}, err
	}

	tokenUsdtRate := map[string]float64{}
	var nativeTokenUsdtRate float64
	ethUsdtRate, err := c.tokenRate.USDRate(common.ChainBase, tokenrate.EthPair, blockTimestamp)
	if err != nil {
		log.Errorw("error when get token rate", "err", err)
		return common.Tradelog{}, err
	}
	nativeTokenUsdtRate = ethUsdtRate
	tokenUsdtRate = map[string]float64{
		blockchain.BaseWEthAddr: ethUsdtRate,
		blockchain.BaseUsdAddr:  1,
		blockchain.BaseUsdcAddr: 1,
		blockchain.BaseDaiAddr:  1,
	}

	setTokenUsdtRate := func(amountIn, amountOut *big.Int, tokenInAddress, tokenOutAddress ethereum.Address) error {
		inAmount, err := c.tokenFormatter.FromWei(c.chain, tokenInAddress, amountIn)
		if err != nil {
			log.Errorw("error when fromWei token in", "err", err)
			return err
		}
		outAmount, err := c.tokenFormatter.FromWei(c.chain, tokenOutAddress, amountOut)
		if err != nil {
			log.Errorw("error when fromWei token out", "err", err)
			return err
		}

		inUsdtRate, existIn := tokenUsdtRate[strings.ToLower(tokenInAddress.String())]
		outUsdtRate, existOut := tokenUsdtRate[strings.ToLower(tokenOutAddress.String())]
		if existIn && !existOut {
			tokenUsdtRate[strings.ToLower(tokenOutAddress.String())] = inUsdtRate * inAmount / outAmount
		} else if !existIn && existOut {
			tokenUsdtRate[strings.ToLower(tokenInAddress.String())] = outUsdtRate * outAmount / inAmount
		}
		return nil
	}

	transferLog := []common.TransferLogData{}
	for i, l := range txReceipt.Logs {
		if l.Removed || c.topic.IsSwapTopic(*l) {
			continue
		}
		transfer, err := c.topic.GetTransferLog(*txReceipt.Logs[i])
		if err != nil {
			continue
		}
		log.Debugw("transfer log",
			"transfer.To", transfer.To,
			"transfer.Value", transfer.Value)
		transferLog = append(transferLog, transfer)
	}

	routes := []common.Routes{}
	var txIndex uint
	for _, l := range txReceipt.Logs {
		if l.Removed || !c.topic.IsSwapTopic(*l) {
			continue
		}
		txIndex = l.TxIndex
		swapLogData, err := c.topic.GetSwapLogInfo(*l)
		log.Infow("swap data", "index", l.Index, "swapLogData", swapLogData, "err", err)
		if err != nil {
			log.Errorw("error when swap data", "err", err)
			return common.Tradelog{}, err
		}

		tokenInAmount := swapLogData.AmountIn
		tokenOutAmount := swapLogData.AmountOut
		var tokenInAddress, tokenOutAddress ethereum.Address
		foundTokenIn := false
		foundTokenOut := false

		if c.topic.IsFilledData(*l) {
			tokenInAddress = swapLogData.TokenIn
			tokenOutAddress = swapLogData.TokenOut
			foundTokenIn = true
			foundTokenOut = true
		} else {
			address := swapLogData.Address
			recipientAddress := swapLogData.RecipientAddress

			// compare extractly
			for _, transfer := range transferLog {
				addressCmp := address.Cmp(transfer.To) == 0
				amountInCmp := transfer.Value.Cmp(swapLogData.AmountIn) == 0
				recipientCmp := recipientAddress.Cmp(transfer.To) == 0
				amountOutCmp := transfer.Value.Cmp(swapLogData.AmountOut) == 0
				// log.Debugw("compare_exactly_log",
				// 	"foundTokenIn", foundTokenIn,
				// 	"address", address,
				// 	"transfer.To", transfer.To,
				// 	"transfer.Value", transfer.Value,
				// 	"swapLogData.AmountIn", swapLogData.AmountIn,
				// 	"foundTokenOut", foundTokenOut,
				// 	"recipientAddress", recipientAddress,
				// 	"swapLogData.AmountOut", swapLogData.AmountOut,
				// 	"equal_address_in", addressCmp,
				// 	"equal_value_in", amountInCmp,
				// 	"equal_address_in", recipientCmp,
				// 	"equal_value_out", amountOutCmp,
				// )
				if !foundTokenIn && addressCmp && amountInCmp {
					tokenInAddress = transfer.Raw.Address
					foundTokenIn = true
				} else if !foundTokenOut && recipientCmp && amountOutCmp {
					tokenOutAddress = transfer.Raw.Address
					foundTokenOut = true
				}
			}
			log.Debugw("after_compare_exactly_log",
				"foundTokenIn", foundTokenIn,
				"tokenInAddress", tokenInAddress,
				"foundTokenOut", foundTokenOut,
				"tokenOutAddress", tokenOutAddress,
			)

			// compare base on percent
			for _, transfer := range transferLog {
				addressCmp := address.Cmp(transfer.To) == 0
				amountInCmp := c.isEqual(transfer.Value, swapLogData.AmountIn)
				recipientCmp := recipientAddress.Cmp(transfer.To) == 0
				amountOutCmp := c.isEqual(transfer.Value, swapLogData.AmountOut)
				log.Debugw("compare_percent_log",
					"foundTokenIn", foundTokenIn,
					"address", address,
					"transfer.To", transfer.To,
					"transfer.Value", transfer.Value,
					"swapLogData.AmountIn", swapLogData.AmountIn,
					"foundTokenOut", foundTokenOut,
					"recipientAddress", recipientAddress,
					"swapLogData.AmountOut", swapLogData.AmountOut,
					"equal_address_in", addressCmp,
					"equal_value_in", amountInCmp,
					"equal_address_in", recipientCmp,
					"equal_value_out", amountOutCmp,
				)
				if !foundTokenIn && addressCmp && amountInCmp {
					tokenInAddress = transfer.Raw.Address
					foundTokenIn = true
				} else if !foundTokenOut && recipientCmp && amountOutCmp {
					tokenOutAddress = transfer.Raw.Address
					foundTokenOut = true
				}
			}
			log.Debugw("after_compare_percent_log",
				"foundTokenIn", foundTokenIn,
				"tokenInAddress", tokenInAddress,
				"foundTokenOut", foundTokenOut,
				"tokenOutAddress", tokenOutAddress,
			)
		}

		if foundTokenIn && foundTokenOut {
			setTokenUsdtRate(swapLogData.AmountIn, swapLogData.AmountOut, tokenInAddress, tokenOutAddress)
			tokenInAmountFloat, err := c.tokenFormatter.FromWei(c.chain, tokenInAddress, tokenInAmount)
			if err != nil {
				log.Errorw("can not convert token in amount",
					"tokenInAddress", tokenInAddress, "tokenInAmount", tokenInAmount, "err", err)
				return common.Tradelog{}, err
			}

			tokenOutAmountFloat, err := c.tokenFormatter.FromWei(c.chain, tokenOutAddress, tokenOutAmount)
			if err != nil {
				log.Errorw("can not convert token out amount",
					"tokenOutAddress", tokenOutAddress, "tokenOutAmount", tokenOutAmount, "err", err)
				return common.Tradelog{}, err
			}
			routes = append(routes, common.Routes{
				LogIndex: l.Index,

				TokenInAddress: tokenInAddress,
				TokenInAmount:  tokenInAmountFloat,

				TokenOutAddress: tokenOutAddress,
				TokenOutAmount:  tokenOutAmountFloat,
				Exchange:        swapLogData.Address,
			})

			// re-calculate-token/usdt rate before
			for j := len(routes) - 1; j >= 0; j-- {
				log := routes[j]
				setTokenUsdtRate(log.AmountIn, log.AmountOut, log.TokenInAddress, log.TokenOutAddress)
			}
		} else {
			return common.Tradelog{}, fmt.Errorf("not found token in or token out for tx: %s at index %d", txHash, l.Index)
		}
	}
	for i, t := range routes {
		if v, exist := tokenUsdtRate[strings.ToLower(t.TokenInAddress.String())]; !exist {
			log.Errorw("not found tokenIn", "t.TokenInAddress", t.TokenInAddress, "tokenUsdtRate", tokenUsdtRate)
			return common.Tradelog{}, fmt.Errorf("not found tokenIn(%s) in tokenUsdtRate for tx: %s at index %d",
				t.TokenInAddress, txHash, t.LogIndex)
		} else {
			routes[i].TokenInUsdtRate = v
		}

		if v, exist := tokenUsdtRate[strings.ToLower(t.TokenOutAddress.String())]; !exist {
			log.Errorw("not found tokenOut", "t.TokenOutAddress", t.TokenOutAddress, "tokenUsdtRate", tokenUsdtRate)
			return common.Tradelog{}, fmt.Errorf("not found tokenOut(%s) in tokenUsdtRate for tx: %s at index %d", t.TokenOutAddress, txHash, t.LogIndex)
		} else {
			routes[i].TokenOutUsdtRate = v
		}
	}
	log.Debugw("end handle swap topic", "duration(ms)", time.Since(now).Milliseconds())
	if len(routes) == 0 {
		return common.Tradelog{}, nil
	}
	routesByte, err := json.Marshal(routes)
	if err != nil {
		return common.Tradelog{}, err
	}
	return common.Tradelog{
		BlockTimestamp: blockTimestamp,
		BlockNumber:    eventLogs[0].BlockNumber,
		TxIndex:        txIndex,
		TxHash:         txHash,
		Sender:         txSender,

		TokenInAddress:  routes[0].TokenInAddress,
		TokenInAmount:   routes[0].TokenInAmount,
		AmountIn:        routes[0].AmountIn,
		TokenInUsdtRate: routes[0].TokenInUsdtRate,

		TokenOutAddress:     routes[len(routes)-1].TokenOutAddress,
		TokenOutAmount:      routes[len(routes)-1].TokenOutAmount,
		AmountOut:           routes[len(routes)-1].AmountOut,
		TokenOutUsdtRate:    routes[len(routes)-1].TokenOutUsdtRate,
		Routes:              routesByte,
		NativeTokenUsdtRate: nativeTokenUsdtRate,
	}, nil
}

func (c *EvmCrawler) assembleTradeLogs(eventLogs []types.Log,
	timeout time.Duration) ([]common.Tradelog, []common.Transferlog, error) {
	var (
		trades   []common.Tradelog
		transfer []common.Transferlog
	)
	now := time.Now()

	tx := map[ethereum.Hash][]types.Log{}
	// event logs can be a splitted part of tx hash, so we combine them by txhash
	for _, log := range eventLogs {
		if log.Removed {
			continue // Removed due to chain reorg
		}
		if len(log.Topics) == 0 {
			c.log.Debugw("log item has no topic")
			continue
		}

		// check it is cex transfer log or not
		if c.topic.IsTransferTopic(log) {
			transferLog, err := c.topic.GetCexTransferLog(log)
			if err == nil {
				// is cex transfer log
				blockTimestamp, err := c.blockTimeResolver.Resolve(c.chain, log.BlockNumber)
				if err != nil {
					c.log.Errorw("error when resolve block timestamp", "err", err)
					return []common.Tradelog{}, []common.Transferlog{}, err
				}

				amount, err := c.tokenFormatter.FromWei(c.chain, log.Address, transferLog.Value)
				if err != nil {
					c.log.Errorw("error when get normal value", "value", transferLog.Value, "address", log.Address, "err", err)
					continue
				}

				exchange := transferLog.From
				if transferLog.IsCexIn {
					exchange = transferLog.To
				}
				ethTransferlog := common.Transferlog{
					BlockTimestamp: blockTimestamp,
					BlockNumber:    log.BlockNumber,
					TxIndex:        log.TxIndex,
					TxHash:         log.TxHash,

					FromAddress: transferLog.From,
					ToAddress:   transferLog.To,

					TokenAddress: log.Address,
					TokenAmount:  amount,
					IsCexIn:      transferLog.IsCexIn,
					Exchange:     exchange,
				}

				transfer = append(transfer, ethTransferlog)
				c.log.Infow("is cex transfer log", "transferLog", transferLog, "addr", log.Address, "ethTransferlog", ethTransferlog)

				continue
			}
		}
		tx[log.TxHash] = append(tx[log.TxHash], log)
	}
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for hash, v := range tx {
		wg.Add(1)
		go func(txHash ethereum.Hash, logs []types.Log) {
			defer wg.Done()
			tradeLog, err := c.handleSwapTopic(txHash, logs, timeout)
			if err != nil {
				c.log.Errorw("error when handle swap topic", "txHash", txHash, "err", err)
				return
			}
			mutex.Lock()
			if tradeLog.TokenInUsdtRate > 0 && tradeLog.TokenOutUsdtRate > 0 {
				trades = append(trades, tradeLog)
			}
			mutex.Unlock()
		}(hash, v)
	}
	wg.Wait()
	c.log.Debugw("assembleTradeLogs", "duration(ms)", time.Since(now).Milliseconds(), "trades", trades)

	return trades, transfer, nil
}
