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
	wk := work(1, 10)

	var sum int
	for i := 0; i < wk.siz; i++ {
		r := wk.next()
		sum += r.value
	}

	fmt.Println(sum)
	return
}

type worker struct {
	sync.WaitGroup

	siz int
	job chan job
	res chan result
}

type job struct {
	value int
}

type result struct {
	value int
}

func work(size, buf int) (w *worker) {
	w = &worker{
		sync.WaitGroup{},
		size,
		make(chan job, buf),
		make(chan result),
	}

	go w.run()
	return
}

func (w *worker) run() {
	defer w.wait()
	go w.push()
	for j := 0; j < w.siz; j++ {
		w.Add(1)
		go w.worker(j)
	}
}

func (w *worker) worker(id int) {
	defer w.Done()
	for j := range w.job {
		w.res <- result{j.value}
	}
}

func (w *worker) push() {
	defer close(w.job)
	for i := 1; i <= cap(w.job); i++ {
		w.job <- job{i}
	}
}

func (w *worker) next() result {
	return <-w.res
}

func (w *worker) wait() {
	w.Wait()
	close(w.res)
}
