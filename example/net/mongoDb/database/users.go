package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
	err = db.Collection("users").FindOne(ctx, filter).Decode(&user)
	return
}

func (db *Database) InsertUser(ctx context.Context, doc interface{}) (id primitive.ObjectID, err error) {
	res, err := db.Collection("users").InsertOne(ctx, doc)
	if err != nil {
		return
	}

	return res.InsertedID.(primitive.ObjectID), err
}

func (db *Database) UpdateUser(ctx context.Context, id primitive.ObjectID, doc interface{}, upsert bool) (up *User, err error) {
	b, err := bson.Marshal(doc)
	if err != nil {
		return
	}

	var upt bson.D
	err = bson.Unmarshal(b, &upt)
	if err != nil {
		return
	}

	log.Println(doc)
	log.Println(upt)

	res := db.Collection("users").FindOneAndUpdate(ctx, bson.D{{"_id", id}}, bson.D{{"$set", upt}})

	if res.Err() != nil {
		return
	}

	if err = res.Decode(&up); err != nil {
		return
	}

	return
}

func (db *Database) DeleteUser(ctx context.Context, filter interface{}) (delCount int, err error) {
	del, err := db.Collection("users").DeleteOne(ctx, filter)
	return int(del.DeletedCount), err
}
