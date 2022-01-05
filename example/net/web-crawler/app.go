package main

import (
	"net/http"
	"sync"
)

type app struct {
	sync.WaitGroup
	urls  []string
	items map[string]item
}

func (a *app) run() {
	for _, u := range a.urls {
		a.Add(1)
		go a.crawl(u)
	}
	a.Wait()
}

func (a *app) crawl(u string) {
	defer a.Done()

	r, err := http.Get(u)
	if err != nil {
		return
	}
	defer r.Body.Close()
	l := lex(r.Body, 10)

	// FIXME: this needs to be in a go routine
	for item := range l.items {
		if _, ok := a.items[item.url.String()]; !ok {
			a.items[item.url.String()] = item
		}
	}
}
