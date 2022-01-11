package main

import (
	"context"
	"fmt"

	"github.com/topheruk/go/src/parse"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Datum interface {
	ToBSON() (*primitive.D, error)
	String() string
}

type Database interface {
	InsertOne(ctx context.Context, f bson.D, d Datum) error
	FindOne(ctx context.Context, f bson.D) error
}

type Foo struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func (f Foo) ToBSON() (primitive.D, error) {
	return parse.ToBSON(f)
}

type Service struct {
	c *mongo.Collection
}

func (s *Service) FindOne(ctx context.Context, f bson.D, d Datum) error {
	return s.c.FindOne(ctx, f, nil).Decode(d)
}

func main() {
	var f = Foo{Name: "foo", Age: 20}
	doc, _ := f.ToBSON()
	fmt.Println(bson.D{{Key: "$set", Value: doc}})
}
