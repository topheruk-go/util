package database

import (
	"context"
	"fmt"

	"github.com/topheruk/go/src/parse"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	*mongo.Database
	// cmds map[string]
}

func New(ctx context.Context, uri, dbName string) (db *Database, err error) {
	c, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	err = c.Connect(ctx)
	if err != nil {
		return
	}

	return &Database{c.Database(dbName)}, err
}

func (db *Database) SetupValidation(ctx context.Context, schemaUrls ...string) error {
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

// func (db *Database) SetValidator(ctx context.Context, schemaUri string) error {
// 	doc, err := parse.BsonCmd(schemaUri)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println(doc)
// 	// return func(ctx context.Context) error {
// 	// 	}, err
// 	return db.RunCommand(ctx, doc).Err()
// }

func (db *Database) InsertOne(ctx context.Context, collName string, v interface{}) (primitive.ObjectID, error) {
	doc, err := parse.ToDoc(v) //may not be needed
	if err != nil {
		return primitive.NilObjectID, err
	}

	res, err := db.Collection(collName).UpdateOne(ctx, doc, bson.D{{"$set", doc}}, options.Update().SetUpsert(true))
	if err != nil {
		return primitive.NilObjectID, err
	}

	if res.UpsertedID == nil {
		return primitive.NilObjectID, fmt.Errorf("user already exists in collection")
	}

	id := res.UpsertedID.(primitive.ObjectID)
	return id, err
}
