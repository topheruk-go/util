package encoding

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func ValidateWithTags() (f func(rw http.ResponseWriter, r *http.Request, v interface{}) error, err error) {
	var (
		valid *validator.Validate
		init  sync.Once
	)

	init.Do(func() { valid = validator.New() })

	return func(rw http.ResponseWriter, r *http.Request, v interface{}) (err error) {
		if err = json.NewDecoder(r.Body).Decode(v); err != nil {
			return
		}

		return valid.Struct(v)
	}, err
}

func ValidateWithSchema(url string) (f func(rw http.ResponseWriter, r *http.Request, data *interface{}) error, err error) {
	var (
		init sync.Once

		schema *jsonschema.Schema
	)

	init.Do(func() { schema, err = jsonschema.Compile(url) })
	if err != nil {
		return nil, err
	}

	f = func(rw http.ResponseWriter, r *http.Request, v *interface{}) (err error) {
		if err = json.NewDecoder(r.Body).Decode(v); err != nil {
			return
		}

		return schema.Validate(*v)
	}

	return f, err
}
