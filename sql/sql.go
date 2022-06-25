package sql

import (
	"database/sql"
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

func QueryRow(db *sql.DB, callback func(row *sql.Row) error, query string, args ...any) error {
	return callback(db.QueryRow(query))
}

func Exec(db *sql.DB, query string, args ...any) error {
	_, err := db.Exec(query, args...)
	return err
}
