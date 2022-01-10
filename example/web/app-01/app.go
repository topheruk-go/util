package app01

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/topheruk/go/src/encoding"
)

type App struct {
	m *chi.Mux
}

func (a *App) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	a.m.ServeHTTP(rw, r)
}

func (a *App) respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	encoding.Respond(rw, r, data, status)
}

func (a *App) decode(rw http.ResponseWriter, r *http.Request, data interface{}) (err error) {
	return encoding.Decode(rw, r, data)
}

func New() *App {
	a := &App{
		m: chi.NewMux(),
	}
	a.routes()
	return a
}

func (a *App) routes() {
	a.m.Get("/ping", a.handlePing())
	a.m.Get("/echo", a.handleEcho("this is an example of an echo handler"))

	a.m.Post("/user", a.handleCreateUser("example/web/app-01/model/schema/user.schema.json"))
}

func (a *App) validate(url string) (f func(rw http.ResponseWriter, r *http.Request, data *interface{}) error, err error) {
	return encoding.ValidateWithSchema(url)
}
