package app01

import (
	"net/http"

	e "github.com/topheruk/go/src/encoding"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *App) routes() {
	a.m.Get("/ping", a.handlePing())
	a.m.Get("/echo", a.handleEcho("this is an example of an echo handler"))

	a.m.Post("/user", a.handleCreateUser("example/web/json/user.schema.json"))
	a.m.Get("/user/{id}", a.handleSearchUser())
}

func (a *App) handlePing() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		e.Respond(rw, r, "ping", http.StatusOK)
	}
}

func (a *App) handleEcho(message string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		e.Respond(rw, r, message, http.StatusOK)
	}
}

func (a *App) handleCreateUser(schema string) http.HandlerFunc {
	decode, err := e.Validator(schema)
	return func(rw http.ResponseWriter, r *http.Request) {
		var dto DtoUser
		err = decode(rw, r, &dto)
		if err != nil {
			e.Respond(rw, r, err, http.StatusBadRequest)
			return
		}

		err = a.c.InsertOne(r.Context(), &dto)
		if err != nil {
			e.Respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		e.Respond(rw, r, dto, http.StatusOK)
	}
}

func (a *App) handleSearchUser() http.HandlerFunc {
	var (
		oid primitive.ObjectID
		err error
	)

	return func(rw http.ResponseWriter, r *http.Request) {
		oid, err = a.OID(r)
		if err != nil {
			e.Respond(rw, r, err, http.StatusBadRequest)
			return
		}

		var user User
		err = a.c.FindOne(r.Context(), a.filterByID(oid), &user)
		if err != nil {
			e.Respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		e.Respond(rw, r, user, http.StatusOK)
	}
}
