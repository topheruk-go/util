package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

func (a *app) routes() {
	a.m.Use(middleware.Logger)
	a.m.Handle("/favicon.ico", http.NotFoundHandler())
	a.m.Get("/*", a.handleIndex("example/index.html"))
	a.m.Get("/laptop-loan", a.handleLaptopLoan("example/laptop-loan.html"))
}
