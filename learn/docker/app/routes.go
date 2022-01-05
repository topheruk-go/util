package app

import "net/http"

func (a *App) handleEcho(message string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, message, http.StatusOK)
	}
}

func (a *App) handleSetUser() http.HandlerFunc {
	type response struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		var dto response

		if err := a.decode(rw, r, &dto); err != nil {
			a.respond(rw, r, err, http.StatusBadRequest)
			return
		}

		a.respond(rw, r, "todo", http.StatusOK)
	}
}

func (a *App) handleGetUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, "todo", http.StatusOK)
	}
}

func (a *App) handleGetAll() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, "todo", http.StatusOK)
	}
}

func (a *App) handleDelUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, "todo", http.StatusOK)
	}
}
