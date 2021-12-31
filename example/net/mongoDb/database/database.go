package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	*mongo.Database
}

func New(ctx context.Context, uri, dbName string) (db *Database, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	err = client.Connect(ctx)
	if err != nil {
		return
	}

	return &Database{client.Database(dbName)}, nil
}

func (db *Database) Disconnect(ctx context.Context) error {
	return db.Client().Disconnect(ctx)
}
