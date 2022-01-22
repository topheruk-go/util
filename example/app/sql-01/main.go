package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/topheruk/go/example/app/sql-01/app/v2"
	"github.com/topheruk/go/example/app/sql-01/db/v2"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if err := RemoveContents("example/app/sql-01/tmp"); err != nil {
		return err
	}

	db := db.New("sqlite3", "example/app/sql-01/bin/test-01.sqlite3")
	if _, err := db.ExecContext(context.TODO(), `
		drop table if exists "laptop_loan";
		create table if not exists "laptop_loan" (
			id blob,
			student_id text,
			start_date datetime not null,
			end_date datetime not null,
			tmp_path text not null,
			primary key("id")
		);
	`); err != nil {
		return err
	}

	println("Listening... http://localhost:8000")
	return http.ListenAndServe(":8000", app.New(db))
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
