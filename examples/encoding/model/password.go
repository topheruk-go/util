package model

import (
	gpv "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
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

type PasswordHash string

func (p Password) Hash() (PasswordHash, error) {
	return HashPassword(p)
}

func (p Password) MustHash() PasswordHash {
	hash, err := HashPassword(p)
	if err != nil {
		panic(err)
	}
	return hash
}

func HashPassword(password Password) (PasswordHash, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return PasswordHash(bytes), err
}

func CheckPasswordHash(password string, hash PasswordHash) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
