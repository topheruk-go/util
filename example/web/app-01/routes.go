package app01

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/topheruk/go/example/web/app-01/model"
	e "github.com/topheruk/go/src/encoding"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
		var dto model.DtoUser
		if err = decode(rw, r, &dto); err != nil {
			e.Respond(rw, r, err, http.StatusBadRequest)
			return
		}
		if err = a.db.Insert(r.Context(), &dto); err != nil {
			e.Respond(rw, r, err, http.StatusInternalServerError)
			return
		}
		e.Respond(rw, r, dto, http.StatusOK)
	}
}

func (a *App) handleSearchUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		id, err := primitive.ObjectIDFromHex(chi.URLParamFromCtx(r.Context(), "id"))
		if err != nil {
			e.Respond(rw, r, err, http.StatusBadRequest)
			return
		}
		user, err := a.db.Search(r.Context(), id)
		if err != nil {
			e.Respond(rw, r, err, http.StatusInternalServerError)
			return
		}
		e.Respond(rw, r, user, http.StatusOK)
	}
}

func (a *App) handleDeleteUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		id, err := primitive.ObjectIDFromHex(chi.URLParamFromCtx(r.Context(), "id"))
		if err != nil {
			e.Respond(rw, r, err, http.StatusBadRequest)
			return
		}
		if err = a.db.Delete(r.Context(), id); err != nil {
			e.Respond(rw, r, err, http.StatusInternalServerError)
			return
		}
		e.Respond(rw, r, id, http.StatusOK)
	}
}

func (a *App) handleUpdateUser(schema string) http.HandlerFunc {
	decode, err := e.Validator(schema)
	return func(rw http.ResponseWriter, r *http.Request) {
		var dto model.DtoUser
		if err = decode(rw, r, &dto); err != nil {
			e.Respond(rw, r, err, http.StatusBadRequest)
			return
		}
		id, err := primitive.ObjectIDFromHex(chi.URLParamFromCtx(r.Context(), "id"))
		if err != nil {
			e.Respond(rw, r, err, http.StatusBadRequest)
			return
		}
		user, err := a.db.Update(r.Context(), id, &dto)
		if err != nil {
			e.Respond(rw, r, err, http.StatusInternalServerError)
			return
		}
		e.Respond(rw, r, user, http.StatusOK)
	}
}
