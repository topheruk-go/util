package main

import (
	"fmt"
	"os"

	"github.com/topheruk/go/learn/csv/app"
)

// Find unique values from csv tables
// command list

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() (err error) {
	f, err := os.Open("learn/csv/data/terms.csv")
	if err != nil {
		return
	}

	app := app.New(f)

	for app.Scan() {
		for value := range app.Record() {
			fmt.Printf("%s ", app.Record()[value])
		}
	}

	_, err = app.All()
	if err != nil {
		return
	}

	return
}
