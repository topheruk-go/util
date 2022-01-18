package sqli

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/topheruk/go/src/encoding"
)

type DB struct {
	*sqlx.DB
}

func NewDB(db *sqlx.DB, driverName string) *DB {
	go db.MapperFunc(encoding.ToSnake)
	return &DB{DB: db}
}

func Open(driverName string, dataSourceName string) (*DB, error) {
	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	go db.MapperFunc(encoding.ToSnake)
	return &DB{DB: db}, nil
}

func MustOpen(driverName string, dataSourceName string) *DB {
	db := sqlx.MustOpen(driverName, dataSourceName)
	go db.MapperFunc(encoding.ToSnake)
	return &DB{DB: db}
}

func Connect(driverName string, dataSourceName string) (*DB, error) {
	return ConnectContext(context.Background(), driverName, dataSourceName)
}

func ConnectContext(ctx context.Context, driverName string, dataSourceName string) (*DB, error) {
	db, err := Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

func MustConnect(driverName string, dataSourceName string) *DB {
	db, err := Connect(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	return db
}

func (db *DB) Queryi(query string, args ...interface{}) error {
	return db.QueryiContext(context.Background(), query, args...)
}

func (db *DB) QueryiContext(ctx context.Context, query string, args ...interface{}) error {
	return preparedNamedQuery(ctx, db, query, args...)
}

func (db *DB) BeginTxi(opts *sql.TxOptions) (*Tx, error) {
	return db.BeginTxiContext(context.Background(), opts)
}

func (db *DB) BeginTxiContext(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	tx, err := db.BeginTxx(ctx, opts)
	return &Tx{Tx: tx}, err
}
