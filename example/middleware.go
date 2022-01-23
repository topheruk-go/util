package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

type Middleware func(hf http.HandlerFunc) http.HandlerFunc

func Chain(hf http.HandlerFunc, mws ...Middleware) http.HandlerFunc {
	for _, m := range mws {
		hf = m(hf)
	}
	return hf
}

func ServePrefix(prefix string) Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, prefix) {
				hf(rw, r)
			}
		}
	}
}

func ServeHTML(hf http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, ".html") {
			hf(rw, r)
		}
	}
}

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() { log.Printf("%s %v", r.URL.Path, time.Since(start)) }()
		f(w, r)
	}
}
