package main

import (
	"time"

	"github.com/google/uuid"
)

type laptoploan struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	StudentID string     `json:"student_id" db:"student_id"`
	StartDate *time.Time `json:"start_date" db:"start_date"`
	EndDate   *time.Time `json:"end_date" db:"end_date"`
	File      []byte     `json:"file" db:"file"`
}
