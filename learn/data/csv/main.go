package main

import (
	"bytes"
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
	f, err := os.Open(fmt.Sprintf("learn/data/csv/data/%s.csv", *filename))
	if err != nil {
		return
	}
	defer f.Close()

	var buf bytes.Buffer
	app, err := app.New(&buf, f, &app.AppOptions{})
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
	for _, u := range users {
		if err = app.Encode(u); err != nil {
			return
		}
	}

	if err = app.Flush(); err != nil {
		return
	}

	fmt.Println(buf.String())
	return
}

type User struct {
	Name     string `csv:"name"`
	Age      int    `csv:"age"`
	Location string `csv:"location"`
}
