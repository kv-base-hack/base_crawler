package main

import (
	"os"
	"sort"
	"time"

	"github.com/joho/godotenv"
	obc "github.com/kv-base-hack/base-client/kaivest_binance"
	"github.com/kv-base-hack/base_crawler/common"
	"github.com/kv-base-hack/base_crawler/lib/blockchain"
	"github.com/kv-base-hack/base_crawler/lib/tokenrate"
	basetopic "github.com/kv-base-hack/base_crawler/lib/topic/base"
	"github.com/kv-base-hack/base_crawler/storage"
	"github.com/kv-base-hack/base_crawler/storage/db"
	"github.com/kv-base-hack/base_crawler/workers"
	"github.com/kv-base-hack/common/logger"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Load()
	app := cli.NewApp()
	app.Action = run
	app.Flags = append(app.Flags, logger.NewSentryFlags()...)
	app.Flags = append(app.Flags, newBaseCrawlerFlags()...)
	app.Flags = append(app.Flags, NewPostgreSQLFlags()...)
	app.Flags = append(app.Flags, NewRatesFlags()...)
	sort.Sort(cli.FlagsByName(app.Flags))

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func run(c *cli.Context) error {
	logger, flusher, err := logger.NewLogger(c)
	if err != nil {
		return err
	}
	defer flusher()

	zap.ReplaceGlobals(logger)
	log := logger.Sugar()
	log.Debugw("Starting application...")

	store := storage.NewStorage()

	database, err := NewDBFromContext(c)
	if err != nil {
		log.Errorw("error when connect to database", "err", err)
		return err
	}

	pg := db.NewPostgres(database)

	baseFromBlock := c.Int64(baseFromBlockFlag)

	baseMaxWorkers := c.Int(baseMaxWorkersFlag)
	baseMaxBlocksForWorker := c.Int64(baseMaxBlocksForWorkerFlag)

	nodeRequestTimeOut := time.Second * 10
	nodePool, err := blockchain.NewNodePool(log, c.String(infuraNodePoolDefaultUrlFlag), c.String(infuraNodePoolKeysFlag), c.Duration(infuraNodeRelaxDuration))
	if err != nil {
		log.Errorw("error when create node pool", "err", err)
		return err
	}

	tokenAmountFormatter, err := blockchain.NewTokenAmountFormatter(log, nodePool)
	if err != nil {
		log.Errorw("error when create token amount formatter", "err", err)
		return err
	}

	resolver, err := blockchain.NewBlockTimeResolver(log, nodePool, nodeRequestTimeOut)
	if err != nil {
		log.Errorw("error when create block time resolver", "err", err)
		return err
	}

	err = resolver.Sync(common.ChainBase, uint64(baseFromBlock), uint64(baseFromBlock))
	if err != nil {
		log.Errorw("error when sync block timestamp", "baseFromBlock", baseFromBlock, "err", err)
		return err
	}

	baseFromBlockTimestamp, err := resolver.Resolve(common.ChainBase, uint64(baseFromBlock))
	if err != nil {
		log.Errorw("error when create block timestamp for bsc", "err", err)
		return err
	}

	kaivestBinance := obc.NewKaivestBinanceClient(c.String(kaivestBinanceUrlFlag))
	kaivestBinanceClient := tokenrate.NewKaivestBinance(log, c.Duration(kaivestRateWorkerDuration),
		baseFromBlockTimestamp, kaivestBinance)
	kaivestBinanceClient.Init()
	go kaivestBinanceClient.Run()

	blockConfirmations := int64(25)

	baseEnable := c.Bool(baseEnableFlag)

	basePlanner, err := workers.NewCrawlerPlanner(log, common.ChainBase, baseEnable,
		blockConfirmations, store, nodePool, c.Duration(basePlannerDurationFlag))
	if err != nil {
		log.Errorw("error when create planner for base", "err", err)
		return err
	}
	basePlanner.Init()
	go basePlanner.Run()

	baseTopic, err := basetopic.NewBaseTopic(log)
	if err != nil {
		log.Errorw("error when create bsc swap topic", "err", err)
		return err
	}

	baseCrawler := workers.NewEvmCrawler(log, common.ChainBase, baseEnable, store,
		baseFromBlock, pg, tokenAmountFormatter, baseMaxWorkers,
		baseMaxBlocksForWorker, nodeRequestTimeOut, resolver, kaivestBinanceClient,
		baseTopic, nodePool, db.BaseTradeLogs, db.BaseTransferTable)
	err = baseCrawler.Init()
	if err != nil {
		log.Errorw("error when init data", "err", err)
		return err
	}

	return baseCrawler.Run()
}
