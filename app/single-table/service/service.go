package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/topheruk/go/src/encoding"
	"github.com/topheruk/go/src/http/handler"
)

type Service struct {
	m  *chi.Mux
	db *sqlx.DB
}

func New(db *sqlx.DB) *Service {
	s := &Service{
		db: db,
		m:  chi.NewMux(),
	}
	s.migrate()
	s.routes()
	return s
}

func (s *Service) ServeHTTP(rw http.ResponseWriter, r *http.Request) { s.m.ServeHTTP(rw, r) }

func (s *Service) Close() error { return s.db.Close() }

func (s *Service) Respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	encoding.Respond(rw, r, data, status)
}

func (s *Service) Err(rw http.ResponseWriter, r *http.Request, err error, status int) {
	type e struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	s.Respond(rw, r, &e{status, err.Error()}, status)
}

func (s *Service) Todo(rw http.ResponseWriter, r *http.Request) {
	s.Err(rw, r, errors.New("todo"), http.StatusInternalServerError)
}

func (s *Service) getId(rw http.ResponseWriter, r *http.Request) (int, error) {
	return strconv.Atoi(chi.URLParamFromCtx(r.Context(), "id"))
}

func (s *Service) Decode(rw http.ResponseWriter, r *http.Request, data interface{}) error {
	return encoding.Decode(rw, r, data)
}

func (s *Service) Echo(message string) http.HandlerFunc { return handler.Echo(message) }

func (s *Service) AbsoluteURL(rw http.ResponseWriter, r *http.Request) string {
	return fmt.Sprintf("%s://%s%s", strings.ToLower(strings.SplitN(r.Proto, "/", 2)[0]), r.Host, r.URL)
}

func (s *Service) migrate() {
	// pull from migration folder if this was a bigger project
	s.db.MustExec(`CREATE TABLE person (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT UNIQUE NOT NULL)`)
}
