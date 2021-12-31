package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/topheruk/go/example/net/mongoDb/database"
)

type app struct {
	m  *chi.Mux
	db *database.Database
}

func (a *app) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// if r.RequestURI == "*" {
	// 	if r.ProtoAtLeast(1, 1) {
	// 		rw.Header().Set("Connection", "close")
	// 	}
	// 	rw.WriteHeader(StatusBadRequest)
	// 	return
	// }
	// h, _ := mux.Handler(r)
	// h.ServeHTTP(rw, r)

	a.m.ServeHTTP(rw, r)
}

func New(db *database.Database) (a *app) {
	a = &app{chi.NewMux(), db}
	a.routes()
	return
}

func (a *app) respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(rw).Encode(data)
		if err != nil {
			http.Error(rw, "Could not encode in json", status)
		}
	}
}

func (a *app) decode(rw http.ResponseWriter, r *http.Request, data interface{}) (err error) {
	if err = json.NewDecoder(r.Body).Decode(data); err != nil {
		return
	}

	return validator.New().Struct(data)
}
