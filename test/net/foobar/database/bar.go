package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bar struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Value string             `json:"value" bson:"value"`
}
