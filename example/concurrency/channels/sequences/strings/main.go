package main

import (
	"fmt"
	"sync"
)

func main() {
	c := run(char, "hello")
	println(c)
}

type seqFn func(s string, c chan string)

func run(f seqFn, s string) <-chan string {
	c := make(chan string, len(s))

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		defer wg.Done()
		go func() {
			f(s, c)
		}()
	}
	wg.Wait()

	return c
}

func println(c <-chan string) {
	for i := range c {
		fmt.Println(i)
	}
}
