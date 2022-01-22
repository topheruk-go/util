package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/topheruk/go/example/app/sql-01/app"
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

	// drop table from database

	println("Listening... http://localhost:8000")
	return http.ListenAndServe(":8000", app.New())
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
