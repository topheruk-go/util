package person

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/topheruk/go/examples/encoding/model"
)

type Person struct {
	ID       uuid.UUID
	Password model.PasswordHash
	DTO
}

func New(dto DTO) *Person {
	return Must(NewPerson(dto))
}

func Must(p *Person, err error) *Person {
	if err != nil {
		panic(err)
	}
	return p
}

func NewPerson(dto DTO) (*Person, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	hash, err := dto.Password.Hash()
	if err != nil {
		return nil, err
	}

	p := Person{
		ID:       uid,
		Password: hash,
		DTO:      dto,
	}
	return &p, nil
}

type DTO struct {
	Name     model.Name     `json:"name"`
	Password model.Password `json:"password"`
}

func (dto *DTO) UnmarshalJSON(data []byte) error {
	if string(data) == `{}` {
		return model.ErrEmpty
	}

	type A DTO
	aux := &struct{ *A }{
		A: (*A)(dto),
	}

	return json.Unmarshal(data, &aux)
}
