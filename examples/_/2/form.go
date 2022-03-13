package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type LoanFormDto struct {
	StudentID string
	StartDate *time.Time
	EndDate   *time.Time
	TmpPath   string
}

type LoanForm struct {
	ID uuid.UUID
	LoanFormDto
}

func newLoanForm(dto LoanFormDto) *LoanForm {
	return &LoanForm{
		ID:          uuid.New(),
		LoanFormDto: dto,
	}
}

// Debugging
func (lf LoanFormDto) String() string {
	return fmt.Sprintf("user %s; start %v; end %v;file: %s", lf.StudentID, lf.StartDate, lf.EndDate, lf.TmpPath)
}
