package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

var (
	datasourceName = "./learn/fs/sql/example/app01/web.sqlite3"
	sqlTables      = map[string]string{
		"user": `"id" BLOB PRIMARY KEY,	"email"	TEXT NOT NULL UNIQUE, "password" BLOB NOT NULL,	"created_at" DATETIME NOT NULL`,
	}
)

func run() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := newService(ctx, datasourceName)
	go db.Migrate(ctx, sqlTables)

	app := newApp(db)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: app,
	}

	// topheruk.sql.DB

	fmt.Println("listening... http://localhost:8000/ping")
	return srv.ListenAndServe()
}
