package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	ether "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	cm "github.com/kv-base-hack/base_crawler/common"
	"github.com/kv-base-hack/base_crawler/lib/contracts"
	"go.uber.org/zap"
)

const sleepTime = time.Second / 2

// enumer -type=NodeType -linecomment -json=true -text=true -sql=true
type NodeType uint64

const (
	NodeInfura NodeType = iota + 1 // infura
)

type nodeEnpoint struct {
	nodeType          NodeType
	client            *ethclient.Client
	blockUsedDuration time.Duration // we dont use this key if it still in relax time
	blockUntil        time.Time
	usedCount         int
}

type ChainNode struct {
	mutex        sync.Mutex
	client       []nodeEnpoint
	keys         []string
	nodeTypeUsed int
}

type NodePool struct {
	log *zap.SugaredLogger

	chains map[cm.Chain]*ChainNode
}

func NewNodePool(log *zap.SugaredLogger, infuraUrl, infuraKey string, infuraNodeRelaxDuration time.Duration) (*NodePool, error) {
	baseClient := make([]nodeEnpoint, 0)
	baseKeys := []string{}
	nodeBaseTypeUsed := 0
	if len(infuraKey) > 0 {
		nodeBaseTypeUsed++
		keys := strings.Split(infuraKey, ",")
		baseKeys = append(baseKeys, keys...)
		for _, k := range keys {
			client, err := NewEthereumClient(infuraUrl + "/" + k)
			if err != nil {
				return &NodePool{}, err
			}
			baseClient = append(baseClient, nodeEnpoint{
				nodeType:          NodeInfura,
				client:            client,
				blockUsedDuration: infuraNodeRelaxDuration,
			})
		}
	}

	return &NodePool{
		log: log,
		chains: map[cm.Chain]*ChainNode{
			cm.ChainBase: {
				client:       baseClient,
				keys:         baseKeys,
				nodeTypeUsed: nodeBaseTypeUsed,
			},
		},
	}, nil
}

// we try to get data from many node endpoint : infura, alchemy, chain node, quick node, etc
// when we got failed from a node endpoint, we will try to get it from other node types
func (n *NodePool) getWorkerWithFailedNode(nodeChain *ChainNode, failedNodeType []NodeType) int {
	nodeChain.mutex.Lock()
	defer nodeChain.mutex.Unlock()
	found := -1
	minUsed := 1_000_000_000
	for i := 0; i < len(nodeChain.client); i++ {
		if time.Now().After(nodeChain.client[i].blockUntil) {
			failed := false
			for _, j := range failedNodeType {
				if j == nodeChain.client[i].nodeType {
					failed = true
					break
				}
			}

			if failed {
				continue
			}

			if nodeChain.client[i].usedCount < minUsed {
				minUsed = nodeChain.client[i].usedCount
				found = i
			}
		}
	}
	if found != -1 {
		nodeChain.client[found].usedCount++
		nodeChain.client[found].blockUntil = time.Now().Add(nodeChain.client[found].blockUsedDuration)
	}
	return found
}

func (n *NodePool) GetTxData(chain cm.Chain, txHash common.Hash, timeout time.Duration) (*types.Receipt, common.Address, error) {
	var txReceipt *types.Receipt
	var tx *types.Transaction
	var err error
	var txSender common.Address
	log := n.log.With("txHash", txHash, "chain", chain, "action", "GetTxData")
	// now := time.Now()

	failedNode := []NodeType{}
	nodeChain, exist := n.chains[chain]
	if !exist {
		return nil, common.Address{}, fmt.Errorf("not found this chain")
	}
	for {
		if len(failedNode) == nodeChain.nodeTypeUsed {
			break
		}
		workerID := n.getWorkerWithFailedNode(nodeChain, failedNode)
		if workerID == -1 {
			time.Sleep(sleepTime)
		} else {
			ctx := context.Background()

			// startTime := time.Now()
			// defer func() {
			// log.Debugw("end. release worker", "workerID", workerID, "duration", time.Since(now), "workDur", time.Since(startTime))
			// }()

			// log.Debugw("doing at", "workerID", workerID)
			txReceipt, err = nodeChain.client[workerID].client.TransactionReceipt(ctx, txHash)
			if err != nil {
				log.Errorw("error when get tx receipt", "key", nodeChain.keys[workerID], "err", err)
				failedNode = append(failedNode, nodeChain.client[workerID].nodeType)
				continue
			}
			// log.Debugw("start get tx by hash", "workerID", workerID)

			tx, _, err = nodeChain.client[workerID].client.TransactionByHash(ctx, txHash)
			if err != nil {
				log.Errorw("error when get tx by hash", "key", nodeChain.keys[workerID], "err", err)
				failedNode = append(failedNode, nodeChain.client[workerID].nodeType)
				continue
			}

			// log.Debugw("start get sender", "workerID", workerID)
			txSender, err = nodeChain.client[workerID].client.TransactionSender(ctx, tx, txReceipt.BlockHash, txReceipt.TransactionIndex)
			if err != nil {
				log.Errorw("error when get tx sender", "key", nodeChain.keys[workerID], "err", err)
				failedNode = append(failedNode, nodeChain.client[workerID].nodeType)
				continue
			}
			// log.Debugw("finished", "workerID", workerID)
			break
		}
	}

	return txReceipt, txSender, err
}

