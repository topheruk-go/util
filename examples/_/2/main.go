package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/topheruk/go/src/database/sqli"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if err := RemoveContents("app/tmp"); err != nil {
		return err
	}

	db := sqli.MustConnect("sqlite3", "app/.sqlite3")
	if _, err := db.ExecContext(context.TODO(), `
		DROP TABLE IF EXISTS laptop_loan;
		CREATE TABLE IF NOT EXISTS laptop_loan (
			id BLOB,
			student_id TEXT,
			start_date DATETIME NOT NULL,
			end_date DATETIME NOT NULL,
			tmp_path TEXT NOT NULL,
			PRIMARY KEY(id)
		);
	`); err != nil {
		return err
	}

	fmt.Println("Listening... http://localhost:8000")

	return http.ListenAndServe(":8000", newApp(db))
}
