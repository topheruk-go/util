package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() (err error) {
	if len(os.Args) == 1 {
		return errors.New("please pass in a url as an argument")
	}

	r, err := http.Get(os.Args[1])
	if err != nil {
		return
	}
	defer r.Body.Close()
	l := lex(r.Body)

	for item := range l.items {
		fmt.Println(item.href)
	}

	return
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
	l.items <- item{href}
}

type item struct {
	// TODO: try `type URL` for href
	href string
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
		case strings.HasPrefix(a.Val, "/study/courses/"):
			l.emit(a.Val)
			return lexStartTag
		}
	}

	return lexStartTag
}
