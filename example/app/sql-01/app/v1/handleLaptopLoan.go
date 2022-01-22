package app

import (
	"net/http"
	"time"

	"github.com/topheruk/go/example/app/sql-01/model/v1"
)

func (a *App) handleLaptopLoan(path ...string) http.HandlerFunc {
	type response struct {
		PostURL          string
		MinDate, MaxDate string
	}

	render := a.render(path...)

	return func(rw http.ResponseWriter, r *http.Request) {
		sd, ed := timeDuration(time.Now())

		resp := &response{
			PostURL: r.URL.Path,
			MinDate: sd,
			MaxDate: ed,
		}

		render(rw, r, resp)
	}
}

func (a *App) handleLaptopLoanPost(tmpPath, postPath, redirectPath string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		lf := &model.LoanForm{TmpPath: tmpPath}

		if err := parseMultiPartForm(rw, r, lf); err != nil {
			a.respond(rw, r, err.Error(), http.StatusInternalServerError)
			return
		}

		url := parseURL(r, postPath)
		if err := postLoanForm(url, "application/json", lf); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		a.redirect(rw, r, redirectPath, http.StatusFound)
	}
}

func (a *App) handleAPILaptopLoan() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var lf model.LoanForm
		if err := a.decode(rw, r, &lf); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}
		// FIXME: add to sql database
		a.db.Forms = append(a.db.Forms, lf)
		a.respond(rw, r, lf, http.StatusOK)
	}
}

func (a *App) handleAPISearchLaptopLoan() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// TODO: search from database
		a.respond(rw, r, a.db.Forms, http.StatusOK)
	}
}
