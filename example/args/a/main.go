package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() (err error) {
	fetch(os.Args[1:])

	return
}

func fetch(urls []string) error {
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			return err
		}

		fmt.Println(res.StatusCode)
	}

	return nil
}
