package bsc

import (
	"fmt"
	"math/big"
	"strings"

	eth "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/kv-base-hack/base_crawler/common"
	"github.com/kv-base-hack/base_crawler/lib/contracts"
	"github.com/kv-base-hack/base_crawler/lib/mathutil"
	"go.uber.org/zap"
)

var (
	cexWallets = []string{}
)

const (
	swapTopic1 = "0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"
	swapTopic2 = "0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"
	swapTopic6 = "0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83" // https://bscscan.com/tx/0xb70a19f26054170911ef160ae0c1857032ba6da46592c4733862cb16d5d0b361#eventlog

	transfer1 = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

// transfer2 = "0xe59fdd36d0d223c0c7d996db7ad796880f45e1936cb0bb7ac102e7082e031487" // erc 20 transfer
)

var (
	bigZero = big.NewInt(0)
)

type BaseTopic struct {
	log        *zap.SugaredLogger
	swapTopic1 contracts.SwapTopic1
	swapTopic2 contracts.SwapTopic2

	swapTopic6 contracts.SwapTopic6

	transfer1 contracts.Transfer1

	cexAddress map[string]bool
}

func NewBaseTopic(log *zap.SugaredLogger) (*BaseTopic, error) {
	swapTopic1, err := contracts.NewSwapTopic1(eth.Address{}, nil)
	if err != nil {
		return &BaseTopic{}, err
	}

	swapTopic2, err := contracts.NewSwapTopic2(eth.Address{}, nil)
	if err != nil {
		return &BaseTopic{}, err
	}
	// swapTopic3, err := contracts.NewSwapTopic3(eth.Address{}, nil)
	// if err != nil {
	// 	return &BaseTopic{}, err
	// }
	// swapTopic4, err := contracts.NewSwapTopic4(eth.Address{}, nil)
	// if err != nil {
	// 	return &BaseTopic{}, err
	// }
	// swapTopic5, err := contracts.NewSwapTopic5(eth.Address{}, nil)
	// if err != nil {
	// 	return &BaseTopic{}, err
	// }
	swapTopic6, err := contracts.NewSwapTopic6(eth.Address{}, nil)
	if err != nil {
		return &BaseTopic{}, err
	}
	// swapTopic7, err := contracts.NewSwapTopic7(eth.Address{}, nil)
	// if err != nil {
	// 	return &BaseTopic{}, err
	// }

	transfer1, err := contracts.NewTransfer1(eth.Address{}, nil)
	if err != nil {
		return &BaseTopic{}, err
	}
	// transfer2, err := contracts.NewTransfer2(eth.Address{}, nil)
	// if err != nil {
	// 	return &BaseTopic{}, err
	// }

	cexAddress := map[string]bool{}
	for _, addr := range cexWallets {
		cexAddress[strings.ToLower(addr)] = true
	}

	return &BaseTopic{
		log:        log,
		swapTopic1: *swapTopic1,
		swapTopic2: *swapTopic2,
		// swapTopic3: *swapTopic3,
		// swapTopic4: *swapTopic4,
		// swapTopic5: *swapTopic5,
		swapTopic6: *swapTopic6,
		// swapTopic7: *swapTopic7,

		transfer1: *transfer1,
		// transfer2:  *transfer2,
		cexAddress: cexAddress,
	}, nil
}

func (e *BaseTopic) GetTopics() [][]eth.Hash {
	topics := [][]eth.Hash{
		{
			eth.HexToHash(swapTopic1),
			eth.HexToHash(swapTopic2),
			eth.HexToHash(swapTopic6),

			eth.HexToHash(transfer1),
		},
	}
	return topics
}

func (e *BaseTopic) IsSwapTopic(log types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0].Hex() == swapTopic1 ||
		log.Topics[0].Hex() == swapTopic2 || log.Topics[0].Hex() == swapTopic6
}

func (e *BaseTopic) IsTransferTopic(log types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0].Hex() == transfer1
}

