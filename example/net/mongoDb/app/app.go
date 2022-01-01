package app

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/topheruk/go/src/parse"

	db "github.com/topheruk/go/example/net/mongoDb/database"
)

type app struct {
	m  *chi.Mux
	db *db.Database
}

func New(db *db.Database) *app {
	a := &app{chi.NewMux(), db}
	a.routes()
	return a
}

func (a *app) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	a.m.ServeHTTP(rw, r)
}

func (a *app) SetupValidation(ctx context.Context, urls ...string) error {
	for _, u := range urls {
		doc, err := parse.BsonCmd(u)
		if err != nil {
			return err
		}

		if err = a.db.RunCommand(ctx, doc).Err(); err != nil {
			return err
		}
	}
	return nil
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
	return json.NewDecoder(r.Body).Decode(data)
}
