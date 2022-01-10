package app01

import (
	"log"
	"net/http"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/topheruk/go/example/web/app-01/model"
)

func (a *App) handlePing() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, "ping", http.StatusOK)
	}
}

func (a *App) handleEcho(message string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, message, http.StatusOK)
	}
}

func (a *App) handleCreateUser(url string) http.HandlerFunc {
	schema, err := jsonschema.Compile(url)
	return func(rw http.ResponseWriter, r *http.Request) {
		var dto model.DtoUser
		if err = a.decode(rw, r, &dto); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		if err = schema.Validate(dto); err != nil {

		}

		// FIXME: create a logger that logs the response of each handler
		log.Println(dto)
		a.respond(rw, r, dto, http.StatusOK)
	}
}
