package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/topheruk/go/src/parse"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database interface {
	Connect(context.Context) error
	Disconnect(context.Context) error
	Validate(ctx context.Context, schemaDir string) error
	Ping(context.Context) error
	Insert(ctx context.Context, collName string, v interface{}) error
	InsertMany(ctx context.Context, collName string, v []interface{}) error
}

type TutorialDatabase struct {
	*mongo.Database
}

func (db *TutorialDatabase) Validate(ctx context.Context, schemaDir string) error {
	f, err := os.Open(schemaDir)
	if err != nil {
		return err
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			doc, err := parse.BsonCmd(schemaDir + file.Name())
			if err != nil {
				return err
			}

			err = db.RunCommand(ctx, doc).Err()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (db *TutorialDatabase) Connect(ctx context.Context) error    { return db.Client().Connect(ctx) }
func (db *TutorialDatabase) Disconnect(ctx context.Context) error { return db.Client().Disconnect(ctx) }
func (db *TutorialDatabase) Ping(ctx context.Context) error {
	return db.Client().Ping(ctx, readpref.Primary())
}

func (db *TutorialDatabase) Insert(ctx context.Context, collName string, v interface{}) error {
	_, err := db.Collection(collName).InsertOne(ctx, v)
	if err != nil {
		return err
	}

	return nil
}

func (db *TutorialDatabase) InsertMany(ctx context.Context, collName string, v []interface{}) error {
	_, err := db.Collection(collName).InsertMany(ctx, v)
	if err != nil {
		return err
	}

	return nil
}

type Podcast struct {
	ID     string   `bson:"_id,omitempty"`
	Title  string   `bson:"title,omitempty"`
	Author string   `bson:"author,omitempty"`
	Tags   []string `bson:"tags,omitempty"`
}

type Episode struct {
	ID          string `bson:"_id,omitempty"`
	PodcastID   string `bson:"podcast_id,omitempty"`
	Title       string `bson:"title,omitempty"`
	Description string `bson:"description,omitempty"`
	Duration    int32  `bson:"duration,omitempty"`
}
