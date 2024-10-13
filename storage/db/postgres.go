package db

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/kv-base-hack/base_crawler/common"
	_ "github.com/lib/pq" // sql driver name: "postgres"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	BaseTradeLogs     = "base_trade_logs"
	BaseTransferTable = "base_transfer_logs"
)

type Postgres struct {
	db *sqlx.DB
	l  *zap.SugaredLogger
}

func NewPostgres(db *sqlx.DB) *Postgres {
	return &Postgres{
		db: db,
		l:  zap.S(),
	}
}

func (pg *Postgres) GetLastStoredBlock(table string) (int64, error) {
	query, _, err := sq.
		Select("MAX(block_number) as block_number").
		From(table).ToSql()
	if err != nil {
		return 0, err
	}
	var result int64
	err = pg.db.Get(&result, query)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (pg *Postgres) GetSmallestStoredBlock(table string) (int64, error) {
	query, _, err := sq.
		Select("MIN(block_number) as block_number").
		From(table).ToSql()
	if err != nil {
		return 0, err
	}
	var result int64
	err = pg.db.Get(&result, query)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (pg *Postgres) getTradeTx(tradeLogs []common.Tradelog, table string) (string, []interface{}, error) {
	b := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Insert(table).
		Columns("block_timestamp", "block_number", "tx_index", "tx_hash", "sender",
			"token_in_address", "token_in_amount", "token_in_usdt_rate",
			"token_out_address", "token_out_amount", "token_out_usdt_rate",
			"native_token_usdt_rate", "routes", "created",
		)

	for _, t := range tradeLogs {
		b = b.Values(t.BlockTimestamp, t.BlockNumber, t.TxIndex, t.TxHash.String(), t.Sender.String(),
			t.TokenInAddress.String(), t.TokenInAmount, t.TokenInUsdtRate,
			t.TokenOutAddress.String(), t.TokenOutAmount, t.TokenOutUsdtRate,
			t.NativeTokenUsdtRate, t.Routes, time.Now(),
		)
	}

	q, p, err := b.Suffix("ON CONFLICT (block_number, tx_index) DO nothing").ToSql()
	return q, p, err
}
func (pg *Postgres) getTransferTx(transferLogs []common.Transferlog, table string) (string, []interface{}, error) {
	b := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Insert(table).
		Columns("block_timestamp", "block_number", "tx_index", "tx_hash",
			"from_address", "to_address",
			"token_address", "token_amount",
			"is_cex_in", "exchange", "created",
		)

	for _, t := range transferLogs {
		b = b.Values(t.BlockTimestamp, t.BlockNumber, t.TxIndex, t.TxHash.String(),
			t.FromAddress.String(), t.ToAddress.String(),
			t.TokenAddress.String(), t.TokenAmount,
			t.IsCexIn, t.Exchange.String(), time.Now(),
		)
	}

	q, p, err := b.Suffix("ON CONFLICT (block_number, tx_index) DO nothing").ToSql()
	return q, p, err
}

func (pg *Postgres) SaveLogs(tradeLogs []common.Tradelog, transferLogs []common.Transferlog, tradeTable, transferTable string) error {
	tx, err := pg.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "create transaction error")
	}
	defer func(t *sqlx.Tx) {
		if err := t.Rollback(); err != nil && err != sql.ErrTxDone {
			pg.l.Errorw("failed to roll back transaction", "err", err)
		}
	}(tx)
	if len(tradeLogs) > 0 {
		qTrade, pTrade, errTrade := pg.getTradeTx(tradeLogs, tradeTable)
		if errTrade != nil {
			return err
		}
		if _, err = tx.Exec(qTrade, pTrade...); err != nil {
			pg.l.Errorw("error when exec tx for eth trade", "err", err)
			return err
		}
	}

	if len(transferLogs) > 0 {
		qTransfer, pTransfer, errTransfer := pg.getTransferTx(transferLogs, transferTable)
		if errTransfer != nil {
			return err
		}

		if _, err = tx.Exec(qTransfer, pTransfer...); err != nil {
			pg.l.Errorw("error when exec tx for eth transfer", "err", err)
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		pg.l.Infow("cannot add eth data to database", "err", err)
		return err
	}

	return nil
}
