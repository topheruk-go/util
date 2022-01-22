package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/topheruk/go/example/app/sql-01/app/v2"
	"github.com/topheruk/go/example/app/sql-01/db/v2"

	"github.com/topheruk/go/src/fs"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if err := fs.RemoveContents("example/app/sql-01/tmp"); err != nil {
		return err
	}

	db := db.New("sqlite3", "example/app/sql-01/bin/test-01.sqlite3")
	if _, err := db.ExecContext(context.TODO(), `
		drop table if exists "laptop_loan";
		create table if not exists "laptop_loan" (
			"id" blob,
			"student_id" text,
			"start_date" datetime not null,
			"end_date" datetime not null,
			"tmp_path" text not null,
			primary key("id")
		);
	`); err != nil {
		return err
	}

	println("Listening... http://localhost:8000")
	return http.ListenAndServe(":8000", app.New(db))
}
