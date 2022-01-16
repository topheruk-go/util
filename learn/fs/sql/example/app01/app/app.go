package app

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/topheruk/go/learn/fs/sql/example/app01/service"
	"github.com/topheruk/go/src/encoding"
)

type app struct {
	m  *chi.Mux
	db *service.DB
}

func (a *app) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	a.m.ServeHTTP(rw, r)
}

func (a *app) respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	encoding.Respond(rw, r, data, status)
}

func (a *app) decode(rw http.ResponseWriter, r *http.Request, data interface{}) (err error) {
	return encoding.Decode(rw, r, data)
}

func (a *app) GetID(ctx context.Context) uuid.UUID {
	return uuid.MustParse(chi.URLParamFromCtx(ctx, "id"))
}

func New(db *service.DB) *app {
	a := &app{
		m:  chi.NewMux(),
		db: db,
	}
	a.routes()
	return a
}

func (a *app) routes() {
	a.m.HandleFunc("/ping", a.echo("ping"))

	a.m.Post("/user", a.handleCreateUser(`INSERT INTO user VALUES (:id, :email, :password, :created_at)`))
	a.m.Get("/user/{id}", a.handleSearchUser(`SELECT * FROM user WHERE id = :id`))
	a.m.Get("/user/", a.handleSearchAll(`SELECT * FROM user`))
}
