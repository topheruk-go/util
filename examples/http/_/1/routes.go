package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

func (a *app) routes() {
	a.m.Use(middleware.Logger)
	a.m.Handle("/favicon.ico", http.NotFoundHandler())

	a.m.Get("/*", a.getIndex("example/index.html"))

	a.m.Get("/laptop-loan", a.getLaptopLoan("example/laptop-loan.html"))

	a.m.Post("/laptop-loan/success", a.getLaptopLoanSuccess("example/laptop-loan-success.html"))
	a.m.Post("/laptop-loan/error", a.getLaptopLoanError("example/laptop-loan-error.html"))

	a.m.Post("/laptop-loan", a.postLaptopLoan("/laptop-loan", "/api/v1/laptop-loan"))

	a.m.Post("/api/v1/laptop-loan", a.postApiLaptopLoan())
}

// GET /
//
// Home page
func (a *app) getIndex(path string) http.HandlerFunc {
	render := a.render(path)
	return func(rw http.ResponseWriter, r *http.Request) {
		render(rw, r, nil)
	}
}

// GET /laptop-loan
//
// Html page where form is located
func (a *app) getLaptopLoan(path string) http.HandlerFunc {
	type response struct {
		Success bool
	}
	render := a.render(path)
	return func(rw http.ResponseWriter, r *http.Request) {
		var resp response
		if err := a.decode(rw, r, &resp); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		render(rw, r, nil)
	}
}

// GET /laptop-loan-success
//
// If successful
func (a *app) getLaptopLoanSuccess(path string) http.HandlerFunc {
	render := a.render(path)
	return func(rw http.ResponseWriter, r *http.Request) {
		render(rw, r, nil)
	}
}

// GET /laptop-loan-error
//
// If unsuccessful
func (a *app) getLaptopLoanError(path string) http.HandlerFunc {
	render := a.render(path)
	return func(rw http.ResponseWriter, r *http.Request) {
		render(rw, r, nil)
	}
}

// POST /laptop-loan
//
// Form values transformed into JSON.
func (a *app) postLaptopLoan(redirUrl string, postURL string) http.HandlerFunc {
	type response struct {
		Todo string
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		postURL = parseURL(r, postURL)
		// TODO: resolve form data
		var todo = &response{Todo: "todo"}
		_, err := postRequest(postURL, "application/json", &todo)
		if err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		a.redirect(rw, r, redirUrl, http.StatusFound)
	}
}

// POST /api/v1/laptop-loan
//
// Insert to db
func (a *app) postApiLaptopLoan() http.HandlerFunc {
	type response struct {
		Todo string
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		var resp response
		if err := a.decode(rw, r, &resp); err != nil {
			a.respond(rw, r, nil, http.StatusInternalServerError)
			// TODO: redirect to failure page
			return
		}

		a.respond(rw, r, &resp, http.StatusOK)
	}
}
