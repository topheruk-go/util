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
	sum, fact := doSomething(w.res)
	fmt.Println(sum, fact)
	// var y = w.sum()
	// var x = w.fact()

	// // fmt.Println(y)
	// fmt.Println(x)

	return
}

type worker struct {
	job, res chan int
}

func newWorker(size int) (w *worker) {
	w = &worker{make(chan int, size), make(chan int, size)}
	go w.spawn()
	return
}

func (w *worker) spawn() { spawn(w.job, w.res) }

// func (w *worker) sum() int  { return sum(w.res) }
func (w *worker) fact() int { return fact(w.res) }

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
			if x == 1 {
				fact = 1
			} else {
				fact *= x
			}
		}(tmp)
	}
	wg.Wait()
	return
}

func sum(res <-chan int) (sum int) {
	for x := 0; x < cap(res); x++ {
		sum += <-res
	}
	return
}

func fact(res <-chan int) (sum int) {
	for x := 0; x <= cap(res); x++ {
		if x == 1 {
			sum = 1
		} else {
			sum *= x
		}
	}
	return
}

type app struct {
	job, res chan int
}

func newApp(size int) (a *app) {
	a = &app{
		job: make(chan int, size),
		res: make(chan int, size),
	}
	go a.spawn()

	return
}

func (a *app) spawn() {
	defer close(a.job)
	for i := 0; i < cap(a.job); i++ {
		go a.work()
		a.job <- i
	}
}

func (a *app) work() {
	for j := range a.job {
		a.res <- j + 1
	}
}

func (a *app) sum() (sum int) {
	// TODO: the range synax is nicer
	for x := 0; x < cap(a.res); x++ {
		sum += <-a.res
	}
	return
}
