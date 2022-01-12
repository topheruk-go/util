package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// 12 Methods per document is required
// FindUser(ctx context.Context, f bson.D, d Datum) error
// FindUserMany(ctx context.Context, f bson.D) ([]User, error)
// FindUserAll(ctx context.Context) ([]User, error)
// InsertUser
// InsertUserMany
// UpsertUser
// UpsertUserMany
// DeleteUser
// DeleteUserMany
// DeleteUserAll
// UpdateUser
// UpdateUserMany

func (s *Service) FindUser(ctx context.Context, f bson.D, d Datum) error {
	return s.CompanyUsers().FindOne(ctx, f).Decode(d)
}

func (s *Service) FindUserAll(ctx context.Context) ([]User, error) {
	return s.FindUserMany(ctx, s.Nil())
}

func (s *Service) FindUserMany(ctx context.Context, f bson.D) ([]User, error) {
	cur, err := s.CompanyUsers().Find(ctx, f)
	if err != nil {
		return nil, err
	}

	var us []User
	for cur.Next(ctx) {
		var u User
		err := cur.Decode(&u)
		if err != nil {
			panic(err)
		}

		us = append(us, u)
	}

	return us, nil
}

func (s *Service) InsertUser(ctx context.Context, d Datum) (err error) {
	_, err = s.CompanyUsers().InsertOne(ctx, d)
	return
}

func (s *Service) UpsertUser(ctx context.Context, f bson.D, d Datum) error {
	// // Only required for Upserting
	// doc, err := d.ToBSON()
	// if err != nil {
	// 	return
	// }

	return nil
}

func (s *Service) DeleteUser(ctx context.Context, f bson.D) error {
	dr, err := s.CompanyUsers().DeleteOne(ctx, f)
	if dr.DeletedCount < 1 {
		return fmt.Errorf("mongo: no documents in result")
	}
	return err
}

func (s *Service) DeleteUserAll(ctx context.Context) error {
	return s.DeleteUserMany(ctx, s.Nil())
}

func (s *Service) DeleteUserMany(ctx context.Context, f bson.D) error {
	dr, err := s.CompanyUsers().DeleteMany(ctx, f)
	if dr.DeletedCount < 1 {
		return fmt.Errorf("mongo: no documents in result")
	}
	return err
}
