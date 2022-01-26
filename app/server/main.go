package main

import (
	"fmt"
	"net/http"
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

func refreshApp(db *sqli.DB) error {
	if err := removeAll("app/tmp"); err != nil {
		return err
	}

	// db := sqli.MustConnect("sqlite3", "app/data/db.sqlite3")
	// defer db.Close()
	_, err := db.Exec(`DROP TABLE IF EXISTS "loan"`)
	if err != nil {
		return err
	}
	return nil
}

func run() error {
	db := sqli.MustConnect("sqlite3", "app/data/db.sqlite3")
	defer db.Close()

	if err := refreshApp(db); err != nil {
		return err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS "loan" (
			"id" BLOB,
			"name" TEXT NOT NULL,
			"age" INTEGER NOT NULL,
			"path" TEXT NOT NULL UNIQUE,
			PRIMARY KEY ("id")
		)`,
	); err != nil {
		return err
	}

	return http.ListenAndServe(":8000", newApp(db))
}
