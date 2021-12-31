package src

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
)

func Validate() (f func(rw http.ResponseWriter, r *http.Request, v interface{}) error, err error) {
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
