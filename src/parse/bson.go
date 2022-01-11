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

	*v, err = ToBSON(bcmd)
	return
}

func ToBSON(v interface{}) (primitive.D, error) {
	b, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}

	var d = primitive.D{}
	return d, bson.Unmarshal(b, &d)
}

func FromBSON(v interface{}) {}
