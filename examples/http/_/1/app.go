package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/topheruk/go/src/encoding"
	"github.com/topheruk/go/src/template"
)

type app struct {
	m  *chi.Mux
	db *sqlx.DB
}

func newApp(db *sqlx.DB) *app {
	a := &app{
		m:  chi.NewMux(),
		db: db,
	}
	a.routes()
	return a
}

func (a *app) ServeHTTP(rw http.ResponseWriter, r *http.Request) { a.m.ServeHTTP(rw, r) }

func (a *app) respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	encoding.Respond(rw, r, data, status)
}

func (a *app) decode(rw http.ResponseWriter, r *http.Request, data interface{}) (err error) {
	return encoding.Decode(rw, r, data)
}

func (a *app) redirect(rw http.ResponseWriter, r *http.Request, url string, status int) {
	http.Redirect(rw, r, url, status)
}

func (a *app) render(path ...string) template.RenderFunc {
	render, err := template.Render(path...)
	if err != nil {
		panic(fmt.Errorf("parsing template error: %w", err))
	}
	return render
}
