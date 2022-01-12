package main

import (
	"github.com/topheruk/go/src/parse"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Datum interface {
	ToBSON() (primitive.D, error)
	String() string
}

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Email        string             `bson:"email" json:"email"`
	PasswordHash string             `bson:"hash" json:"-"`
}

func (u User) ToBSON() (primitive.D, error) { return parse.ToBSON(u) }
func (u User) String() string               { return u.ID.String() + " " + u.Email }

func (dto *DTO) New() (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 0)

	return &User{
		ID:           primitive.NewObjectID(),
		Email:        dto.Email,
		PasswordHash: string(hash),
	}, err
}

type DTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto DTO) ToBSON() (primitive.D, error) { return parse.ToBSON(dto) }
func (dto DTO) String() string               { return dto.Email }