func (e *BaseTopic) getSwapLogInfoFromTopic1(l types.Log) (common.SwapLogData, error) {
	log := e.log.With("func", "getSwapLogInfoFromTopic1")
	trade, err := e.swapTopic1.ParseSwap(l)
	if err != nil {
		return common.SwapLogData{}, err
	}

	log.Debugw("trade_log_topic1", "trade", trade)

	amountIn := mathutil.MinBigInt(trade.Amount0In, trade.Amount1In)
	if amountIn.Cmp(bigZero) == 0 {
		amountIn = mathutil.MaxBigInt(trade.Amount0In, trade.Amount1In)
	}
	amountOut := mathutil.MinBigInt(trade.Amount0Out, trade.Amount1Out)
	if amountOut.Cmp(bigZero) == 0 {
		amountOut = mathutil.MaxBigInt(trade.Amount0Out, trade.Amount1Out)
	}
	adress := trade.Raw.Address
	to := trade.To

	return common.SwapLogData{
		AmountIn:         amountIn,
		AmountOut:        amountOut,
		Address:          adress,
		RecipientAddress: to,
	}, nil
}

func (e *BaseTopic) GetSwapLogInfo(log types.Log) (common.SwapLogData, error) {
	switch log.Topics[0].Hex() {
	// case pancakeSwap:
	// 	return e.getSwapLogInfoForPancakeSwap(log)
	case swapTopic1:
		return e.getSwapLogInfoFromTopic1(log)
	case swapTopic2:
		return e.getSwapLog2(log)
	case swapTopic6:
		return e.getSwapLogInfoForSwapTopic6(log)
		// case swapTopic7:
		// 	return e.getSwapLogInfoForSwapTopic7(log)
	}
	return common.SwapLogData{}, fmt.Errorf("not handle swap log")
}

func (e *BaseTopic) getSwapLog2(l types.Log) (common.SwapLogData, error) {
	log := e.log.With("func", "getSwapLog2")

	amount0, amount1, address, to, err := e.getSwap2TradeInfo(l)
	if err != nil {
		return common.SwapLogData{}, nil
	}

	log.Debugw("trade_log_topic2",
		"amount0", amount0,
		"amount1", amount1,
		"address", address,
		"to", to)

	var amountIn, amountOut *big.Int
	if amount0.Cmp(bigZero) > 0 {
		amountIn = amount0
		amountOut = amount1.Neg(amount1)
	} else {
		amountIn = amount1
		amountOut = amount0.Neg(amount0)
	}

	return common.SwapLogData{
		AmountIn:         amountIn,
		AmountOut:        amountOut,
		Address:          address,
		RecipientAddress: to,
	}, nil
}

func (e *BaseTopic) getSwap2TradeInfo(l types.Log) (amount0, amount1 *big.Int, address, to eth.Address, err error) {
	trade, err := e.swapTopic2.ParseSwap(l)
	if err == nil {
		amount0 = trade.Amount0
		amount1 = trade.Amount1
		address = trade.Raw.Address
		to = trade.Recipient
		return
	}
	return
}

func (e *BaseTopic) IsFilledData(log types.Log) bool {
	// switch log.Topics[0].Hex() {
	// case swapTopic7:
	// 	return true
	// }
	return false
}

func (e *BaseTopic) getSwapLogInfoForSwapTopic6(l types.Log) (common.SwapLogData, error) {
	log := e.log.With("func", "getSwapLogInfoForSwapTopic6")
	trade, err := e.swapTopic6.ParseSwap(l)
	if err != nil {
		return common.SwapLogData{}, err
	}

	log.Debugw("getSwapLogInfoForSwapTopic6", "trade", trade)

	log.Debugw("trade_log_topic2",
		"amount0", trade.Amount0,
		"amount1", trade.Amount1,
		"address", trade.Raw.Address,
		"to", trade.Recipient)

	var amountIn, amountOut *big.Int
	if trade.Amount0.Cmp(bigZero) > 0 {
		amountIn = trade.Amount0
		amountOut = trade.Amount1.Neg(trade.Amount1)
	} else {
		amountIn = trade.Amount1
		amountOut = trade.Amount0.Neg(trade.Amount0)
	}

	return common.SwapLogData{
		AmountIn:         amountIn,
		AmountOut:        amountOut,
		Address:          trade.Raw.Address,
		RecipientAddress: trade.Recipient,
	}, nil
}

