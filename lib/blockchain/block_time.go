package blockchain

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum/core/types"
	cm "github.com/kv-base-hack/base_crawler/common"
)

const retry = 5

// BlockTimeResolver is a helper to get transaction timestamp from block number.
// It has a cache for one block.
type BlockTimeResolver struct {
	mu            sync.RWMutex
	log           *zap.SugaredLogger
	cachedHeaders map[cm.Chain]map[uint64]*types.Header
	nodePool      *NodePool
	timeout       time.Duration
}

// NewBlockTimeResolver returns BlockTimeResolver instance given a ethereum client.
func NewBlockTimeResolver(log *zap.SugaredLogger, nodePool *NodePool, timeout time.Duration) (*BlockTimeResolver, error) {
	return &BlockTimeResolver{
		log: log,
		cachedHeaders: map[cm.Chain]map[uint64]*types.Header{
			cm.ChainBase: map[uint64]*types.Header{},
		},
		nodePool: nodePool,
		timeout:  timeout,
	}, nil
}

func (btr *BlockTimeResolver) Sync(chain cm.Chain, fromBlock, toBlock uint64) error {
	var wg sync.WaitGroup
	for i := fromBlock; i <= toBlock; i++ {
		wg.Add(1)
		go func(blockNumber uint64) {
			defer wg.Done()
			for j := 0; j < retry; j++ {
				blockHeader, err := btr.nodePool.GetBlockTimeStamp(chain, timeout, blockNumber)
				if err != nil {
					btr.log.Errorw("error when get block header of block", "blockNumber", blockNumber, "retry", j)
					continue
				}
				btr.mu.Lock()
				btr.cachedHeaders[chain][blockNumber] = blockHeader
				btr.mu.Unlock()
				break
			}
		}(i)
	}
	wg.Wait()
	return nil
}

// Resolve returns timestamp from block number.
// It cachedHeaders block number and block header to reduces the number of request
// to node.
func (btr *BlockTimeResolver) Resolve(chain cm.Chain, blockNumber uint64) (time.Time, error) {
	// cache hit happy path
	btr.mu.RLock()
	defer btr.mu.RUnlock()
	header, ok := btr.cachedHeaders[chain][blockNumber]
	if ok {
		ts := time.Unix(int64(header.Time), 0).UTC()
		return ts, nil
	}

	return time.Now(), fmt.Errorf("not found block timestamp for block number: %d", blockNumber)
}
