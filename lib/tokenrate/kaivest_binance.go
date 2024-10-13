package tokenrate

import (
	"strconv"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
	obc "github.com/kv-base-hack/base-client/kaivest_binance"
	"github.com/kv-base-hack/base_crawler/common"
	"go.uber.org/zap"
)

const name = "kaivest_binance"
const EthPair = "ETHUSDT"
const WbtcPair = "BTCUSDT"
const BnbPair = "BNBUSDT"
const interval = "1m"
const limit = 1000

type Network struct {
	klines             []binance.Kline
	fromBlockTimeStamp time.Time
	pair               string
}

type ChainWithPair struct {
	chain common.Chain
	pair  string
}

type KaivestBinance struct {
	log      *zap.SugaredLogger
	duration time.Duration

	kaivestBinanceClient *obc.KaivestBinanceClient
	mutex                sync.RWMutex
	chain                map[ChainWithPair]*Network
}

func NewKaivestBinance(log *zap.SugaredLogger, duration time.Duration,
	baseFromBlockTimeStamp time.Time,
	kaivestBinanceClient *obc.KaivestBinanceClient) *KaivestBinance {
	return &KaivestBinance{
		log:      log,
		duration: duration,
		chain: map[ChainWithPair]*Network{
			{
				chain: common.ChainBase,
				pair:  EthPair,
			}: {
				klines:             make([]binance.Kline, 0),
				fromBlockTimeStamp: baseFromBlockTimeStamp.Add(-time.Minute * 10), // get history of 10 ago from this timestamp
				pair:               EthPair,
			},
		},
		kaivestBinanceClient: kaivestBinanceClient,
	}
}

func (kb *KaivestBinance) Init() {
	now := time.Now().UnixMilli()

	for key, network := range kb.chain {
		start := network.fromBlockTimeStamp.UnixMilli()
		kb.InsertNewKline(key.chain, key.pair, start, now)
	}
}

func (kb *KaivestBinance) InsertNewKline(chain common.Chain, pair string, start, end int64) {
	kb.mutex.Lock()
	defer kb.mutex.Unlock()
	for {
		kline, err := kb.kaivestBinanceClient.GetSpotKLine(pair, interval, limit, start)
		if err != nil {
			panic(err)
		}
		if len(kline) == 0 {
			break
		}
		start = kline[len(kline)-1].CloseTime + 1
		chainWithPair := ChainWithPair{
			chain: chain,
			pair:  pair,
		}
		kb.chain[chainWithPair].klines = append(kb.chain[chainWithPair].klines, kline...)
		// kb.log.Infow("insert new klines", "chain", chain.String(), "kline", kline)
		if start > end {
			break
		}
	}
}

func (kb *KaivestBinance) Run() {
	// call to kai vest binance to get data
	ticker := time.NewTicker(kb.duration)
	for ; ; <-ticker.C {
		now := time.Now().Unix()
		for key, network := range kb.chain {
			lastInsert := network.klines[len(network.klines)-1].CloseTime
			kb.InsertNewKline(key.chain, key.pair, lastInsert+1, now)
		}
	}
}

// the historical price of native token of chain.
func (kb *KaivestBinance) USDRate(chain common.Chain, pair string, timestamp time.Time) (float64, error) {
	kb.mutex.RLock()
	defer kb.mutex.RUnlock()
	fi := 0
	chainWithPair := ChainWithPair{
		chain: chain,
		pair:  pair,
	}
	la := len(kb.chain[chainWithPair].klines) - 1
	tokenUsdtRate := 0.0
	for fi <= la {
		mid := (fi + la) / 2
		if kb.chain[chainWithPair].klines[mid].OpenTime <= timestamp.UnixMilli() {
			rate, err := strconv.ParseFloat(kb.chain[chainWithPair].klines[mid].Open, 64)
			if err != nil {
				return 0, err
			}
			tokenUsdtRate = rate
			fi = mid + 1
		} else {
			la = mid - 1
		}
	}
	return tokenUsdtRate, nil
}

// Name return name of CoinGecko provider name
func (kb *KaivestBinance) Name() string {
	return name
}
