package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

func (a *app) routes() {
	a.m.Use(middleware.Logger)
	a.m.Handle("/static/*", a.fileServer("/static/", "app/client/public"))
	a.m.Get("/*", a.handleIndex("app/views/index.html"))

	a.m.Post("/api/v1/user", a.handleApi())
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

func (a *app) handleApi() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var dto LoanDto
		if err := a.decode(rw, r, &dto); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		f, err := btof(dto.File)
		if err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}
		defer f.Close()

		l := newLoan(dto, f.Name())
		if err = a.db.Queryi(`INSERT INTO "loan" VALUES (:id,:name,:age,:path)`, l); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, "ok", http.StatusOK)
	}
}
