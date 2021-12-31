package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	Name string
}

func main() {
	foo := &Foo{Name: "Foo"}

	Fn(foo)
}

func Fn(v interface{}) {
	type Bar struct {
		Name string
	}

	r := reflect.Indirect(reflect.ValueOf(v))
	fmt.Println(r.Field(0).String())
}
