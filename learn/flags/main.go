package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := run2(); err != nil {
		panic(err)
	}
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email int    `json:"email"`
}

func run2() (err error) {
	name, pass, _ := CmdFlags()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := newDatabase(ctx, fmt.Sprintf("mongodb://%s:%s@localhost:27017", *name, *pass), "company")
	if err != nil {
		return
	}
	defer db.Client().Disconnect(ctx)

	app := NewApp(db)

	if err = app.handleInsertUser(ctx, "./learn/flags/users.validation.json"); err != nil {
		return
	}

	if err = app.handleInsertUserTwo(ctx, "./learn/flags/users.validation.json"); err != nil {
		return
	}

	fmt.Println("success")
	return
}

// APP.GO
type app struct {
	db *database
}

func NewApp(db *database) *app {
	a := &app{db}
	return a
}

func (a *app) handleInsertUser(ctx context.Context, schema string) (err error) {
	type res struct {
		Name string `bson:"not_name"`
	}

	if err = a.db.SetValidator(ctx, "users", schema); err != nil {
		return
	}

	//inside closure
	r := &res{"One"}
	id, err := a.db.InsertOneUser(ctx, r)
	if err != nil {
		return
	}

	fmt.Println(id)
	return nil
}

func (a *app) handleInsertUserTwo(ctx context.Context, schema string) (err error) {
	type res struct {
		Name string `bson:"name"`
	}

	if err = a.db.SetValidator(ctx, "users", schema); err != nil {
		return
	}

	r := &res{"Two"}
	id, err := a.db.InsertOneUser(ctx, r)
	if err != nil {
		return
	}

	fmt.Println(id)
	return nil
}

// DATABASE.GO
type database struct {
	*mongo.Database
}

func newDatabase(ctx context.Context, uri, dbName string) (db *database, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	err = client.Connect(ctx)
	if err != nil {
		return
	}

	return &database{client.Database(dbName)}, nil
}

func (d *database) SetValidator(ctx context.Context, collMod, schema string) (err error) {
	doc, err := parseSchema(schema)
	if err != nil {
		return
	}

	return d.RunCommand(ctx, doc).Err()
}

func (d *database) InsertOneUser(ctx context.Context, v interface{}) (*primitive.ObjectID, error) {
	doc, err := toDoc(v)
	if err != nil {
		return nil, err
	}

	res, err := d.Collection("users").UpdateOne(ctx, doc, bson.D{{"$set", doc}}, options.Update().SetUpsert(true))
	if err != nil {
		return nil, err
	}

	if res.UpsertedID == nil {
		return nil, fmt.Errorf("user already exists in collection")
	}

	id := res.UpsertedID.(primitive.ObjectID)
	return &id, err
}
