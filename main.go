package main

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	app "github.com/topheruk/go/app/rest/demo/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello, World!")

	db := sqlx.MustConnect("sqlite3", ":memory:")
	app := app.New(db)

	db.MustExec(`INSERT INTO person VALUES (1, "John")`)

	http.ListenAndServe(":8000", app)
}
