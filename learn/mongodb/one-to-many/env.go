package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	uri  string
	addr string
}

func env(path string) (cfg *Config, err error) {
	err = godotenv.Load(path)
	if err != nil {
		return
	}

	return &Config{
		os.ExpandEnv("mongodb://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/"),
		os.ExpandEnv("$SERVER_HOST:$SERVER_PORT"),
	}, nil
}
