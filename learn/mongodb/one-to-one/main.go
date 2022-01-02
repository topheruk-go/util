package main

import (
	"context"
	"encoding/json"
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
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type TestCase struct {
	content string
	id      string
}

var insertOneTestCases = []TestCase{
	{content: `{"email":"40092345@stu.fra.ac.uk"}`},
	{content: `{"email":"40092346@stu.fra.ac.uk"}`},
	{content: `{"email":"40092347@stu.fra.ac.uk"}`},
	{content: `{"email":"40092348@stu.fra.ac.uk"}`},
	{content: `{"email":"Rahim.Affulbrown@fra.ac.uk"}`},
}

var deleteTestCases = []TestCase{
	{id: "40092345"},
	{id: "40092346"},
	{id: "40092347"},
	{id: "40092348"},
	{id: "Rahim.Affulbrown"},
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
	doc, err := parse.BsonCmd("./learn/mongodb/one-to-one/baz.schema.json")
	if err != nil {
		return
	}

	if err = db.RunCommand(ctx, doc).Err(); err != nil {
		return
	}

	// Insert One
	for _, tc := range insertOneTestCases {
		if err = InsertOne(ctx, db, &tc); err != nil {
			return
		}
	}

	// Delete One
	for _, tc := range deleteTestCases {
		if err = DeleteOne(ctx, db, &tc); err != nil {
			return
		}
	}

	return
}

func InsertOne(ctx context.Context, db *database.MongoDB, tc *TestCase) (err error) {
	type response struct {
		Id    string `json:"-" bson:"_id"`
		Email string `json:"email" bson:"email"`
	}

	var res response
	if err = json.NewDecoder(strings.NewReader(tc.content)).Decode(&res); err != nil {
		return
	}
	res.Id = res.Email[:strings.Index(res.Email, "@")]

	err = db.InsertOne(ctx, "baz", res)
	if err != nil {
		return
	}

	return
}

func DeleteOne(ctx context.Context, db *database.MongoDB, tc *TestCase) (err error) {
	// TODO: find info for `options.Delete()` struct
	return db.DeleteOne(ctx, "baz", tc.id)
}
