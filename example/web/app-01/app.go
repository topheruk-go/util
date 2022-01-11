package app01

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/topheruk/go/example/web/app-01/model"
	"github.com/topheruk/go/src/encoding"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type App struct {
	m  *chi.Mux
	db UserService
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
		// us: TODO:,
	}
	a.routes()
	return a
}

func (a *App) routes() {
	a.m.Get("/ping", a.handlePing())
	a.m.Get("/echo", a.handleEcho("this is an example of an echo handler"))

	a.m.Post("/user", a.handleCreateUser("example/web/json/user.schema.json"))
	a.m.Put("/user/{id}", a.handleUpdateUser("example/web/json/user.schema.json"))
	a.m.Get("/user/{id}", a.handleSearchUser())
	a.m.Delete("/user/{id}", a.handleDeleteUser())
}

type UserService interface {
	Search(context.Context, primitive.ObjectID) (*model.User, error)
	Insert(context.Context, *model.DtoUser) error
	Delete(context.Context, primitive.ObjectID) error
	Update(context.Context, primitive.ObjectID, *model.DtoUser) (*model.User, error)
}
