package pkg

import "fmt"

type Bar struct {
	Name string
}

func Fn(v interface{}) {
	fmt.Println(v.(*Bar).Name)
}
