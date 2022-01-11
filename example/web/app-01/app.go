package app01

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/topheruk/go/src/encoding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type App struct {
	m  *chi.Mux
	db Database
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

func (a *App) param(r *http.Request, param string) string {
	return chi.URLParamFromCtx(r.Context(), param)
}

func (a *App) OID(r *http.Request) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(a.param(r, "id"))
}

func (a *App) filterByID(id primitive.ObjectID) bson.D { return bson.D{{Key: "_id", Value: id}} }

type Datum interface {
	ToBSON() (primitive.D, error)
	String() string
}

type Database interface {
	InsertOne(ctx context.Context, d Datum) error
	FindOne(ctx context.Context, f bson.D, d Datum) error
}
