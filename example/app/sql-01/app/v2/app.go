package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/topheruk/go/example/app/sql-01/db/v2"
	"github.com/topheruk/go/src/encoding"
	"github.com/topheruk/go/src/template"
)

type App struct {
	m  *chi.Mux
	db *db.DB
}

func New(db *db.DB) *App {
	a := &App{
		m:  chi.NewMux(),
		db: db,
	}
	a.routes()
	return a
}

func (a *App) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	a.m.ServeHTTP(rw, r)
}

// panic if error
func (a *App) render(path ...string) template.RenderFunc {
	render, err := template.Render(path...)
	if err != nil {
		panic(err)
	}
	return render
}

func (a *App) respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	encoding.Respond(rw, r, data, status)
}

func (a *App) decode(rw http.ResponseWriter, r *http.Request, data interface{}) error {
	return encoding.Decode(rw, r, data)
}

func (a *App) redirect(rw http.ResponseWriter, r *http.Request, url string, status int) {
	http.Redirect(rw, r, url, status)
}
