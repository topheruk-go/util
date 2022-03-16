package model

type Person struct {
	ID int `json:"id"`
	PersonDTO
}

type PersonDTO struct {
	Name *string `json:"name"`
}
