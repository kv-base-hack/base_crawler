package main

import (
	"time"

	"github.com/urfave/cli/v2"
)

const (
	baseEnableFlag    = "base-enable"
	baseFromBlockFlag = "base-from-block"

	baseMaxWorkersFlag    = "base-max-workers"
	defaultBaseMaxWorkers = 3

	baseMaxBlocksForWorkerFlag    = "base-max-blocks-for-worker"
	defaultMaxBlocksForBaseWorker = 100

	basePlannerDurationFlag    = "base-planner-duration"
	basePlannerDurationDefault = time.Second * 1
)

var baseCrawlerFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    baseEnableFlag,
		EnvVars: []string{"BASE_ENABLE"},
	},
	&cli.StringFlag{
		Name:    baseFromBlockFlag,
		Usage:   "Fetch trade logs from block for base",
		EnvVars: []string{"BASE_FROM_BLOCK"},
	},
	&cli.IntFlag{
		Name:    baseMaxWorkersFlag,
		Usage:   "The maximum number of worker to fetch trade logs for base",
		EnvVars: []string{"BASE_MAX_WORKER"},
		Value:   defaultBaseMaxWorkers,
	},
	&cli.IntFlag{
		Name:    baseMaxBlocksForWorkerFlag,
		Usage:   "The maximum number of block on each query",
		EnvVars: []string{"BASE_MAX_BLOCKS_FOR_WORKER"},
		Value:   defaultMaxBlocksForBaseWorker,
	},
	&cli.DurationFlag{
		Name:    basePlannerDurationFlag,
		EnvVars: []string{"BASE_PLANNER_DURATION"},
		Value:   basePlannerDurationDefault,
	},
}

func newBaseCrawlerFlags() (flags []cli.Flag) {
	return baseCrawlerFlags
}
