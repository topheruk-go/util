package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type app struct {
	db Database
}

func New(db Database) (a *app) {
	a = &app{db: db}
	return
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() (err error) {
	cfg, err := env("learn/mongodb/one-to-many/.env")

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.uri))
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app := New(&TutorialDatabase{client.Database("tutorial")})

	if err = app.db.Connect(ctx); err != nil {
		return
	}
	defer app.db.Disconnect(ctx)

	if err = app.db.Ping(ctx); err != nil {
		return
	}

	err = app.db.Validate(ctx, "learn/mongodb/one-to-many/schema/")
	if err != nil {
		return
	}

	// Insert Podcast
	type podcast struct {
		ID     string   `bson:"_id,omitempty"`
		Title  string   `bson:"title,omitempty"`
		Author string   `bson:"author,omitempty"`
		Tags   []string `bson:"tags,omitempty"`
	}

	pod := podcast{
		ID:     "1",
		Title:  "The Polyglot Developer Podcast",
		Author: "Nic Raboy",
		Tags:   []string{"development", "programming", "coding"},
	}

	err = app.db.Insert(ctx, "podcasts", pod)
	if err != nil {
		return
	}

	type episode struct {
		ID          string `bson:"_id,omitempty"`
		PodcastID   string `bson:"podcast_id,omitempty"`
		Title       string `bson:"title,omitempty"`
		Description string `bson:"description,omitempty"`
		Duration    int32  `bson:"duration,omitempty"`
	}

	ep1 := episode{
		ID:          "1",
		PodcastID:   "1",
		Title:       "GraphQL for API Development",
		Description: "Learn about GraphQL from the co-creator of GraphQL, Lee Byron.",
		Duration:    25,
	}

	ep2 := episode{
		ID:          "2",
		PodcastID:   "1",
		Title:       "Progressive Web Application Development",
		Description: "Learn about PWA development with Tara Manicsic.",
		Duration:    32,
	}

	// Insert Episode
	err = app.db.InsertMany(ctx, "episodes", bson.A{ep1, ep2})
	if err != nil {
		return
	}

	return
}
