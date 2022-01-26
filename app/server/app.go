package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/topheruk/go/src/database/sqli"
	"github.com/topheruk/go/src/encoding"
	tmpl "github.com/topheruk/go/src/template"
)

type app struct {
	m  *chi.Mux
	db *sqli.DB
}

func newApp(db *sqli.DB) *app {
	a := &app{m: chi.NewMux(), db: db}
	a.routes()
	return a
}

func (a *app) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	a.m.ServeHTTP(rw, r)
}

func (a *app) render(filenames ...string) tmpl.RenderFunc {
	render, err := tmpl.Render(filenames...)
	if err != nil {
		panic(fmt.Errorf("error parsing template: %w", err))
	}
	return render
}

func (a *app) decode(rw http.ResponseWriter, r *http.Request, data interface{}) error {
	return encoding.Decode(rw, r, data)
}

func (a *app) respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	encoding.Respond(rw, r, data, status)
}
