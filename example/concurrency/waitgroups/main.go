package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var ch = make(chan int, 100)

	for i := 0; i < cap(ch); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- i + 1
		}(i)
	}

	wg.Wait()
	close(ch)

	var sum int = 0
	for i := range ch {
		sum += i
	}

	fmt.Println(sum)
}
