package service

import (
	"errors"
	"net/http"

	"github.com/topheruk/go/app/model"
)

func (s *Service) routes() {
	s.m.Get("/ping", s.Echo("ping"))

	s.m.Post("/person", s.handleInsertPerson("INSERT INTO person (name) VALUES (:name) RETURNING *"))
	s.m.Get("/person", s.handleSelectPersonSlice("SELECT * FROM person"))
	s.m.Get("/person/{id}", s.handleSelectPerson("SELECT * FROM person WHERE id=?"))
	s.m.Delete("/person/{id}", s.handleDeletePerson("DELETE FROM person WHERE id=?"))
}

func (s *Service) handleInsertPerson(query string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var dto model.PersonDTO
		if err := s.Decode(rw, r, &dto); err != nil {
			s.Err(rw, r, err, http.StatusBadRequest)
			return
		}
		var p model.Person
		stmt, err := s.db.PrepareNamedContext(r.Context(), query)
		if err != nil {
			s.Err(rw, r, err, http.StatusInternalServerError)
			return
		}
		defer stmt.Close()
		if err := stmt.Get(&p, dto); err != nil {
			s.Err(rw, r, err, http.StatusInternalServerError)
			return
		}
		s.Respond(rw, r, &p, http.StatusOK)
	}
}

func (s *Service) handleSelectPersonSlice(query string) http.HandlerFunc {
	// TODO: impl search params
	return func(rw http.ResponseWriter, r *http.Request) {
		var ps []model.Person
		if err := s.db.Select(&ps, query); err != nil {
			s.Err(rw, r, err, http.StatusInternalServerError)
			return
		}
		s.Respond(rw, r, &ps, http.StatusOK)
	}
}

func (s *Service) handleSelectPerson(query string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		uid, err := s.ID(rw, r)
		if err != nil {
			s.Err(rw, r, err, http.StatusBadRequest)
			return
		}
		var p model.Person
		if err := s.db.Get(&p, query, uid); err != nil {
			s.Err(rw, r, err, http.StatusInternalServerError)
			return
		}
		s.Respond(rw, r, &p, http.StatusOK)
	}
}

func (s *Service) handleDeletePerson(query string) http.HandlerFunc {
	// type response struct{}
	return func(rw http.ResponseWriter, r *http.Request) {
		uid, err := s.ID(rw, r)
		if err != nil {
			s.Err(rw, r, err, http.StatusBadRequest)
			return
		}
		// TODO: should passing a value out-of-bounds or non-int types
		// create an error?
		res, err := s.db.Exec(query, uid)
		if err != nil {
			s.Err(rw, r, err, http.StatusInternalServerError)
			return
		}
		// TODO: feels like a hack, investigate
		if i, _ := res.RowsAffected(); i == 0 {
			s.Err(rw, r, errors.New("error: could not find a match"), http.StatusInternalServerError)
			return
		}
		s.Respond(rw, r, uid, http.StatusOK)
	}
}
