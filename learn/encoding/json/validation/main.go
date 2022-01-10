package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/qri-io/jsonschema"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	var p Person

	if err := validation(
		"learn/encoding/json/validation/person.schema.json",
		"learn/encoding/json/validation/person.json",
		&p,
	); err != nil {
		log.Fatalln(err)
	}

	// decode("learn/encoding/json/validation/person.json", &p)

	log.Printf("success %v\n", p)
}

func decode(bodyUrl string, v interface{}) (err error) {
	bodyBytes, err := ioutil.ReadFile(bodyUrl)
	if err != nil {
		return
	}

	return json.Unmarshal(bodyBytes, v)
}

func validation(schemaUrl, bodyUrl string, v interface{}) (err error) {
	var schema = &jsonschema.Schema{}

	schemeBytes, err := ioutil.ReadFile(schemaUrl)
	if err != nil {
		return
	}

	bodyBytes, err := ioutil.ReadFile(bodyUrl)
	if err != nil {
		return
	}

	if err = json.Unmarshal(schemeBytes, schema); err != nil {
		return
	}

	if _, err = schema.ValidateBytes(context.TODO(), bodyBytes); err != nil {
		return
	}

	if err = json.Unmarshal(bodyBytes, v); err != nil {
		return
	}

	return
}
