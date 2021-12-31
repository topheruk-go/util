package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	db "github.com/topheruk/go/example/net/mongoDb/database"
)

func (a *app) routes() {
	a.m.HandleFunc("/ping", a.ping)
	// FIXME: would like to still be able to have trailing slashes route
	a.m.Get("/api/v1/users/", a.findAllUsers())
	a.m.Get("/api/v1/users/{id}", a.findUser())
	a.m.Post("/api/v1/users/", a.createUser())
	a.m.Put("/api/v1/users/{id}", a.updateUser())
	a.m.Delete("/api/v1/users/{id}", a.deleteUser())
}

func (a *app) ping(rw http.ResponseWriter, r *http.Request) {
	a.respond(rw, r, "ping", http.StatusOK)
}

func (a *app) findAllUsers() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// I can assume this is almost impossible to fail
		users, err := a.db.FindManyUsers(r.Context(), bson.D{})
		if err != nil {
			a.respond(rw, r, nil, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, &users, http.StatusOK)
	}
}

func (a *app) findUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		id, err := primitive.ObjectIDFromHex(chi.URLParamFromCtx(r.Context(), "id"))
		if err != nil {
			a.respond(rw, r, nil, http.StatusBadRequest)
			return
		}

		user, err := a.db.FindUser(r.Context(), bson.M{"_id": id})
		if err != nil {
			a.respond(rw, r, nil, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, user, http.StatusOK)
	}
}

func (a *app) createUser() http.HandlerFunc {
	type response struct {
		Name string `bson:"name"`
		Age  int    `bson:"age" validate:"required"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		var res response

		if err := a.decode(rw, r, &res); err != nil {
			a.respond(rw, r, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO: This will fail if user already exists in Database
		id, err := a.db.InsertUser(r.Context(), res)
		if err != nil {
			a.respond(rw, r, nil, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, db.User{id, res.Name, res.Age}, http.StatusOK)
	}
}

func (a *app) updateUser() http.HandlerFunc {
	type response struct {
		Name string `bson:"name"`
		Age  int    `bson:"age" validate:"required"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		var res response

		if err := a.decode(rw, r, &res); err != nil {
			a.respond(rw, r, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := primitive.ObjectIDFromHex(chi.URLParamFromCtx(r.Context(), "id"))
		if err != nil {
			a.respond(rw, r, nil, http.StatusBadRequest)
			return
		}

		count, err := a.db.UpdateUser(r.Context(), id, res, false)
		if err != nil {
			a.respond(rw, r, nil, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, fmt.Sprintf("%d users updated", count), http.StatusOK)
	}
}

func (a *app) deleteUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		id, err := primitive.ObjectIDFromHex(chi.URLParamFromCtx(r.Context(), "id"))
		if err != nil {
			a.respond(rw, r, nil, http.StatusBadRequest)
			return
		}

		count, err := a.db.DeleteUser(r.Context(), bson.M{"_id": id})
		if err != nil {
			a.respond(rw, r, nil, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, fmt.Sprintf("%d users deleted", count), http.StatusOK)
	}
}
