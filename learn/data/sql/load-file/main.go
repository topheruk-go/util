package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jszwec/csvutil"
	_ "github.com/mattn/go-sqlite3"
	"github.com/topheruk/go/src/database/sqli"
	"github.com/topheruk/go/src/measure"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// TODO: populate DB
	db := sqli.MustConnect("sqlite3", "learn/data/sql/load-file/.sqlite3")

	db.SetMaxOpenConns(3)

	db.Migrate(`
	drop table if exists "city";
	create table if not exists "city" (
		"name" text not null,
		"population" integer not null
	);`)

	_, err := loadFromCSV(db, "learn/data/sql/load-file/city.csv")
	if err != nil {
		return err
	}

	return nil
}

type city struct {
	Name       string `csv:"name"`
	Population int    `csv:"population"`
}

// "learn/data/sql/load-file/city.csv"
func loadFromCSV(db *sqli.DB, name string) (int, error) {
	defer measure.Time(time.Now(), "loadFromCSV")

	f, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return 0, err
	}

	var cc []city
	if err := csvutil.Unmarshal(b, &cc); err != nil {
		return 0, err
	}

	var g errgroup.Group
	for _, c := range cc {
		c := c
		g.Go(func() error {
			return db.Queryi(`insert into "city" values (:name,:population)`, &c)
		})
	}

	if err := g.Wait(); err != nil {
		return 0, err
	}

	return len(cc), nil
}
