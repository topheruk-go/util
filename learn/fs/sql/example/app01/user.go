package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type DtoUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Email        string    `db:"email" json:"email"`
	PasswordHash []byte    `db:"password" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"-"`
}

func NewUser(dto *DtoUser) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 0)
	if err != nil {
		return nil, fmt.Errorf("encrypting password failed: %w", err)
	}

	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("generating uuid failed: %w", err)
	}

	return &User{
		ID:           uid,
		Email:        dto.Email,
		PasswordHash: hash,
		CreatedAt:    time.Now(),
	}, nil
}

func (u *User) Valid(password string) error {
	err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password))
	if err != nil {
		return fmt.Errorf("password mismatch: %w", err)
	}
	return nil
}
