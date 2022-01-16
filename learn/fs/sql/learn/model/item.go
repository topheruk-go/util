package model

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}

type DtoItem struct {
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}

func NewItem(dto *DtoItem) (*Item, error) {
	i := &Item{
		uuid.New(),
		dto.Title,
		dto.CreatedAt,
	}
	return i, nil
}
