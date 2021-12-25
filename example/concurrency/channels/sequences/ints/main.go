package main

import "fmt"

func main() {
	c := run(sum, 10)
	println(c)
}

type seqFn func(n int, c chan int)

func run(f seqFn, size int) <-chan int {
	c := make(chan int, size)
	go f(cap(c), c)
	return c
}

func println(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
