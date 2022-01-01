package app

import (
	"net/http"

	db "github.com/topheruk/go/test/net/foobar/database"
)

func (a *app) routes() {
	a.m.Get("/api/v1/foo/", a.echo("foo"))
	a.m.Get("/api/v1/bar/", a.echo("bar"))

	a.m.Post("/api/v1/foo/", a.foo())
	a.m.Post("/api/v1/bar/", a.bar())
}

func (a *app) echo(message string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, message, http.StatusOK)
	}
}

func (a *app) foo() http.HandlerFunc {
	type response struct {
		Value int `json:"value" bson:"value"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		var res response
		if err := a.decode(rw, r, &res); err != nil {
			a.respond(rw, r, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := a.db.InsertOne(r.Context(), "foo", res)
		if err != nil {
			a.respond(rw, r, err.Error(), http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, db.Foo{id, res.Value}, http.StatusOK)
	}
}

func (a *app) bar() http.HandlerFunc {
	type response struct {
		Value string `json:"value" bson:"value"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		var res response
		if err := a.decode(rw, r, &res); err != nil {
			a.respond(rw, r, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := a.db.InsertOne(r.Context(), "bar", res)
		if err != nil {
			a.respond(rw, r, err.Error(), http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, db.Bar{id, res.Value}, http.StatusOK)
	}
}
