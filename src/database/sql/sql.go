package sql

import (
	"context"
	"database/sql"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/topheruk/go/src/encoding"
)

type DB struct {
	sync.WaitGroup
	s *sqlx.DB
}

func New(ctx context.Context, driverName string, dataSourceName string) *DB {
	s := sqlx.MustOpen(driverName, dataSourceName)
	go s.MapperFunc(encoding.ToSnake)
	return &DB{
		s: s,
	}
}

func (db *DB) Query(ctx context.Context, query string, args ...interface{}) error {
	return Query(ctx, db.s, query, args...)
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	tx, err := db.s.BeginTxx(ctx, opts)
	return &Tx{x: tx}, err
}
