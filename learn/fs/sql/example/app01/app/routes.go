package app

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/topheruk/go/learn/fs/sql/example/app01/model"
)

func (a *app) echo(msg string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, msg, http.StatusOK)
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
		uid := a.GetID(r.Context())

		u, err := a.db.SearchUser(r.Context(), query, response{uid})
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
		if err := a.db.Query(r.Context(), query, &uu); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, uu, http.StatusOK)
	}
}
