package service

import (
	"fmt"
	"net/http"

	"github.com/topheruk/go/app/rest/demo/model"
)

func (s *Service) routes() {
	s.m.Get("/ping", s.Echo("ping"))

	s.m.Post("/person", s.handleInsertPerson("INSERT INTO person (name) VALUES (?) RETURNING *"))
	s.m.Get("/person", s.handleSelectPersonSlice("SELECT * FROM person"))
	s.m.Get("/person/{id}", s.handleSelectPerson("SELECT * FROM person WHERE id=?"))
	s.m.Delete("/person/{id}", s.handleDeletePerson("DELETE FROM person WHERE id=?"))
}

// 201 if created
// Add newly created uri to location header
func (s *Service) handleInsertPerson(query string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var dto model.PersonDTO
		if err := s.Decode(rw, r, &dto); err != nil {
			s.Err(rw, r, err, http.StatusBadRequest)
			return
		}

		var p model.Person
		if err := s.create(&p, query, dto.Name); err != nil {
			s.Err(rw, r, err, http.StatusInternalServerError)
			return
		}
		// TODO: respond 201
		rw.Header().Add("Location", s.AbsoluteURL(rw, r)+"/"+fmt.Sprint(p.ID))
		s.Respond(rw, r, p, http.StatusCreated)
	}
}

func (s *Service) handleSelectPersonSlice(query string) http.HandlerFunc {
	// TODO: impl search params
	return func(rw http.ResponseWriter, r *http.Request) {
		var ps []model.Person
		if err := s.readSlice(&ps, query); err != nil {
			s.Err(rw, r, err, http.StatusInternalServerError)
			return
		}

		s.Respond(rw, r, ps, http.StatusOK)
	}
}

func (s *Service) handleSelectPerson(query string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		uid, err := s.getId(rw, r)
		if err != nil {
			s.Err(rw, r, err, http.StatusBadRequest)
			return
		}

		var p model.Person
		if err := s.read(&p, query, uid); err != nil {
			s.Err(rw, r, err, http.StatusInternalServerError)
			return
		}
		// if empty then its an error?
		s.Respond(rw, r, p, http.StatusOK)
	}
}

// 204 on successful req (Put, Patch, Delete)
func (s *Service) handleDeletePerson(query string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		uid, err := s.getId(rw, r)
		if err != nil {
			s.Err(rw, r, err, http.StatusBadRequest)
			return
		}

		if err := s.delete(query, uid); err != nil {
			s.Err(rw, r, err, http.StatusInternalServerError)
			return
		}

		s.Respond(rw, r, nil, http.StatusNoContent)
	}
}
