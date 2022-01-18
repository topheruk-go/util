package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (a *app) echo(msg string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, msg, http.StatusOK)
	}
}

func (a *app) handleCreateUser() http.HandlerFunc {
	var stmt = `INSERT INTO user VALUES (:id, :email, :password, :created_at)`
	return func(rw http.ResponseWriter, r *http.Request) {
		var dto DtoUser
		if err := a.decode(rw, r, &dto); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		u, err := a.db.InsertUser(r.Context(), stmt, &dto)
		if err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, u, http.StatusOK)
	}
}

func (a *app) handleSearchUser() http.HandlerFunc {
	type response struct {
		ID uuid.UUID `json:"id"`
	}
	var stmt = `SELECT * FROM user WHERE id = :id`
	return func(rw http.ResponseWriter, r *http.Request) {
		uid := a.GetID(r.Context())

		u, err := a.db.SearchUser(r.Context(), stmt, response{uid})
		if err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, u, http.StatusOK)
	}
}

func (a *app) handleSearchAll() http.HandlerFunc {
	var stmt = `SELECT * FROM user`
	return func(rw http.ResponseWriter, r *http.Request) {
		var uu []User
		if err := a.db.db.QueryiContext(r.Context(), stmt, &uu); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, uu, http.StatusOK)
	}
}
