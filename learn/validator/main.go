package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type response struct {
	Name string `json:"name"`
	Age  int    `json:"age" validate:"required"`
}

func run() (err error) {

	encode, _ := Validate()

	var v response
	r := strings.NewReader(`{"name":"John", "age":32 }`)
	if err = encode(r, &v); err != nil {
		return
	}
	fmt.Println(v.Name)

	var u response
	r = strings.NewReader(`{"name":"John", "password":"jenny" }`)
	if err = encode(r, &u); err != nil {
		return
	}
	fmt.Println(u.Name)

	return nil
}

func Validate() (f func(r io.Reader, v interface{}) error, err error) {
	var (
		valid *validator.Validate
		init  sync.Once
	)

	init.Do(func() { valid = validator.New() })

	return func(r io.Reader, v interface{}) (err error) {
		if err = json.NewDecoder(r).Decode(v); err != nil {
			return
		}

		return valid.Struct(v)
	}, err
}
