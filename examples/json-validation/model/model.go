package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var ErrTodo = errors.New("todo")

type User struct {
	ID uuid.UUID
	UserDTO
	Password  []byte
	CreatedAt *time.Time
}

type UserDTO struct {
	Name     string
	Email    string
	Password Password
}

type Password string

func (p Password) Hash() (string, error) {
	return HashPassword(string(p))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
