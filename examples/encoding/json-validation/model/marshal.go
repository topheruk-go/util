package model

import (
	"errors"

	validator "github.com/wagslane/go-password-validator"
)

var ErrMarshal = errors.New("error: cannot marshal")
var ErrUnmarshal = errors.New("error: cannot unmarshal")

// this number can be what ever I want.
// for testing I am having it at 10 but will amend to a more realistic number later
const minEntropy = 10

func (p *Password) UnmarshalText(data []byte) error {
	if err := validator.Validate(string(data), minEntropy); err != nil {
		return err
	}

	*p = Password(data)
	return nil
}

func (p Password) MarshalText() ([]byte, error) {
	return []byte(p), nil
}
