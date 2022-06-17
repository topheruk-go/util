package pg

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func Query[T any](db *pgx.Conn, query string, scanner func(rows pgx.Rows, v *T) error, args ...any) ([]T, error) {
	return QueryContext(context.Background(), db, query, scanner, args...)
}

func QueryContext[T any](ctx context.Context, db *pgx.Conn, query string, scanner func(rows pgx.Rows, v *T) error, args ...any) ([]T, error) {
	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var vs []T
	for rows.Next() {
		var v T
		err = scanner(rows, &v)
		if err != nil {
			return nil, err
		}
		vs = append(vs, v)
	}
	return vs, rows.Err()
}

func QueryRow(db *pgx.Conn, query string, scanner func(row pgx.Row) error, args ...any) error {
	return scanner(db.QueryRow(context.Background(), query, args...))
}

func Exec(db *pgx.Conn, query string, args ...any) error {
	_, err := db.Exec(context.Background(), query, args...)
	return err
}
