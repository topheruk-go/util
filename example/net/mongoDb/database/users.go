package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
	Age  int                `bson:"age" json:"age"`
}

func (db *Database) FindManyUsers(ctx context.Context, filter interface{}) (users []User, err error) {
	cur, err := db.Collection("users").Find(ctx, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var el User
		cur.Decode(&el)
		users = append(users, el)
	}

	return users, nil
}

func (db *Database) FindUser(ctx context.Context, filter interface{}) (user *User, err error) {
	var el User
	err = db.Collection("users").FindOne(ctx, filter).Decode(&el)
	return &el, err
}

func (db *Database) InsertUser(ctx context.Context, doc interface{}) (id primitive.ObjectID, err error) {
	res, err := db.Collection("users").InsertOne(ctx, doc)
	if err != nil {
		return
	}

	return res.InsertedID.(primitive.ObjectID), err
}

func (db *Database) DeleteUser(ctx context.Context, filter interface{}) {
	// del, err := db.Collection("users").DeleteOne(ctx, filter)

	// del.DeletedCount
}
