package main

import (
	"flag"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id, omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

type Episodes struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PodcastID   primitive.ObjectID `bson:"podcast_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() (err error) {

	return
}
