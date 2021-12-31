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
	username = "topheruk"
	password = "T^*G7!Pf"
	host     = "192.168.1.173"
	cliPort  = 27017
	srvPort  = 8000
	uri      = fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, cliPort)
)

func run() (err error) {
	// TODO: arg flags parsing
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(ctx, uri, "company")
	if err != nil {
		return
	}
	defer db.Disconnect(ctx)

	app := app.New(db)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", srvPort),
		Handler: app,
	}

	fmt.Println("running... http://localhost:8000/api/v1/users/")
	return srv.ListenAndServe()
}
