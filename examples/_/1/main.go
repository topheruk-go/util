package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	db := sqlx.MustConnect("sqlite3", "example/.sqlite3")
	app := newApp(db)
	srv := &http.Server{
		Addr:    ":8000",
		Handler: app,
	}
	log.Println("Listening... http://localhost:8000/")
	return srv.ListenAndServe()
}
