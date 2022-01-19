package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

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

func validation(schemaUrl, bodyUrl string, v interface{}) (err error) {
	sch, err := loadSchemaFromString(schemaUrl)
	if err != nil {
		return
	}

	var bodyBytes = []byte(`{
		"firstName" : "George",
		"lastName" : "Michael"
		}`)

	errs, err := sch.ValidateBytes(context.TODO(), bodyBytes)
	if err != nil {
		return err
	}

	if len(errs) > 0 {
		return fmt.Errorf("%v", errs)
	}

	return json.Unmarshal(bodyBytes, v)
}

func loadSchemaFromString(url string) (*jsonschema.Schema, error) {
	b, err := os.ReadFile(url)
	if err != nil {
		return nil, err
	}

	v := &jsonschema.Schema{}
	return v, json.Unmarshal(b, v)
}
