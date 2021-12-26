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
	a := newApp(100)

	var sum int
	for r := range a.w.r {
		sum += r.value
	}

	fmt.Println(sum)
	return
}

type app struct {
	sync.WaitGroup

	s int
	w *work
}

func newApp(size int) (a *app) {
	a = &app{sync.WaitGroup{}, size, newWork(size)}
	go a.run()
	return
}

func (a *app) run() {
	defer a.close()
	for j := 0; j < 2; j++ {
		a.Add(1)
		go a.put(j)
	}

	go a.add()
}

func (a *app) add() {
	defer close(a.w.j)
	for i := 1; i <= a.s; i++ {
		a.w.j <- job{i}
	}
}

func (a *app) put(id int) {
	defer a.Done()
	fmt.Printf("Worker %d created\n", id+1)
	for j := range a.w.j {
		a.w.r <- result{j.value * 100}
	}
	fmt.Printf("Worker %d closed\n", id+1)
}

func (a *app) close() {
	a.Wait()
	close(a.w.r)
}

type job struct {
	value int
}

type result struct {
	value int
}

type work struct {
	j chan job
	r chan result
}

func newWork(size int) (w *work) {
	return &work{make(chan job, size), make(chan result, size)}
}
