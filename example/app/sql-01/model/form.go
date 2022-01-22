package model

import (
	"fmt"
	"time"
)

type LoanForm struct {
	UserID    string
	StartDate *time.Time
	EndDate   *time.Time
	TmpPath   string
}

// Debugging
func (lf LoanForm) String() string {
	return fmt.Sprintf("user %s; start %v; end %v;file: %s", lf.UserID, lf.StartDate, lf.EndDate, lf.TmpPath)
}
