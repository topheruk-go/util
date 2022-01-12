package app01

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	*mongo.Client
}

func NewService(ctx context.Context, uri string) (*Service, error) {
	cli, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err = cli.Connect(ctx); err != nil {
		return nil, err
	}

	return &Service{cli}, nil
}

func (s *Service) Close(ctx context.Context) error {
	return s.Disconnect(ctx)
}

func (s *Service) FindOne(ctx context.Context, f bson.D, d Datum) error {
	return s.Database("test").Collection("users").FindOne(ctx, f).Decode(d)
}

func (s *Service) InsertOne(ctx context.Context, d Datum) error {
	r, err := s.Database("test").Collection("users").InsertOne(ctx, d)

	_ = r.InsertedID
	return err
}

func (s *Service) UpsertOne(ctx context.Context, f bson.D, d Datum) error {
	return nil
}
