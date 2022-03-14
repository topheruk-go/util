package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/topheruk/go/src/database/sqli"
	"github.com/topheruk/go/src/encoding"
	"github.com/topheruk/go/src/template"
)

type app struct {
	m  *chi.Mux
	db *sqli.DB
}

func newApp(db *sqli.DB) *app {
	a := &app{
		m:  chi.NewMux(),
		db: db,
	}
	a.routes()
	return a
}

func (a *app) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	a.m.ServeHTTP(rw, r)
}

// panic if error
func (a *app) render(filenames ...string) template.RenderFunc {
	render, err := template.Render(filenames...)
	if err != nil {
		panic(err)
	}
	return render
}

func (a *app) respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	encoding.Respond(rw, r, data, status)
}

func (a *app) decode(rw http.ResponseWriter, r *http.Request, data interface{}) error {
	return encoding.Decode(rw, r, data)
}

func (a *app) redirect(rw http.ResponseWriter, r *http.Request, url string, status int) {
	http.Redirect(rw, r, url, status)
}
