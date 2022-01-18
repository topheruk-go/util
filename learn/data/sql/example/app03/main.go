package main

import (
	"fmt"
	"os"

	"github.com/topheruk/go/src/database/sqli"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	db := sqli.MustOpen("sqlite3", "./learn/sql/transaction/db.sqlite3")

	db.Queryi("DROP TABLE IF EXISTS user")

	db.Migrate(
		`CREATE IF NOT EXISTS user (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"name" TEXT NOT NULL
		);`,
	)

	db.Transaction(func(tx *sqli.Tx) (err error) {
		type foo struct {
			Name string
		}

		if err = tx.Queryi(`INSERT INTO user (name) VALUES (:name)`, foo{"A"}); err != nil {
			return err
		}
		if err = tx.Queryi(`INSERT INTO user (name) VALUES (:name)`, foo{"B"}); err != nil {
			return err
		}
		if err = tx.Queryi(`INSERT INTO user (name) VALUES (:name)`, foo{"C"}); err != nil {
			return err
		}
		if err = tx.Queryi(`INSERT INTO user (name) VALUES (:name)`, foo{"D"}); err != nil {
			return err
		}
		if err = tx.Queryi(`INSERT INTO user (name) VALUES (:name)`, foo{"E"}); err != nil {
			return err
		}
		if err = tx.Queryi(`INSERT INTO user (name) VALUES (:name)`, foo{"F"}); err != nil {
			return err
		}

		return nil
	})

	return nil
}
