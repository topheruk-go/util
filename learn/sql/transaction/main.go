package main

import (
	"context"
	"fmt"
	"os"

	"github.com/topheruk/go/src/database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

var ctx = context.TODO()

func run() error {
	db := sql.New(ctx, "sqlite3", "./learn/sql/transaction/db.sqlite3")

	db.Query(ctx, "DROP TABLE IF EXISTS user")

	db.Migrate(ctx, map[string]string{
		"user": `"id" INTEGER PRIMARY KEY AUTOINCREMENT, "name" TEXT NOT NULL`,
	})
	// TODO: use an error channel
	db.Transaction(ctx, func(tx *sql.Tx) (err error) {
		type foo struct {
			Name string
		}
		if err = tx.Query(ctx, `INSERT INTO user (name) VALUES (:name)`, foo{"A"}); err != nil {
			return err
		}
		if err = tx.Query(ctx, `INSERT INTO user (name) VALUES (:name)`, foo{"B"}); err != nil {
			return err
		}
		if err = tx.Query(ctx, `INSERT INTO user (name) VALUES (:name)`, foo{"C"}); err != nil {
			return err
		}
		if err = tx.Query(ctx, `INSERT INTO user (name) VALUES (:name)`, foo{"D"}); err != nil {
			return err
		}
		if err = tx.Query(ctx, `INSERT INTO user (name) VALUES (:name)`, foo{"E"}); err != nil {
			return err
		}
		if err = tx.Query(ctx, `INSERT INTO user (name) VALUES (:name)`, foo{"F"}); err != nil {
			return err
		}
		return nil
	})

	return nil
}
