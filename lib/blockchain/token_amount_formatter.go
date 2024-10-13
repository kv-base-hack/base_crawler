package blockchain

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	cm "github.com/kv-base-hack/base_crawler/common"
	"go.uber.org/zap"
)

var timeout = 30 * time.Second

// TokenAmountFormatter is a helper to convert token amount from/to wei
type TokenAmountFormatter struct {
	mu             sync.RWMutex
	cachedDecimals map[cm.Chain]map[string]int64
	nodePool       *NodePool
	log            *zap.SugaredLogger
}

// NewTokenAmountFormatter returns a new TokenAmountFormatter instance.
func NewTokenAmountFormatter(log *zap.SugaredLogger, nodePool *NodePool) (*TokenAmountFormatter, error) {
	return &TokenAmountFormatter{
		log: log,
		cachedDecimals: map[cm.Chain]map[string]int64{
			cm.ChainBase: {},
		},
		nodePool: nodePool,
	}, nil
}

// FromWei formats the given amount in wei to human friendly
// number with token decimals from contract.
func (f *TokenAmountFormatter) FromWei(chain cm.Chain, address common.Address, amount *big.Int) (float64, error) {
	if amount == nil {
		return 0, nil
	}
	decimals, err := f.GetDecimals(chain, address)
	if err != nil {
		return 0, fmt.Errorf("failed to get decimals: %s, err: %s", address.Hex(), err.Error())
	}
	floatAmount := new(big.Float).SetInt(amount)
	power := new(big.Float).SetInt(new(big.Int).Exp(
		big.NewInt(10), big.NewInt(decimals), nil,
	))
	res := new(big.Float).Quo(floatAmount, power)
	result, _ := res.Float64()
	return result, nil
}

// GetDecimals return decimals of a token
func (f *TokenAmountFormatter) GetDecimals(chain cm.Chain, address common.Address) (int64, error) {
	f.mu.RLock()
	if chainDecimal, exist := f.cachedDecimals[chain]; exist {
		if decimals, ok := chainDecimal[address.String()]; ok {
			f.mu.RUnlock()
			return decimals, nil
		}
	}
	f.mu.RUnlock()
	decimal, err := f.nodePool.GetDecimal(chain, address, timeout)
	if err != nil {
		return 0, err
	}
	f.mu.Lock()
	f.cachedDecimals[chain][address.String()] = int64(decimal)
	f.mu.Unlock()
	return int64(decimal), err
}
