package main

import (
	"log"
	"net/http"
)

func (a *app) routes() {
	// a.m.Use(middleware.Logger)
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
	type p struct {
		// Name  string `json:"name"`
		// Age   int    `json:"age"`
		// Email string `json:"email"`
		File []byte
		// StartDate *time.Time `json:"start_date"`
		// EndDate   *time.Time `json:"end_date"`
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{}
		if err := a.decode(rw, r, &resp); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}
		// resp["file"]
		log.Println(resp)
		a.respond(rw, r, "ok", http.StatusOK)
	}
}
