package app01

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	db *mongo.Database
}

func (s *Service) FindOne(ctx context.Context, f bson.D, d Datum) error {
	return s.db.Collection("users").FindOne(ctx, f).Decode(d)
}

func (s *Service) InsertOne(ctx context.Context, d Datum) error {
	_, err := s.db.Collection("users").InsertOne(ctx, d)
	return err
}

func (s *Service) UpsertOne(ctx context.Context, f bson.D, d Datum) error {
	return nil
}
