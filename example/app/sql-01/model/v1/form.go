package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type LoanForm struct {
	ID        uuid.UUID
	StudentID string
	StartDate *time.Time
	EndDate   *time.Time
	TmpPath   string
}

// Debugging
func (lf LoanForm) String() string {
	return fmt.Sprintf("user %s; start %v; end %v;file: %s", lf.StudentID, lf.StartDate, lf.EndDate, lf.TmpPath)
}
