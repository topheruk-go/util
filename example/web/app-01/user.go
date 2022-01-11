package app01

import (
	"fmt"
	"time"

	"github.com/topheruk/go/src/parse"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"-"`
}

func (u User) ToBSON() (primitive.D, error) { return parse.ToBSON(u) }
func (u User) String() string               { return fmt.Sprintf("%v: %s", u.ID, u.Email) }

type DtoUser struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (dto DtoUser) ToBSON() (primitive.D, error) { return parse.ToBSON(dto) }
func (dto DtoUser) String() string               { return dto.Email }

func (dto DtoUser) User() *User {
	return &User{
		ID:       newUserID(),
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func newUserID() primitive.ObjectID {
	return primitive.NewObjectIDFromTimestamp(time.Now())
}
