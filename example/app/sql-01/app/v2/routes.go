package app

import (
	"net/http"
)

func (a *App) routes() {
	a.m.HandleFunc("/", a.handleEcho("ping"))
	a.m.Handle("/favicon.ico", http.NotFoundHandler())

	a.m.Get("/laptop-loan", a.handleLaptopLoan("example/app/sql-01/views/laptop-loan.html"))
	a.m.Post("/laptop-loan", a.handleLaptopLoanPost("example/app/sql-01/tmp", "/api/v1/laptop-loan", "/laptop-loan"))

	a.m.Post("/api/v1/laptop-loan", a.handleAPILaptopLoan(`insert into "laptop_loan" values (?, ?, ?, ?, ?)`))
	a.m.Get("/api/v1/laptop-loan/", a.handleAPISearchLaptopLoan())
}

func (a *App) handleEcho(message string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, message, http.StatusOK)
	}
}
