package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/topheruk/go/learn/mongodb/canvas/v1/model"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	var humanities model.Account
	var english model.Account
	var spanish model.Account

	if err = decode(`{
		"id":"A001",
		"status":"active",
		"name":"Humanities"
	}`, &humanities); err != nil {
		return
	}

	if err = decode(`{
		"id":"A002",
		"parent_account_id":"A001",
		"status":"active",
		"name":"English"
	}`, &english); err != nil {
		return
	}

	if err = decode(`{
		"id":"A003",
		"parent_account_id":"A001",
		"status":"deleted",
		"name":"Spanish"
	}`, &spanish); err != nil {
		return
	}

	for _, account := range []interface{}{humanities, english, spanish} {
		respond(account)
	}

	return
}

func decode(content string, v interface{}) error {
	return json.NewDecoder(strings.NewReader(content)).Decode(&v)
}

func respond(v interface{}) error {
	return json.NewEncoder(os.Stdout).Encode(v)
}
