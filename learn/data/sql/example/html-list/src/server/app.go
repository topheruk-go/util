package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/topheruk/go/src/database/sqli"
	e "github.com/topheruk/go/src/encoding"
)

type app struct {
	m  *chi.Mux
	db *sqli.DB
}

func (a *app) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	a.m.ServeHTTP(rw, r)
}

func (a *app) respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	e.Respond(rw, r, data, status)
}

func (a *app) decode(rw http.ResponseWriter, r *http.Request, data interface{}) (err error) {
	return e.Decode(rw, r, data)
}

func newApp(db *sqli.DB) *app {
	a := &app{
		m:  chi.NewMux(),
		db: db,
	}
	go a.routes()
	return a
}

func (a *app) routes() {
	a.m.HandleFunc("/index.html", a.handleIndex("html-list/src/views/index.html"))

	a.m.Get("/account/", a.handleGetAll())

	a.m.Get("/ping", a.handleEcho("ping"))
}
