package db

import "github.com/kv-base-hack/base_crawler/common"

type DB interface {
	GetLastStoredBlock(table string) (int64, error)
	GetSmallestStoredBlock(table string) (int64, error)
	SaveLogs(tradeLogs []common.Tradelog, transferLogs []common.Transferlog, tradeTable, transferTable string) error
}
