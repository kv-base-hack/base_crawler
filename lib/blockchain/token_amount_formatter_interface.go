package blockchain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	cm "github.com/kv-base-hack/base_crawler/common"
)

// TokenAmountFormatterInterface interface convert token amount from/to wei
type TokenAmountFormatterInterface interface {
	FromWei(cm.Chain, common.Address, *big.Int) (float64, error)
	GetDecimals(cm.Chain, common.Address) (int64, error)
}
