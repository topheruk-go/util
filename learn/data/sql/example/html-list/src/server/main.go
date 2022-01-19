package main

import (
	"fmt"
	"log"
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

func run() error {
	db := sqli.MustConnect("sqlite3", "html-list/src/db/example.sqlite3")
	go db.Migrate(`
	DROP TABLE IF EXISTS "account";
	CREATE TABLE IF NOT EXISTS "account" (
		"id" integer primary key autoincrement,
		"name" text not null,
		"age" integer not null,
		"address" text not null,
		"salary" integer not null 
	);
	INSERT INTO "account" ("name","age","address","salary") VALUES
		("Paul",32,"California",20000.0),
		("Allen",25,"Texas",15000.0),
		("Teddy",23,"Norway",20000.0),
		("Mark",25,"Rich-Mond",65000.0),
		("David",27,"Texas",85000.0),
		("Kim",22,"South-Hall",45000.0),
		("James",24,"Houston",10000.0)
	`)

	log.Printf("listening... http://localhost:8000/")
	return http.ListenAndServe(":8000", newApp(db))
}
