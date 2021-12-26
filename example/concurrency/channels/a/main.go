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
	var job = make(chan int, 1000)
	var res = make(chan int, cap(job))

	for j := 1; j <= cap(job); j++ {
		wg.Add(1)
		go func(job <-chan int) {
			defer wg.Done()
			for j := range job {
				res <- j
			}
		}(job)
		job <- j
	}
	close(job)

	wg.Wait()
	close(res)

	var sum int
	for r := range res {
		sum += r
	}

	fmt.Println(sum)
	return
}
