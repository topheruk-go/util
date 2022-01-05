package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	db "github.com/topheruk/go/learn/docker/database"
)

type App struct {
	m  *chi.Mux
	db db.Database
}

func New(d db.Database) *App {
	a := &App{chi.NewMux(), d}
	a.routes()
	return a
}

func (a *App) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	a.m.ServeHTTP(rw, r)
}

func (a *App) respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(rw).Encode(data)
		if err != nil {
			http.Error(rw, "Could not encode in json", status)
		}
	}
}

func (a *App) decode(rw http.ResponseWriter, r *http.Request, data interface{}) (err error) {
	return json.NewDecoder(r.Body).Decode(data)
}

func (a *App) routes() {
	a.m.HandleFunc("/ping", a.handleEcho("ping"))

	a.m.Get("/users/", a.handleGetAll())
	a.m.Get("/users/{id}", a.handleGetUser())

	a.m.Post("/users", a.handleSetUser())
}
