package main

import (
	"net/http"

	t "github.com/topheruk/go/src/template"
)

type app struct {
	r *http.ServeMux
}

func (a *app) ServeHTTP(rw http.ResponseWriter, r *http.Request) { a.r.ServeHTTP(rw, r) }

func newApp() (a *app) {
	a = &app{
		r: http.NewServeMux(),
	}
	a.routes()
	return
}

func (a *app) routes() {
	a.r.HandleFunc("/", a.handleHome("example/net/htmlTemplaing/views/index.html"))
}

func (a *app) handleHome(path string) http.HandlerFunc {
	type data struct {
		Name string
	}

	exec, _ := t.Render(path)
	return func(rw http.ResponseWriter, r *http.Request) {
		exec(rw, r, data{"Kristopher"})
	}
}
