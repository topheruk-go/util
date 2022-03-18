package database

import (
	"context"
	"fmt"

	"github.com/topheruk/go/x/parse"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Insert(ctx context.Context, collName string, v interface{}) (primitive.D, error)
}

type MongoDB struct {
	*mongo.Database
}

func NewMongoDB(ctx context.Context, uri, dbName string) (db *MongoDB, err error) {
	c, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	err = c.Connect(ctx)
	if err != nil {
		return
	}

	return &MongoDB{c.Database(dbName)}, err
}

func (db *MongoDB) SetupValidation(ctx context.Context, schemaUrls ...string) error {
	for _, u := range schemaUrls {
		cmd, err := parse.BsonCmd(u)
		if err != nil {
			return err
		}

		err = db.RunCommand(ctx, cmd).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *MongoDB) InsertOne(ctx context.Context, collName string, v interface{}) error {
	res, err := db.Collection(collName).UpdateOne(ctx, v, bson.D{{"$set", v}}, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	if res.UpsertedID == nil {
		return fmt.Errorf("user already exists in collection")
	}

	return nil
}

func (db *MongoDB) DeleteOne(ctx context.Context, collName string, id string) error {
	res, err := db.Collection(collName).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("no user was deleted; invalid id %s", id)
	}

	return nil
}
