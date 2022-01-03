package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/topheruk/go/learn/data/csv/app"
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

	app, err := app.New(fw, fr, &app.Options{})
	if err != nil {
		return
	}

	// Reader
	var users []User
	for app.Scan() {
		var u User
		if err := app.Decode(&u); err == io.EOF {
			break
		}

		users = append(users, u)
	}

	// Writer
	// fmt.Println(len(users))
	for _, u := range users {
		if err = app.Encode(u); err != nil {
			return
		}
	}
	if err = app.Flush(); err != nil {
		return
	}

	// fmt.Println(buf.String())
	return
}

type User struct {
	Name     string `csv:"name"`
	Age      int    `csv:"age"`
	Location string `csv:"-"`
}
