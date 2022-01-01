package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/topheruk/go/example/net/mongoDb/app"
	"github.com/topheruk/go/example/net/mongoDb/database"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var (
	username = ""
	password = ""
	host     = ""
	cliPort  = 27017
	srvPort  = 8000
)

func run() (err error) {
	user, pass, port := parseFlags()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(ctx, fmt.Sprintf("mongodb://%s:%s@localhost:27017", *user, *pass), "company")
	if err != nil {
		return
	}
	defer db.Disconnect(ctx)

	app := app.New(db)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: app,
	}

	fmt.Println("running... http://localhost:8000/api/v1/users/")
	return srv.ListenAndServe()
}
