package app

import (
	"net/http"

	t "github.com/topheruk/go/src/template"
)

func (a *app) handleSignInForm(path string) http.HandlerFunc {
	exec, err := t.Render(path)
	return func(rw http.ResponseWriter, r *http.Request) {
		if err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		exec(rw, r, nil)
	}
}

func (a *app) handleSignUpForm(path string) http.HandlerFunc {
	exec, err := t.Render(path)
	return func(rw http.ResponseWriter, r *http.Request) {
		if err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		exec(rw, r, nil)
	}
}
