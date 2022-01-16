package main

import (
	"context"
	"fmt"
)

// no returning
func update(ctx context.Context, s *Service, dto interface{}) error {
	return s.Query(ctx, `UPDATE user SET email = :email WHERE email = :old_email`, dto)
}

// returning the new object
func updateR(ctx context.Context, s *Service, target interface{}, dto interface{}) error {
	return s.Query(ctx,
		`UPDATE user SET email = :email WHERE email = :old_email RETURNING *`,
		dto, target)
}

func insert(ctx context.Context, s *Service, dto *DtoUser) error {
	u, err := NewUser(dto)
	if err != nil {
		return err
	}

	return s.Query(ctx, `INSERT INTO user VALUES (:id, :email, :password, :created_at)`, u)
}

func getAll(ctx context.Context, s *Service) error {
	var uu []User
	err := s.Query(ctx, `SELECT * FROM user`, &uu)
	if err != nil {
		return err
	}
	for _, u := range uu {
		fmt.Printf("get all method: %v\n", u)
	}
	return nil
}

func get(ctx context.Context, s *Service, dto *DtoUser) error {
	var u User
	err := s.Query(ctx, `SELECT * FROM user WHERE email = :email`, dto, &u)
	if err != nil {
		return err
	}
	if err = u.ValidPassword(dto.Password); err != nil {
		return err
	}
	fmt.Printf("get method: %v\n", u)
	return nil
}
