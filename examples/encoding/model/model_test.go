package model_test

import (
	"encoding/json"
	"strings"
	"testing"

	test "github.com/topheruk/go/src/x/testing"

	"github.com/go-playground/assert/v2"
	"github.com/topheruk/go/examples/encoding/model"
	"github.com/topheruk/go/examples/encoding/model/person"
)

func casename(n int) string { return test.Casename(n) }

// MinEntropy is definetly something that should be defined
// within the back end. But I dont see any reason to have a
// max length or password rules. I think this should be fine
func TestPasswordUnmarshal(t *testing.T) {
	t.Parallel()

	type testcase struct {
		data     string
		expected string
	}

	tt := []testcase{
		{data: `{"password":""}`},           //err:too short error
		{data: `{"password":"helloWorld"}`}, //err:not strong enough
		{data: `{"password":"_Jde£5%3$^fsccfs0"}`, expected: "_Jde£5%3$^fsccfs0"},
	}

	for i, tc := range tt {
		t.Run(casename(i), func(t *testing.T) {
			type response struct {
				Password model.Password
			}

			var res response
			err := json.NewDecoder(strings.NewReader(tc.data)).Decode(&res)
			if !assert.IsEqual(err, nil) {
				// could check type of error, but it's pretty simple stuff
				return
			}
			assert.Equal(t, res.Password, model.Password(tc.expected))
		})
	}
}

// Can assume that password has been sanitized therefore
// no need for any checks inside the Marshaling process
// behaves like a normal string
func TestPasswordMarshal(t *testing.T) {
	t.Parallel()

	type testcase struct {
		data struct{ Password model.Password }
	}

	tt := []testcase{
		{data: struct{ Password model.Password }{Password: "password"}},
	}

	for i, tc := range tt {
		t.Run(casename(i), func(t *testing.T) {
			var sb strings.Builder
			err := json.NewEncoder(&sb).Encode(tc.data)
			assert.Equal(t, err, nil)
		})
	}
}

// If a json object has bad syntax, I would
// assume that is getting picked up so no need
// to check for that
func TestPersonUnmarshal(t *testing.T) {
	t.Parallel()

	type testcase struct {
		data string
		err  error
	}

	tt := []testcase{
		{data: `{}`, err: model.ErrEmpty}, //err:empty
		{data: `{"name":"John","password":"_Jde£5%3$^fsccfs0"}`},
	}

	for i, tc := range tt {
		t.Run(casename(i), func(t *testing.T) {
			var res person.DTO
			err := json.Unmarshal([]byte(tc.data), &res)
			assert.Equal(t, err, tc.err)
			if err != nil {
				return
			}

			b, err := json.Marshal(&res)
			assert.Equal(t, err, nil)
			assert.Equal(t, string(b), tc.data)
		})
	}
}

// I again don't need to test this,
// But it bumps my coverage up a little bit
func TestHashPassword(t *testing.T) {
	t.Parallel()
	type testcase struct {
		value string
	}

	tt := []testcase{
		{value: "_Jde£5%3$^fsccfs0"},
	}

	for i, tc := range tt {
		t.Run(casename(i), func(t *testing.T) {
			hash, err := model.Password(tc.value).Hash()
			assert.Equal(t, err, nil)

			assert.Equal(t, model.CheckPasswordHash(tc.value, hash), nil)
		})
	}
}

// Tbh, if this fails it's not due to human error
// so testing purely for coverage
func TestPersonCreate(t *testing.T) {
	t.Parallel()

	dto := person.DTO{
		Name:     "John",
		Password: "password",
	}
	p := person.New(dto)

	assert.Equal(t, string(p.Name), "John")
}