func (n *NodePool) FilterLogs(chain cm.Chain, query ether.FilterQuery) ([]types.Log, error) {
	var logs []types.Log
	var err error
	log := n.log.With("action", "FilterLogs", "chain", chain)
	failedNode := []NodeType{}

	nodeChain, exist := n.chains[chain]
	if !exist {
		return nil, fmt.Errorf("not found this chain")
	}

	for {
		if len(failedNode) == nodeChain.nodeTypeUsed {
			break
		}
		workerID := n.getWorkerWithFailedNode(nodeChain, failedNode)
		if workerID == -1 {
			time.Sleep(sleepTime)
		} else {

			ctx := context.Background()

			// log.Debugw("doing at", "workerID", workerID)
			logs, err = nodeChain.client[workerID].client.FilterLogs(ctx, query)
			if err != nil {
				log.Errorw("error when filter logs", "key", nodeChain.keys[workerID], "err", err)
				failedNode = append(failedNode, nodeChain.client[workerID].nodeType)
				continue
			}
			// log.Debugw("finished", "workerID", workerID)
			break
		}
	}
	return logs, nil
}

func (n *NodePool) GetBlockTimeStamp(chain cm.Chain, timeout time.Duration, blockNumber uint64) (*types.Header, error) {
	var header *types.Header
	var err error
	log := n.log.With("action", "GetBlockTimeStamp", "chain", chain, "blockNumber", blockNumber)

	failedNode := []NodeType{}
	nodeChain, exist := n.chains[chain]
	if !exist {
		return nil, fmt.Errorf("not found this chain")
	}
	for {
		if len(failedNode) == nodeChain.nodeTypeUsed {
			break
		}
		workerID := n.getWorkerWithFailedNode(nodeChain, failedNode)
		if workerID == -1 {
			time.Sleep(sleepTime)
		} else {

			ctx := context.Background()

			// log.Debugw("doing at", "workerID", workerID)
			header, err = nodeChain.client[workerID].client.HeaderByNumber(ctx, big.NewInt(int64(blockNumber)))
			if err != nil && header == nil {
				log.Errorw("error when get header by number", "key", nodeChain.keys[workerID], "err", err)
				failedNode = append(failedNode, nodeChain.client[workerID].nodeType)
				continue
			}
			if err != nil && strings.Contains(err.Error(), "missing required field") {
				log.Errorw("ignore block header error", "key", nodeChain.client[workerID], "err", err)
			}
			// log.Debugw("finished", "workerID", workerID)
			break
		}
	}
	return header, nil
}

func (n *NodePool) GetBlockNumber(chain cm.Chain) (uint64, error) {
	var block uint64
	var err error
	log := n.log.With("action", "GetBlockNumber", "chain", chain)

	failedNode := []NodeType{}
	nodeChain, exist := n.chains[chain]
	if !exist {
		return 0, fmt.Errorf("not found this chain")
	}
	for {
		if len(failedNode) == nodeChain.nodeTypeUsed {
			break
		}
		workerID := n.getWorkerWithFailedNode(nodeChain, failedNode)
		if workerID == -1 {
			time.Sleep(sleepTime)
		} else {
			ctx := context.Background()
			// log.Debugw("doing at", "workerID", workerID)
			block, err = nodeChain.client[workerID].client.BlockNumber(ctx)
			if err != nil {
				log.Errorw("error when get block number", "key", nodeChain.keys[workerID], "err", err)
				failedNode = append(failedNode, nodeChain.client[workerID].nodeType)
				continue
			}
			// log.Debugw("finished", "workerID", workerID)
			break
		}
	}
	return block, nil
}

func (n *NodePool) GetDecimal(chain cm.Chain, address common.Address, timeout time.Duration) (uint8, error) {
	var decimal uint8
	// log := n.log.With("action", "GetDecimal", "address", address)
	failedNode := []NodeType{}
	nodeChain, exist := n.chains[chain]
	if !exist {
		return 0, fmt.Errorf("not found this chain")
	}
	for {
		if len(failedNode) == nodeChain.nodeTypeUsed {
			break
		}
		workerID := n.getWorkerWithFailedNode(nodeChain, failedNode)
		if workerID == -1 {
			time.Sleep(sleepTime)
		} else {
			ctx := context.Background()

			// log.Debugw("doing at", "workerID", workerID)
			tokenContract, err := contracts.NewERC20(address, nodeChain.client[workerID].client)
			if err != nil {
				failedNode = append(failedNode, nodeChain.client[workerID].nodeType)
				continue
			}
			decimal, err = tokenContract.Decimals(&bind.CallOpts{Context: ctx})
			if err != nil {
				failedNode = append(failedNode, nodeChain.client[workerID].nodeType)
				continue
			}
			// log.Debugw("finished", "workerID", workerID)
			break
		}
	}
	return decimal, nil
}
