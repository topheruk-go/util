package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	db, err := sql.Open("sqlite3", "learn/data/sql/uuid/.db")
	if err != nil {
		return err
	}

	db.Exec(`
		drop table if exists "user";
		create table if not exists "user" (
			"id" blob,
			"name" text not null,
			primary key ("id")
		);
	`)

	id1 := uuid.New()
	if err := insert(db, id1, "foo"); err != nil {
		return err
	}

	id2 := uuid.New()
	if err := insert(db, id2, "baz"); err != nil {
		return err
	}

	user, err := get(db, id1)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", user.Name)

	return nil
}

type person struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func insert(db *sql.DB, uuid uuid.UUID, name string) error {
	stmt, err := db.Prepare(`insert into "user" values (?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	p := person{Name: name}
	p.ID = uuid

	stmt.Exec(p.ID, p.Name)
	return nil
}

func get(db *sql.DB, uuid uuid.UUID) (*person, error) {
	stmt, err := db.Prepare(`select * from "user" where id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(uuid)

	var q person
	row.Scan(&q.ID, &q.Name)
	return &q, nil
}
