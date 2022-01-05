package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Setup() (err error) {
	uri := os.ExpandEnv("mongodb://$MONGO_DB_USERNAME:$MONGO_DB_PASSWORD@$MONGO_DB_HOST:$MONGO_DB_PORT")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		return
	}
	defer client.Disconnect(ctx)

	names, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return
	}

	for _, n := range names {
		fmt.Printf("%s\n", n)
	}

	return
}
