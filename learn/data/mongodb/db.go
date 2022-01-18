package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	InsertOne(ctx context.Context, f bson.D, d Datum) error
	FindOne(ctx context.Context, f bson.D) error
}

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

func (s *Service) ObjectIDFromHex(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

func (s *Service) FilterByID(oid primitive.ObjectID) bson.D { return bson.D{{Key: "_id", Value: oid}} }

func (s *Service) Nil() bson.D { return bson.D{} }

func (s *Service) CompanyUsers() *mongo.Collection { return s.Database("company").Collection("users") }
