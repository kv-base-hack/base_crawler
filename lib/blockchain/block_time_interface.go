package blockchain

import (
	"time"

	"github.com/kv-base-hack/base_crawler/common"
)

// BlockTimeResolverInterface define the functionality
type BlockTimeResolverInterface interface {
	Sync(chain common.Chain, fromBlock, toBlock uint64) error
	Resolve(chain common.Chain, blockNumber uint64) (time.Time, error)
}
