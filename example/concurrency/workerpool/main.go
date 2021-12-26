package main

import (
	"fmt"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() (err error) {
	a := newApp(10)
	sum := a.sum()
	fmt.Println(sum)

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
