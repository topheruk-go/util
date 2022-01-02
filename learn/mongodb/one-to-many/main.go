package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
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
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() (err error) {
	cfg, err := env("learn/mongodb/one-to-many/.env")
	if err != nil {
		return
	}

	fmt.Println(cfg)

	return
}

type Config struct {
	Username string
	Password string
	Addr     string
}

func env(path string) (cfg *Config, err error) {
	err = godotenv.Load(path)
	if err != nil {
		return
	}

	return &Config{
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDRESS"),
	}, nil
}
