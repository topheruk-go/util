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
	db, err := sql.Open("sqlite3", "learn/data/sql/uuid/.sqlite3")
	if err != nil {
		return err
	}

	// db.Exec(`
	// 	DROP TABLE IF EXISTS "user";
	// 	CREATE TABLE IF NOT EXISTS "user" (
	// 		"id" BLOB PRIMARY KEY,
	// 		"name" TEXT NOT NULL
	// 	);
	// `)

	db.Exec(`
		drop table if exists "user";
		create table if not exists "user" (
			"id" blob primary key,
			"name" text not null
		);
	`)

	type person struct {
		ID   uuid.UUID `db:"id"`
		Name string    `db:"name"`
	}

	p := person{Name: "Matt"}
	p.ID, _ = uuid.NewUUID()

	stmt, _ := db.Prepare(`insert into "user" values (?, ?)`)
	defer stmt.Close()
	stmt.Exec(p.ID, p.Name)

	// get user
	stmt, _ = db.Prepare(`select * from "user" where id = ?`)
	defer stmt.Close()
	row := stmt.QueryRow(p.ID)

	var q person
	row.Scan(&q.ID, &q.Name)
	fmt.Println(q)

	return nil
}
