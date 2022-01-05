package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	db "github.com/topheruk/go/learn/docker/mongodb"
)

func init() {
	env := flag.String("env", "", "environment varaible file")
	flag.Parse()
	if err := godotenv.Load(*env); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	return db.Setup()
}
