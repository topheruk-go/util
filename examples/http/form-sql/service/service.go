package service

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	h "github.com/topheruk/go/src/x/http/handler"
	q "github.com/topheruk/go/src/x/http/request"
	p "github.com/topheruk/go/src/x/http/response"
	"github.com/topheruk/go/src/x/parse"
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
	p.New(rw, r, data, status)
}

func (s *Service) Err(rw http.ResponseWriter, r *http.Request, err error, status int) {
	p.Err(rw, r, err, status)
}

func (s *Service) Todo(rw http.ResponseWriter, r *http.Request) {
	p.Todo(rw, r)
}

func (s *Service) getId(rw http.ResponseWriter, r *http.Request) (int, error) {
	return strconv.Atoi(chi.URLParamFromCtx(r.Context(), "id"))
}

func (s *Service) Decode(rw http.ResponseWriter, r *http.Request, data interface{}) error {
	return q.New(rw, r, data)
}

func (s *Service) Echo(message string) http.HandlerFunc { return h.Echo(message) }

func (s *Service) AbsoluteURL(rw http.ResponseWriter, r *http.Request) string {
	return parse.AbsoluteURL(rw, r)
}

func (s *Service) migrate() {
	// pull from migration folder if this was a bigger project
	s.db.MustExec(`CREATE TABLE person (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT UNIQUE NOT NULL)`)
}
