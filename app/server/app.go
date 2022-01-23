package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	tmpl "github.com/topheruk/go/src/template"
)

type app struct {
	m *chi.Mux
}

func newApp() *app {
	a := &app{m: chi.NewMux()}
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
