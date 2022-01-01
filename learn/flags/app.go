package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type bsoncommand struct {
	CollMod   string `json:"collMod" bson:"collMod"`
	Validator struct {
		Schema interface{} `json:"$jsonSchema" bson:"$jsonSchema"`
	} `json:"validator" bson:"validator"`
}

func NewBsonCommand(path string) (bcmd *bsoncommand, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	var v bsoncommand
	if err = json.Unmarshal(data, &v); err != nil {
		return
	}

	return &v, err
}

func parseSchema(path string) (v *primitive.D, err error) {
	bcmd, err := NewBsonCommand(path)
	if err != nil {
		return
	}

	return toDoc(bcmd)
}

func CmdFlags() (username, password *string, address *int) {
	username = flag.String("u", "", "database client username:password")
	password = flag.String("p", "", "database client password")
	address = flag.Int("a", 8000, "server port number")

	flag.Parse()
	return
}

func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}
