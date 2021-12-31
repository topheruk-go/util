package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type response struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func run() (err error) {

	encode, err := Validate("learn/jsonSchema/user.json")
	if err != nil {
		return
	}

	var v response
	r := strings.NewReader(`{"name":"John", "age":32 }`)
	if err = encode(r, &v); err != nil {
		return
	}
	fmt.Println(v.Name)

	return nil
}

func Validate(url string) (f func(r io.Reader, v interface{}) error, err error) {
	var (
		schema *jsonschema.Schema
		init   sync.Once
	)

	init.Do(func() { schema, err = jsonschema.Compile(url) })

	return func(r io.Reader, data interface{}) (err error) {
		if err = json.NewDecoder(r).Decode(data); err != nil {
			return
		}

		return schema.Validate(data)
	}, err
}
