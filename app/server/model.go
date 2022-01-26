package main

import (
	"github.com/google/uuid"
)

type Loan struct {
	ID   uuid.UUID `db:"id"`
	Path string    `db:"path"`
	Age  int       `db:"age"`
	Name string    `db:"name"`
}

type LoanDto struct {
	Age  int    `db:"age"`
	Name string `db:"name"`
	File []byte
}

func newLoan(dto LoanDto, fileName string) *Loan {
	l := &Loan{
		ID:   uuid.New(),
		Age:  dto.Age,
		Name: dto.Name,
		Path: fileName,
	}
	return l
}
