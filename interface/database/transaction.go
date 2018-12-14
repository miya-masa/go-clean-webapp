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
	tx, err := t.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, &txKey, tx)
	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return v, nil
}

func GetTx(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(&txKey).(*sqlx.Tx)
	return tx, ok
}

func DoInTx(db *sqlx.DB, f func(tx *sqlx.Tx) (interface{}, error)) (interface{}, error) {
	// start transaction
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	// execution
	v, err := f(tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// commit
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return v, nil
}
