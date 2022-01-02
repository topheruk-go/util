package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/topheruk/go/src/database"
	"github.com/topheruk/go/src/parse"
)

type User struct {
	Id string `json:"id" bson:"_id"`
}

var user = flag.String("user", "", "client username")
var pass = flag.String("pass", "", "client username")

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017", *user, *pass)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.NewMongoDB(ctx, uri, "foobar")
	if err != nil {
		return
	}
	defer db.Client().Disconnect(ctx)

	// setup validation
	doc, err := parse.BsonCmd("./learn/mongoDb/baz.schema.json")
	if err != nil {
		return
	}

	if err = db.RunCommand(ctx, doc).Err(); err != nil {
		return
	}
	// end

	// insert one with custom ID
	type response struct {
		Id    string `json:"-" bson:"_id"`
		Email string `json:"email" bson:"email"`
	}

	var res = response{Email: "40092@stu.fra.ac.uk"}
	res.Id = res.Email[:strings.Index(res.Email, "@")]

	err = db.InsertOne(ctx, "baz", res)
	if err != nil {
		return
	}

	return
}
