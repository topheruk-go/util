package sql

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v4"
)

func Query[T any](db *sql.DB, callback func(rows *sql.Rows, v *T) error, query string, args ...any) ([]T, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var vs []T
	for rows.Next() {
		var v T
		err = callback(rows, &v) //should be reference?
		if err != nil {
			return nil, err
		}
		vs = append(vs, v)
	}
	return vs, rows.Err()
}

func QueryPSQL[T any](db *pgx.Conn, callback func(rows pgx.Rows, v *T) error, query string, args ...any) ([]T, error) {
	rows, err := db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var vs []T
	for rows.Next() {
		var v T
		err = callback(rows, &v) //should be reference?
		if err != nil {
			return nil, err
		}
		vs = append(vs, v)
	}
	return vs, rows.Err()
}

func QueryRow(db *sql.DB, callback func(row *sql.Row) error, query string, args ...any) error {
	return callback(db.QueryRow(query))
}

func QueryRowPSQL(db *pgx.Conn, callback func(row pgx.Row) error, query string, args ...any) error {
	return callback(db.QueryRow(context.Background(), query, args...))
}

func Exec(db *sql.DB, query string, args ...any) error {
	_, err := db.Exec(query, args...)
	return err
}

func ExecPQSL(db *pgx.Conn, query string, args ...any) error {
	_, err := db.Exec(context.Background(), query, args...)
	return err
}
