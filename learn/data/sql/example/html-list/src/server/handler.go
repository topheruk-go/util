package main

import (
	"fmt"
	"log"
	"net/http"

	t "github.com/topheruk/go/src/template"
)

func (a *app) handleEcho(message string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.respond(rw, r, message, http.StatusOK)
	}
}

type Page struct {
	Title string
}

// #?sort=name&dir=asc
func (a *app) handleIndex(path string) http.HandlerFunc {
	type page struct {
		Page
		Accounts []account
	}

	renderFn, err := t.Render(path)
	return func(rw http.ResponseWriter, r *http.Request) {
		if err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		sort := r.URL.Query()["sort"]
		log.Println(sort)

		var aa []account
		if err := a.filterQuery(rw, r, &aa); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		p := &page{
			Page:     Page{Title: "This is an example of using SQL to populate an unordered list"},
			Accounts: aa,
		}

		renderFn(rw, r, p)
	}
}

func (a *app) filterQuery(rw http.ResponseWriter, r *http.Request, accs *[]account) error {
	stmt := `select * from account`

	sort, ok := r.URL.Query()["sort"]
	if ok && len(sort) == 2 {
		fmt.Println(sort)
	}

	return a.db.Queryi(stmt, &accs)
}

func (a *app) handleGetAll() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var aa []account
		if err := a.db.Queryi(`select * from account`, &aa); err != nil {
			a.respond(rw, r, err, http.StatusInternalServerError)
			return
		}

		a.respond(rw, r, aa, http.StatusOK)
	}
}
