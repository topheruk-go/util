package encoding

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/qri-io/jsonschema"
)

func Respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(rw).Encode(data)
		if err != nil {
			http.Error(rw, "Could not encode in json", status)
		}
	}
}

func Decode(rw http.ResponseWriter, r *http.Request, data interface{}) (err error) {
	return json.NewDecoder(r.Body).Decode(data)
}

func loadSchema(url string) (*jsonschema.Schema, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	v := &jsonschema.Schema{}
	return v, json.Unmarshal(b, v)
}

func loadSchemaFromString(url string) (*jsonschema.Schema, error) {
	b, err := os.ReadFile(url)
	if err != nil {
		return nil, err
	}

	v := &jsonschema.Schema{}
	return v, json.Unmarshal(b, v)
}

func Validator(url string) (func(http.ResponseWriter, *http.Request, interface{}) error, error) {
	var (
		init sync.Once

		err error
		sch *jsonschema.Schema
	)

	init.Do(func() { sch, err = loadSchemaFromString(url) })

	var f = func(rw http.ResponseWriter, r *http.Request, v interface{}) (err error) {
		bodb, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}
		defer r.Body.Close()

		errs, err := sch.ValidateBytes(r.Context(), bodb)
		if err != nil {
			return err
		}

		// FIXME: format all error msgs, not just the first one
		if len(errs) > 0 {
			return fmt.Errorf("%v", errs[0].Message)
		}

		return json.Unmarshal(bodb, v)
	}

	return f, err
}
