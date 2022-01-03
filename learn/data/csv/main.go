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

	app, err := app.New(&bytes.Buffer{}, f, &app.AppOptions{})
	if err != nil {
		return
	}

	var users []User
	for app.Scan() {
		var u User
		if err := app.Decode(&u); err == io.EOF {
			break
		}

		users = append(users, u)
	}

	fmt.Println(users)

	return
}

type User struct {
	Name     string `csv:"name"`
	Age      int    `csv:"age"`
	Location string `csv:"location"`
}
