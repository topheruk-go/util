package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/topheruk/go/learn/data/sql/example/app01/app"
	"github.com/topheruk/go/learn/data/sql/example/app01/service"
)

var (
	datasourceName = "./learn/fs/sql/example/app01/sql/.sqlite3"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := service.NewDB(ctx, datasourceName)
	go db.Migrate(
		`CREATE TABLE IF EXISTS table (
			"id" BLOB PRIMARY KEY,
			"email"	TEXT NOT NULL UNIQUE, 
			"password" BLOB NOT NULL,	
			"created_at" DATETIME NOT NULL
		)`,
	)

	app := app.New(db)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: app,
	}

	fmt.Println("listening... http://localhost:8000/ping")
	return srv.ListenAndServe()
}
