package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type User struct {
	Id   primitive.ObjectID `json:"id"`
	Name string             `json:"name"`
	Age  int                `json:"age"`
}

type response struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func run() (err error) {
	user, pass, _ := CmdFlags()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(
		// uri,
		fmt.Sprintf("mongodb://%s:%s@localhost:27017/", *user, *pass),
	))

	if err != nil {
		return
	}

	err = client.Connect(ctx)
	if err != nil {
		return
	}
	defer client.Disconnect(ctx)

	fmt.Println("connected")

	db := client.Database("company")

	// Create User
	response := &struct {
		Name string `bson:"name"`
		Age  int    `bson:"age"`
	}{"Ysidssn", 23}

	err = Validate(ctx, db, "./learn/flags/users.validation.json", response)
	if err != nil {
		return
	}

	upt, err := InsertOne(ctx, db, response)
	if err != nil {
		return
	}

	fmt.Println(upt)

	fmt.Println("successful")
	return
}

func CmdFlags() (username, password *string, address *int) {
	username = flag.String("u", "", "database client username:password")
	password = flag.String("p", "", "database client password")
	address = flag.Int("a", 8000, "server port number")

	flag.Parse()
	return
}

// FIXME: this works(?) but the command can fail at times
// need to understand what those cases are
func InsertOne(ctx context.Context, db *mongo.Database, v interface{}) (*mongo.UpdateResult, error) {
	b, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}

	var doc bson.D
	err = bson.Unmarshal(b, &doc)
	if err != nil {
		return nil, err
	}

	res, err := db.Collection("users").UpdateOne(ctx, doc, bson.D{{"$set", doc}}, options.Update().SetUpsert(true))
	if err != nil {
		return nil, err
	}

	return res, err
}

// TODO: add to main project as I think I undersatnd now
func Validate(ctx context.Context, db *mongo.Database, url string, v interface{}) (err error) {
	doc, err := parseJsonFile(url)
	if err != nil {
		return
	}

	return db.RunCommand(ctx, doc).Decode(v)
}

// FIXME: messy but works
func parseJsonFile(url string) (doc bson.D, err error) {
	content, err := ioutil.ReadFile(url)
	if err != nil {
		return
	}

	var tmp interface{}
	if err = json.Unmarshal(content, &tmp); err != nil {
		return
	}

	vb, err := bson.Marshal(tmp)
	if err != nil {
		return
	}

	err = bson.Unmarshal(vb, &doc)
	return
}
