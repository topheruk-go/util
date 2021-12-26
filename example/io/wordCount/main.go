package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var words map[string]int

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() (err error) {
	f, err := os.Open("example/wordCount/hamlet_gut.txt")
	if err != nil {
		return
	}
	defer f.Close()

	a := newApp(f)
	a.print()
	return
}

type app struct {
	sync.WaitGroup
	s     *bufio.Scanner
	found map[string]int
}

func newApp(f io.ReadCloser) (a *app) {
	a = &app{
		sync.WaitGroup{},
		bufio.NewScanner(f),
		make(map[string]int),
	}
	a.run()
	return a
}

func (a *app) run() {
	a.Add(1)
	defer a.Wait()
	go func() {
		defer a.Done()
		if err := a.tallyWords(); err != nil {
			fmt.Println(err.Error())
		}
	}()
}

func (a *app) add(word string, n int) {
	count, ok := a.found[word]
	if !ok {
		a.found[word] = n
	}

	a.found[word] = count + n
}

func (a *app) print() {
	for word, count := range a.found {
		if count > 1 {
			fmt.Printf("%d %s\n", count, word)
		}
	}
}

func (a *app) tallyWords() (err error) {
	a.s.Split(bufio.ScanWords)
	for a.s.Scan() {
		word := strings.ToLower(a.s.Text())
		a.add(word, 1)
	}

	return a.s.Err()
}
