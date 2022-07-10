// TODO -- [sql.TX, Must*]
package sql

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
)

var ErrPtr = fmt.Errorf("must pass by reference")

func Select[T any](db *sqlx.DB, query string, params ...any) ([]T, error) {
	var dest []T
	return dest, db.Select(&dest, query, params...)
}

func Get[T any](db *sqlx.DB, query string, params ...any) (*T, error) {
	var v T
	return &v, db.Get(&v, query, params...)
}

func Query[T any](db *sqlx.DB, dest T, query string, params ...any) error {
	if rv := reflect.ValueOf(dest); rv.Kind() == reflect.Ptr {
		switch reflect.Indirect(rv).Kind() {
		case reflect.Slice:
			return db.Select(dest, query, params...)
		default:
			return db.Get(dest, query, params...)
		}
	}
	return ErrPtr
}

type DB struct {
	c *sqlx.DB
}

func Connect(driverName, dataSourceName string) *DB {
	return &DB{c: sqlx.MustConnect(driverName, dataSourceName)}
}

func (db *DB) Close() error {
	return db.c.Close()
}

func (db DB) Prepare(query string, params ...any) (*Stmt, error) {
	stmt, err := db.c.Preparex(query)
	return &Stmt{c: stmt}, err
}

func (db DB) Exec(query string, params ...any) (sql.Result, error) {
	return db.c.Exec(query, params...)
}

func (db *DB) Query(dest any, query string, params ...any) error {
	return Query(db.c, dest, query, params...)
}

type Stmt struct {
	c *sqlx.Stmt
}

func (stmt Stmt) Exec(params ...any) (sql.Result, error) {
	return stmt.c.Exec(params...)
}

func (stmt Stmt) Query(dest any, params ...any) error {
	if rv := reflect.ValueOf(dest); rv.Kind() == reflect.Ptr {
		switch reflect.Indirect(rv).Kind() {
		case reflect.Slice:
			return stmt.c.Select(dest, params...)
		default:
			return stmt.c.Get(dest, params...)
		}
	}
	return ErrPtr
}
