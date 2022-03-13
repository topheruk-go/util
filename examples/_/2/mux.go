package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

func (a *app) routes() {
	a.m.Use(middleware.Logger)

	a.m.HandleFunc("/", a.handleEcho("ping"))
	a.m.Handle("/favicon.ico", http.NotFoundHandler())

	a.m.Get("/laptop-loan", a.handleLaptopLoan("app/views/index.html"))
	// send blob and then unmarshal and add to database
	a.m.Post("/laptop-loan", a.handleLaptopLoanPost("app/tmp", "/api/v1"))

	a.m.Post("/api/v1/laptop-loan", a.handleAPILaptopLoan(`INSERT INTO laptop_loan VALUES (:id, :student_id, :start_date, :end_date, :tmp_path)`))
	a.m.Get("/api/v1/laptop-loan/", a.handleAPISearchLaptopLoan())
}

func (a *app) handleEcho(message string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, message, http.StatusOK)
	}
}

func (a *app) handleLaptopLoan(path ...string) http.HandlerFunc {
	type response struct {
		PostURL          string
		MinDate, MaxDate string
	}
	render := a.render(path...)
	return func(rw http.ResponseWriter, r *http.Request) {
		sd, ed := htmlFormattedTimeDuration(time.Now())
		resp := &response{PostURL: r.URL.Path, MinDate: sd, MaxDate: ed}

		render(rw, r, resp)
	}
}

func (a *app) handleLaptopLoanPost(dir, api string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		lf := &LoanFormDto{TmpPath: dir}
		if err := parseMultiPartForm(rw, r, lf); err != nil {
			a.respond(rw, r, err.Error(), http.StatusInternalServerError)
			return
		}
		url := parseURL(r, api+r.URL.Path)
		if err := postLoanForm(url, "application/json", lf); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}
		a.redirect(rw, r, r.URL.Path, http.StatusFound)
	}
}

func (a *app) handleAPILaptopLoan(query string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var dto LoanFormDto
		if err := a.decode(rw, r, &dto); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}
		if err := a.db.NamedQueryi(query, newLoanForm(dto)); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}
		a.respond(rw, r, dto, http.StatusOK)
	}
}

// TODO: search from database
func (a *app) handleAPISearchLaptopLoan() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, "todo", http.StatusOK)
	}
}
