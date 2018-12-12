package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/miya-masa/go-clean-webapp/transaction"
)

var txKey = struct{}{}

type tx struct {
	db *sqlx.DB
}

func NewTransaction(db *sqlx.DB) transaction.Transaction {
	return &tx{db: db}
}

func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, &txKey, tx)
	v, err := f(ctx)
	if err != nil {
		return nil, tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		return nil, tx.Rollback()
	}
	return v, nil
}

func GetTxFromContext(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(&txKey).(*sqlx.Tx)
	return tx, ok
}
