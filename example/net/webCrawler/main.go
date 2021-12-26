package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() (err error) {
	if len(os.Args) == 1 {
		return errors.New("please pass in a url(s) as an argument")
	}

	var wg sync.WaitGroup
	var unique = map[string]*url.URL{}

	for _, u := range os.Args[1:] {
		wg.Add(1)
		go crawl(u, unique, &wg)
	}
	wg.Wait()

	for url := range unique {
		fmt.Println(url)
	}

	return
}

func crawl(u string, unique map[string]*url.URL, wg *sync.WaitGroup) {
	defer wg.Done()

	r, err := http.Get(u)
	if err != nil {
		return
	}
	defer r.Body.Close()
	l := lex(r.Body)

	// FIXME: this needs to be in a go routine
	for item := range l.items {
		if _, ok := unique[item.url.String()]; !ok {
			unique[item.url.String()] = item.url
		}
	}
}

type lexer struct {
	*html.Tokenizer
	items chan item
}

func lex(r io.ReadCloser) (l *lexer) {
	l = &lexer{
		html.NewTokenizer(r),
		make(chan item),
	}
	go l.run()
	return
}

func (l *lexer) run() {
	defer close(l.items)
	for state := lexRoot; state != nil; {
		state = state(l)
	}
}

func (l *lexer) emit(href string) {
	u, _ := url.Parse(href)
	l.items <- item{u}
}

type item struct {
	url *url.URL
}

type stateFn func(l *lexer) stateFn

func lexRoot(l *lexer) stateFn {
	switch typ := l.Next(); typ {
	case html.ErrorToken:
		return nil
	case html.StartTagToken:
		return lexStartTag
	default:
		return lexRoot
	}
}

func lexStartTag(l *lexer) stateFn {
	switch t := l.Token(); t.Data {
	case "a":
		return lexAnchorTag(l, &t)
	default:
		return lexRoot
	}
}

func lexAnchorTag(l *lexer, t *html.Token) stateFn {
	for _, a := range t.Attr {
		switch {
		case strings.HasPrefix(a.Val, "https://"):
			l.emit(a.Val)
			return lexStartTag
		}
	}

	return lexStartTag
}
