package main

import (
	"fmt"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() (err error) {
	fmt.Println("Listening to http://localhost:8000/")
	return http.ListenAndServe(":8000", newApp())
}
