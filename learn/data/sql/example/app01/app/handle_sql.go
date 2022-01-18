package app

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/topheruk/go/learn/fs/sql/example/app01/model"
	"github.com/topheruk/go/src/parse"
)

func (a *app) echo(message string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, message, http.StatusOK)
	}
}

func (a *app) handleCreateUser(query string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var dto model.DtoUser
		if err := a.decode(rw, r, &dto); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		u, err := a.db.InsertUser(r.Context(), query, &dto)
		if err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, u, http.StatusOK)
	}
}

func (a *app) handleSearchUser(query string) http.HandlerFunc {
	type response struct {
		ID uuid.UUID `json:"id"`
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		u, err := a.db.SearchUser(r.Context(), query, response{parse.GetID(r.Context())})
		if err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, u, http.StatusOK)
	}
}

func (a *app) handleSearchAll(query string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var uu []model.User
		if err := a.db.QueryiContext(r.Context(), query, &uu); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, uu, http.StatusOK)
	}
}
