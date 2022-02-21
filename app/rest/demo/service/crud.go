package service

import (
	"fmt"
)

func (s *Service) create(dest interface{}, query string, args ...interface{}) error {
	return s.db.Get(dest, query, args...)
}

func (s *Service) read(dest interface{}, query string, args ...interface{}) error {
	return s.db.Get(dest, query, args...)
}
func (s *Service) readSlice(dest interface{}, query string, args ...interface{}) error {
	return s.db.Select(dest, query, args...)
}

func (s *Service) update() error { return ErrTodo }

func (s *Service) delete(query string, uid int) error {
	res, err := s.db.Exec(query, uid)
	if err != nil {
		return err
	}
	i, err := res.RowsAffected()
	if i == 0 {
		fmt.Println(err)
		return ErrNoMatch
	}
	return nil
}
