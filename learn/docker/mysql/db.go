package mssql

import (
	"context"
	"database/sql"
)

type Database struct {
	*sql.DB
}

func (d *Database) Insert(ctx context.Context, v interface{}) error {

	return nil
}

func (d *Database) SearchAll(ctx context.Context, v interface{}) error {
	d.Exec(`SELECT * FROM person`)
	return nil
}
