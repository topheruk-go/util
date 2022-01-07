package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	t "github.com/topheruk/go/src/template"
	"golang.org/x/crypto/bcrypt"
)

type app struct {
	m  *chi.Mux
	db database
}

func App(db database) *app {
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
	a.m.Get("/sign-in-form", a.handleSignInForm("learn/auth/std-lib/views/sign-in.html"))
	a.m.Get("/sign-up-form", a.handleSignUpForm("learn/auth/std-lib/views/sign-up.html"))

	a.m.Get("/sign-in", a.handleSignIn("learn/auth/std-lib/views/sign-in.html"))
	a.m.Post("/sign-up", a.handleSignUp("learn/auth/std-lib/views/sign-up.html", "/sign-in-form"))
}

func (a *app) handleSignUp(filePath, redirectPath string) http.HandlerFunc {
	exec, _ := t.Render(filePath)

	return func(rw http.ResponseWriter, r *http.Request) {
		var email, password = r.FormValue("email"), r.FormValue("password")

		for k, v := range r.Form {
			fmt.Printf("key: %s; val: %v\n", k, v)
		}

		user, err := User(email, password)
		if err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		if err = a.db.post(r.Context(), user); err != nil {
			exec(rw, r, fmt.Sprintf("User with email %s already exists", email))
			return
		}

		http.Redirect(rw, r, redirectPath, http.StatusFound)
	}
}

func (a *app) handleSignIn(filePath string) http.HandlerFunc {
	exec, _ := t.Render(filePath)
	return func(rw http.ResponseWriter, r *http.Request) {
		var email, password = r.FormValue("email"), r.FormValue("password")
		user, err := a.db.get(r.Context(), email, password)
		if err != nil {
			exec(rw, r, err)
			return
		}

		a.respond(rw, r, user, http.StatusOK)
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
	post(ctx context.Context, u *user) error
	get(ctx context.Context, email, password string) (*user, error)
}

type service struct {
	users map[string]*user
}

func Service() *service {
	s := &service{
		users: map[string]*user{},
	}
	return s
}

func (s *service) post(ctx context.Context, u *user) error {
	if _, ok := s.users[u.Email]; ok {
		return fmt.Errorf("user %s already exists", u.Email)
	}

	s.users[u.Email] = u
	return nil
}

func (s *service) get(ctx context.Context, email, password string) (*user, error) {
	user, ok := s.users[email]
	if !ok {
		return nil, fmt.Errorf("no user exists for %s", email)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password, try again")
	}

	return user, nil
}

type user struct {
	Email    string `json:"email"`
	Password string `json:"-"`
}

func User(email, password string) (*user, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &user{email, string(hash)}, nil
}

type dtoUser struct {
	Email    string
	Password string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	fmt.Println("running... http://localhost:8000/sign-up-form")
	return http.ListenAndServe(":8000", App(Service()))
}

// GET 	/ landing page
// GET 	/sign-in sign-in form -> ok? user page else? sign-in-form w/ error
// GET  / anon page : view only
// GET 	/ [auth] free user page : above + can post, can view
// GET  / [auth] paid user page : above + can change details, can post, can view
// GET 	/ [auth] admin user page : above + ability to delete users
// GET 	/sign-up sign-up form -> ok? sign-page else? sign-up-form w/ error
