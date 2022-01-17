package sql

import (
	"context"
	"sync"

	"github.com/jmoiron/sqlx"
)

type Tx struct {
	sync.WaitGroup
	x *sqlx.Tx
}

func (tx *Tx) Query(ctx context.Context, query string, args ...interface{}) error {
	return Query(ctx, tx.x, query, args...)
}

func (tx *Tx) Commit() error   { return tx.x.Commit() }
func (tx *Tx) Rollback() error { return tx.x.Rollback() }

func (db *DB) Transaction(ctx context.Context, doSomething func(tx *Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err = doSomething(tx); err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
