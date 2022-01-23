package main

import "net/http"

func (a *app) handleIndex(path string) http.HandlerFunc {
	render := a.render(path)
	return func(rw http.ResponseWriter, r *http.Request) {
		render(rw, r, nil)
	}
}

func (a *app) handleLaptopLoan(path string) http.HandlerFunc {
	render := a.render(path)
	return func(rw http.ResponseWriter, r *http.Request) {
		render(rw, r, nil)
	}
}
