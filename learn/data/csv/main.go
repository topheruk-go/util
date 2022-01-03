package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/topheruk/go/learn/data/csv/serde"
)

var filename = flag.String("f", "", "name given for .csv file")

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	fr, err := os.Open(fmt.Sprintf("learn/data/csv/data/%s.csv", *filename))
	if err != nil {
		return
	}
	defer fr.Close()

	fw, err := os.Create("learn/data/csv/data/output.csv")
	if err != nil {
		return
	}
	defer fw.Close()

	csv, _ := serde.NewCSV(fw, fr, &serde.Options{})

	// Reader
	var users []User
	for csv.Scan() {
		var u User
		if csv.Decode(&u) == io.EOF {
			break
		}
		users = append(users, u)
	}

	// Writer
	for _, u := range users {
		csv.Encode(u)
	}
	defer csv.Flush()

	return
}

type User struct {
	Name     string `csv:"name"`
	Age      int    `csv:"age"`
	Location string `csv:"location"`
}
