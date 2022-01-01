package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/topheruk/go/src/parse"
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

// func (db *Database) SetValidator(ctx context.Context, schemaUri string) (err error) {
// 	doc, err := parse.BsonCmd(schemaUri)
// 	if err != nil {
// 		return
// 	}
// 	return db.RunCommand(ctx, doc).Err()
// }

func (db *Database) SetValidator(schemaUri string) (func(ctx context.Context) error, error) {
	doc, err := parse.BsonCmd(schemaUri)
	return func(ctx context.Context) error {
		return db.RunCommand(ctx, doc).Err()
	}, err
}

func (db *Database) Disconnect(ctx context.Context) error {
	return db.Client().Disconnect(ctx)
}