func (e *BaseTopic) GetTransferLog(log types.Log) (common.TransferLogData, error) {
	if len(log.Topics) == 0 {
		return common.TransferLogData{}, fmt.Errorf("not transfer log")
	}
	switch log.Topics[0].Hex() {
	case transfer1:
		return e.getTransferLog1(log)
	}
	return common.TransferLogData{}, fmt.Errorf("not transfer log")
}

func (e *BaseTopic) getTransferLog1(l types.Log) (common.TransferLogData, error) {
	log := e.log.With("func", "getTransferLog1")
	transfer, err := e.transfer1.ParseTransfer(l)
	if err != nil {
		return common.TransferLogData{}, err
	}

	log.Debugw("transfer_log1", "transfer", transfer)

	return common.TransferLogData{
		From:  transfer.From,
		To:    transfer.To,
		Value: transfer.Value,
		Raw:   transfer.Raw,
	}, nil
}

func (e *BaseTopic) GetCexTransferLog(log types.Log) (common.TransferLogData, error) {
	switch log.Topics[0].Hex() {
	case transfer1:
		return e.getCexTransferLog1(log)
		// case transfer2:
		// 	return e.getCexTransferLog2(log)
	}
	return common.TransferLogData{}, fmt.Errorf("not cex transfer log")
}

func (e *BaseTopic) getCexTransferLog1(l types.Log) (common.TransferLogData, error) {
	log := e.log.With("func", "getTransferLog1")
	transfer, err := e.transfer1.ParseTransfer(l)
	if err != nil {
		return common.TransferLogData{}, err
	}

	transferLog := common.TransferLogData{
		From:  transfer.From,
		To:    transfer.To,
		Value: transfer.Value,
		Raw:   transfer.Raw,
	}
	log.Debugw("log", "transferLog1", transferLog)
	isCexIn, isCexTransfer := e.isCexTransfer(transferLog)

	if !isCexTransfer {
		return common.TransferLogData{}, fmt.Errorf("not cex transfer")
	}
	transferLog.IsCexIn = isCexIn

	return transferLog, nil
}

// func (e *BaseTopic) getCexTransferLog2(l types.Log) (common.TransferLogData, error) {
// 	log := e.log.With("func", "getCexTransferLog2")
// 	transfer, err := e.transfer2.ParseERC20Transfer(l)
// 	if err != nil {
// 		return common.TransferLogData{}, err
// 	}

// 	transferLog := common.TransferLogData{
// 		From:  transfer.From,
// 		To:    transfer.To,
// 		Value: transfer.Amount,
// 		Raw:   transfer.Raw,
// 	}
// 	log.Debugw("log", "transferLog2", transferLog)
// 	isCexIn, isCexTransfer := e.isCexTransfer(transferLog)

// 	if !isCexTransfer {
// 		return common.TransferLogData{}, fmt.Errorf("not cex transfer")
// 	}
// 	transferLog.IsCexIn = isCexIn

// 	return transferLog, nil
// }

// return isCexIn, isCexTransfer or not
func (e *BaseTopic) isCexTransfer(log common.TransferLogData) (bool, bool) {
	if _, exist := e.cexAddress[strings.ToLower(log.From.String())]; exist {
		return false, true
	}
	if _, exist := e.cexAddress[strings.ToLower(log.To.String())]; exist {
		return true, true
	}
	return false, false
}
