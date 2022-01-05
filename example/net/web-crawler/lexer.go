package main

import (
	"io"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

type lexer struct {
	*html.Tokenizer
	items chan item
}

func lex(r io.ReadCloser, buf int) (l *lexer) {
	l = &lexer{
		html.NewTokenizer(r),
		make(chan item, buf),
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
