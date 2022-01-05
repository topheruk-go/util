package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/topheruk/go/example/net/mongodb/app"
	db "github.com/topheruk/go/example/net/mongodb/database"
)

var user = flag.String("user", "", "client username")
var pass = flag.String("pass", "", "client username")
var port = flag.Int("port", 8000, "server port number")

var schemas = []string{
	"test/net/foobar/database/schema/foo.json",
	"test/net/foobar/database/schema/bar.json",
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

}

func run() (err error) {
	// user, pass, port := parseFlags()
	uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017", *user, *pass)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := db.New(ctx, uri, "foobar")
	if err != nil {
		return
	}
	defer db.Client().Disconnect(ctx)

	app := app.New(db)
	if err = app.SetupValidation(ctx, schemas...); err != nil {
		return
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: app,
	}

	fmt.Printf("running... http://localhost:%d/api/v1/\n", *port)
	return srv.ListenAndServe()
}

func parseFlags() (username, password *string, port *int) {
	username = flag.String("user", "", "database client username:password")
	password = flag.String("pass", "", "database client password")
	port = flag.Int("port", 8000, "server port number")

	flag.Parse()
	return
}
