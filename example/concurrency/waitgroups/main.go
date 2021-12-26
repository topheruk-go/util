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
	a := newApp(10)
	sum := a.sum()
	fmt.Println(sum)
	return
}

type app struct {
	sync.WaitGroup
	ch chan int
}

func newApp(size int) (a *app) {
	a = &app{
		WaitGroup: sync.WaitGroup{},
		ch:        make(chan int, size),
	}
	go a.spawn()
	return
}

func (a *app) spawn() {
	defer close(a.ch)
	for i := 0; i < cap(a.ch); i++ {
		a.Add(1)
		go func(i int) {
			defer a.Done()
			a.ch <- i + 1
		}(i)
	}
	a.Wait()
}

func (a *app) sum() (sum int) {
	for c := range a.ch {
		sum += c
	}
	return
}
