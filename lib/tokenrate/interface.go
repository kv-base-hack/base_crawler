package tokenrate

import (
	"time"

	"github.com/kv-base-hack/base_crawler/common"
)

// TokenUsdtRateProvider is the common interface to query historical
// rates of token with USD.
type TokenUsdtRateProvider interface {
	USDRate(common.Chain, string, time.Time) (float64, error)
	// SolRate(time.Time) (float64, error)
	// Name return name of provider
	Name() string
}
