package model

import (
	gpv "github.com/wagslane/go-password-validator"
)

var minEntropy = 60.0

type Password string

func (p *Password) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return ErrTooShort
	}
	*p = Password(data)
	return gpv.Validate(string(data), minEntropy)
}
