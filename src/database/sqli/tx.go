package sqli

import (
	"context"
	"sync"

	"github.com/jmoiron/sqlx"
)

type Tx struct {
	sync.WaitGroup
	*sqlx.Tx
}

func (tx *Tx) Queryi(query string, args ...interface{}) error {
	return tx.QueryiContext(context.Background(), query, args...)
}

func (tx *Tx) QueryiContext(ctx context.Context, query string, args ...interface{}) error {
	return preparedNamedQuery(ctx, tx, query, args...)
}

func (db *DB) Transaction(fn func(tx *Tx) error) error {
	return db.TransactionContext(context.Background(), fn)
}

func (db *DB) TransactionContext(ctx context.Context, fn func(tx *Tx) error) error {
	return transaction(ctx, db, fn)
}

func transaction(ctx context.Context, db *DB, fn func(tx *Tx) error) error {
	tx, err := db.BeginTxiContext(ctx, nil)
	if err != nil {
		return err
	}
	if err = fn(tx); err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
