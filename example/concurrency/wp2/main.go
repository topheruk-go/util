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
	w := newWorker(10)
	sum, fact := w.doSomething()
	fmt.Println(sum, fact)

	return
}

type worker struct {
	sync.WaitGroup
	job, res chan int
}

func newWorker(size int) (w *worker) {
	w = &worker{sync.WaitGroup{}, make(chan int, size), make(chan int, size)}
	go w.spawn()
	return
}

func (w *worker) spawn() { spawn(w.job, w.res) }
func spawn(job chan int, res chan<- int) {
	defer close(job)
	for i := 0; i < cap(job); i++ {
		go work(job, res)
		job <- i
	}
}

func work(job <-chan int, res chan<- int) {
	for j := range job {
		res <- j + 1
	}
}

func (w *worker) doSomething() (sum, fact int) {
	return doSomething(w.res)
}

func doSomething(res <-chan int) (sum, fact int) {
	var wg sync.WaitGroup
	for x := 0; x < cap(res); x++ {
		wg.Add(2)
		tmp := <-res

		go func(v int) {
			defer wg.Done()
			sum += v
		}(tmp)

		go func(v int) {
			defer wg.Done()
			fact -= v
		}(tmp)

	}
	wg.Wait()
	return
}
