package topic

import (
	eth "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/kv-base-hack/base_crawler/common"
)

type Topic interface {
	IsSwapTopic(log types.Log) bool
	GetTopics() [][]eth.Hash
	GetSwapLogInfo(log types.Log) (common.SwapLogData, error)

	IsTransferTopic(log types.Log) bool
	GetTransferLog(log types.Log) (common.TransferLogData, error)

	IsFilledData(log types.Log) bool
	GetCexTransferLog(log types.Log) (common.TransferLogData, error)
}
