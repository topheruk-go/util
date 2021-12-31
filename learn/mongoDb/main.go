package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	username = "topheruk"
	password = "T^*G7!Pf"
	host     = "192.168.1.173"
	port     = 27017
	uri      = fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, port)
)

func main() {
	if err := run(); err != nil {
		return
	}
}

func run() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	err = client.Connect(ctx)
	if err != nil {
		return
	}
	defer client.Disconnect(ctx)

	db := client.Database("company").Collection("users")

	var doc bson.D
	if err = db.FindOne(ctx, bson.M{}).Decode(&doc); err != nil {
		return
	}

	fmt.Printf("%v", doc)

	return
}
