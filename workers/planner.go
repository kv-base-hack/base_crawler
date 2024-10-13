package workers

import (
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/kv-base-hack/base_crawler/common"
	"github.com/kv-base-hack/base_crawler/lib/blockchain"
	"github.com/kv-base-hack/base_crawler/storage"
)

var ErrEOF = errors.New("no more planned blocks")

// get next block
type crawlPlanner struct {
	log           *zap.SugaredLogger
	chain         common.Chain
	enable        bool
	confirmations int64 // number of block confirmations before fetching
	inmemStorage  *storage.Storage
	nodePool      *blockchain.NodePool
	duration      time.Duration
}

// NewCrawlerPlanner returns new crawler planner instance with given context.
func NewCrawlerPlanner(
	log *zap.SugaredLogger,
	chain common.Chain,
	enable bool,
	blockConfirmations int64,
	inmemStorage *storage.Storage, nodePool *blockchain.NodePool, duration time.Duration) (*crawlPlanner, error) {

	return &crawlPlanner{
		log:           log,
		chain:         chain,
		enable:        enable,
		confirmations: blockConfirmations,
		inmemStorage:  inmemStorage,
		nodePool:      nodePool,
		duration:      duration,
	}, nil
}
func (p *crawlPlanner) Init() {
	if !p.enable {
		return
	}
	p.process()
}

func (p *crawlPlanner) Run() {
	if !p.enable {
		return
	}
	log := p.log.With("chain", p.chain)
	log.Infow("start to run planner", "enable", p.enable)
	for {
		p.process()
		time.Sleep(p.duration)
	}
}

func (p *crawlPlanner) process() {
	log := p.log.With("chain", p.chain)
	block, err := p.nodePool.GetBlockNumber(p.chain)
	if err != nil {
		log.Errorw("error when get latest block", "err", err)
		return
	}
	lastConfirm := int64(block) - p.confirmations
	latestBlock := p.inmemStorage.GetLatestBlock(p.chain)
	if lastConfirm >= latestBlock {
		p.inmemStorage.SetLatestBlock(p.chain, lastConfirm)
	}
}
