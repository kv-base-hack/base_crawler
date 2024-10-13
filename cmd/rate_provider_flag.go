package main

import (
	"time"

	"github.com/urfave/cli/v2"
)

const (
	kaivestBinanceUrlFlag            = "kaivest-binance-url"
	kaivestRateWorkerDuration        = "kaivest-rate-worker-duration"
	defaultKaivestRateWorkerDuration = time.Minute
)

var ratesFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    kaivestBinanceUrlFlag,
		Usage:   "kaivest binance url",
		EnvVars: []string{"KAIVEST_BINANCE_URL"},
	},
	&cli.DurationFlag{
		Name:    kaivestRateWorkerDuration,
		Usage:   "worker duration to call to kaivest to get rate",
		EnvVars: []string{"KAIVEST_RATE_WORKER_DURATION"},
		Value:   defaultKaivestRateWorkerDuration,
	},
}

func NewRatesFlags() (flags []cli.Flag) {
	return ratesFlags
}
