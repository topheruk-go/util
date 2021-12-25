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
		go func(i int, ch chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- i * 2
		}(i, ch, &wg)
	}
	wg.Wait()
	close(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
