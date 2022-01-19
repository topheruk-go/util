package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Foo struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Value int                `json:"value" bson:"value"`
}
