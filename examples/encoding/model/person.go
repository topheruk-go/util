package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Person struct {
	ID uuid.UUID
	PersonDto
}

func newPerson(dto PersonDto) *Person {
	p := Person{
		ID:        uuid.New(),
		PersonDto: dto,
	}
	return &p
}

type PersonDto struct {
	Name     Name     `json:"name"`
	Password Password `json:"password"`
}

func (dto *PersonDto) UnmarshalJSON(data []byte) error {
	if string(data) == `{}` {
		return ErrEmpty
	}

	type A PersonDto
	aux := &struct{ *A }{
		A: (*A)(dto),
	}

	// var aux Alias
	return json.Unmarshal(data, &aux)
}
