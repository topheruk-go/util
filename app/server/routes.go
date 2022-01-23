package main

import (
	"net/http"
)

func (a *app) routes() {
	// a.m.Use(middleware.Logger)
	a.m.Handle("/static/*", a.fileServer("/static/", "app/client/public"))
	a.m.Get("/*", a.handleIndex("app/views/index.html"))
}

func (a *app) fileServer(prefix, dir string) http.Handler {
	return http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))
}

func (a *app) handleIndex(filenames ...string) http.HandlerFunc {
	type p struct {
		Title string
	}
	render := a.render(filenames...)
	return func(rw http.ResponseWriter, r *http.Request) {
		render(rw, r, &p{Title: "Basic Svelte App"})
	}
}
