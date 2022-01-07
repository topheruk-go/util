package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	t "github.com/topheruk/go/src/template"
)

type app struct {
	m  *chi.Mux
	db database
}

func New(db database) *app {
	a := &app{
		db: db,
		m:  chi.NewMux(),
	}
	a.routes()
	return a
}

func (a *app) ServeHTTP(rw http.ResponseWriter, r *http.Request) { a.m.ServeHTTP(rw, r) }

func (a *app) respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(rw).Encode(data)
		if err != nil {
			http.Error(rw, "Could not encode in json", status)
		}
	}
}

func (a *app) decode(rw http.ResponseWriter, r *http.Request, data interface{}) (err error) {
	return json.NewDecoder(r.Body).Decode(data)
}

func (a *app) routes() {
	a.m.Get("/sign-in-form", a.handleSignInForm("learn/auth/views/sign-in.html"))
	a.m.Get("/sign-up-form", a.handleSignUpForm("learn/auth/views/sign-up.html"))

	a.m.Get("/sign-in", a.ServeHTTP)
	a.m.Post("/sign-up", a.ServeHTTP)
}

func (a *app) handleSignUp() http.HandlerFunc {
	type response struct {
		email    string
		password string
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		user := &response{email: r.FormValue("email"), password: r.FormValue("password")}

		a.respond(rw, r, "POST / sign-up", http.StatusInternalServerError)
	}
}

func (a *app) handleSignIn() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, "GET / sign-in", http.StatusInternalServerError)
	}
}

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

type database interface {
	insertUser(ctx context.Context, v interface{}) error
	getUser(ctx context.Context, v interface{}) error
}

type service struct {
}

func (s *service) insertUser(ctx context.Context, v interface{}) error {

}

type user struct {
	Email    string `json:"email"`
	Password string `json:"-"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() (err error) {

	return
}
