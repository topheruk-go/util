package sql

import (
	"context"
	"database/sql"
)

func Query[T any](db *sql.DB, callback func(rows *sql.Rows, v *T) error, query string, args ...any) ([]T, error) {
	return QueryContext(context.Background(), db, callback, query, args...)
}

func QueryContext[T any](ctx context.Context, db *sql.DB, scanner func(rows *sql.Rows, v *T) error, query string, args ...any) ([]T, error) {
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var vs []T
	for rows.Next() {
		var v T
		err = scanner(rows, &v) //should be reference?
		if err != nil {
			return nil, err
		}
		vs = append(vs, v)
	}
	return vs, rows.Err()
}

func QueryRow(db *sql.DB, scanner func(row *sql.Row) error, query string, args ...any) error {
	return QueryRowContext(context.Background(), db, scanner, query, args...)
}

func QueryRowContext(ctx context.Context, db *sql.DB, scanner func(row *sql.Row) error, query string, args ...any) error {
	return scanner(db.QueryRowContext(ctx, query))
}

func Exec(db *sql.DB, query string, args ...any) error {
	return ExecContext(context.Background(), db, query, args...)
}

func ExecContext(ctx context.Context, db *sql.DB, query string, args ...any) error {
	_, err := db.ExecContext(ctx, query, args...)
	return err
}
