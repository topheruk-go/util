package parse

import (
	"encoding/json"
	"io/ioutil"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type bsonCommand struct {
	CollMod   string `json:"collMod" bson:"collMod"`
	Validator struct {
		Schema interface{} `json:"$jsonSchema" bson:"$jsonSchema"`
	} `json:"validator" bson:"validator"`
}

func newBsonCommand(path string) (bcmd *bsonCommand, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	var v bsonCommand
	if err = json.Unmarshal(data, &v); err != nil {
		return
	}

	return &v, err
}

func BsonCmd(path string) (v *primitive.D, err error) {
	bcmd, err := newBsonCommand(path)
	if err != nil {
		return
	}

	return ToDoc(bcmd)
}

func ToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}
