package main

import (
	"time"

	"github.com/urfave/cli/v2"
)

const (
	infuraNodePoolDefaultUrlFlag = "infura-node-pool-default-url"
	infuraNodePoolKeysFlag       = "infura-node-pool-keys"
	infuraNodeRelaxDuration      = "infura-node-relax-duration"
)

// NewEthereumNodeFlags returns cli flag for ethereum node url input
func NewEthereumNodeFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    infuraNodePoolDefaultUrlFlag,
			Usage:   "infura node default url",
			EnvVars: []string{"INFURA_NODE_POOL_DEFAULT_URL"},
		},
		&cli.StringFlag{
			Name:    infuraNodePoolKeysFlag,
			Usage:   "infura node keys",
			EnvVars: []string{"INFURA_NODE_POOL_KEYS"},
		},
		&cli.DurationFlag{
			Name:    infuraNodeRelaxDuration,
			Usage:   "infura node relax duration",
			Value:   time.Second / 2,
			EnvVars: []string{"INFURA_NODE_RELAX_DURATION"},
		},
	}
}
