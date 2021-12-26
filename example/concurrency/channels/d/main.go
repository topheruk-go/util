package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	inputChannel := make(chan int, 100)
	outputChannel := make(chan int, 100)

	//Create some workers
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go createWorker(inputChannel, outputChannel, &wg, j)
	}

	//Fill the input channl with data from a goroutine
	go run(inputChannel)

	go closeOnFinish(&wg, outputChannel)

	var sum int
	for i := range outputChannel {
		sum += i
	}
	fmt.Println(sum)
}

func createWorker(input, output chan int, wg *sync.WaitGroup, number int) {
	fmt.Printf("Worker %d created\n", number+1)
	for in := range input {
		output <- in * 100
	}
	// fmt.Printf("Worker %d closed\n", number+1)
	wg.Done()
}

//Fill upt the input channel
func run(input chan int) {
	for i := 0; i < 100; i++ {
		// fmt.Printf("Sending work %d\n", i)
		input <- i
	}
	close(input)
}

//just to make sure we close the channel once everything on the waitGroup is Done
func closeOnFinish(wg *sync.WaitGroup, output chan int) {
	wg.Wait()
	close(output)
}
