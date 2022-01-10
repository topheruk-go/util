package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"-"`
}

type DtoUser struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewUser(dto *DtoUser) *User {
	return &User{
		ID:       newUserID(),
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func newUserID() primitive.ObjectID {
	return primitive.NewObjectIDFromTimestamp(time.Now())
}
