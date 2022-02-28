package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func casename(n int) string { return fmt.Sprintf("case_%d", n+1) }

func TestUnmarshalPassword(t *testing.T) {
	t.Parallel()
	type testcase struct {
		value string
		// err   error
	}

	tt := []testcase{
		{value: `{"password":"abcd1234"}`},
	}

	for i, tc := range tt {
		t.Run(casename(i), func(t *testing.T) {
			var v struct {
				Password Password `json:"password"`
			}
			assert.Equal(t, json.NewDecoder(strings.NewReader(tc.value)).Decode(&v), nil)

			var sb strings.Builder
			assert.Equal(t, json.NewEncoder(&sb).Encode(v), nil)
			assert.Equal(t, sb.String()[:len(sb.String())-1], tc.value)
		})
	}
}

func TestHashPassword(t *testing.T) {
	t.Parallel()
	type testcase struct {
		value string
	}

	tt := []testcase{
		{value: "1234abcd"},
	}

	for i, tc := range tt {
		t.Run(casename(i), func(t *testing.T) {
			hash, err := Password(tc.value).Hash()
			assert.Equal(t, err, nil)

			assert.Equal(t, CheckPasswordHash(tc.value, hash), nil)
		})
	}
}
