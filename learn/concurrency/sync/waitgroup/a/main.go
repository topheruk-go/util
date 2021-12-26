package main

import (
	"fmt"
	"sync"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() (err error) {
	var wg sync.WaitGroup
	var ch = make(chan int, 100)

	for i := 1; i <= cap(ch); i++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			ch <- c
		}(i)
	}
	wg.Wait()
	close(ch)

	var sum int
	for c := range ch {
		sum += c
	}

	fmt.Println(sum)
	return
}
