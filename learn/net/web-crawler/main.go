package main

import (
	"errors"
	"fmt"
	"os"
)

// go run ./example/net/webCrawler/ https://schier.co https://insomnia.rest

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() (err error) {
	if len(os.Args) == 1 {
		return errors.New("please pass in a url(s) as an argument")
	}

	a := &app{
		urls:  os.Args[1:],
		items: map[string]item{},
	}

	a.run()

	for _, i := range a.items {
		fmt.Printf("%v\n", i)
	}

	return
}
