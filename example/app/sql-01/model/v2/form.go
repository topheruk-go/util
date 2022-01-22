package model

import (
	"time"

	"github.com/google/uuid"
)

// Auto-Generate PDF using text/template package
//
// STUDENT FILL-OUT
// -- Name
// -- Student ID
// -- Group Lead & Programme Lead Email
// -- Student Signature
// -- -- Date
// -- Guardian Signature (if < 18)
// -- -- Date
// ACADEMIC STAFF FILL-OUT
// -- Group Lead Name & Signature
// -- -- Date
// -- Programme Manager Name & Signature
// -- Loan Duration (probs makes more sense for us to have this control)
// -- -- Date
// IT DEPT FILL-OUT
// -- Laptop Meta
// -- IT Signature

type LoanFormMeta struct {
	StudentMeta StudentMeta
}

type StudentMeta struct {
	ID                 uuid.UUID
	StudentID          string
	Name               string
	Email              string
	GroupLeadEmail     string
	ProgrammeLeadEmail string
	DateRequested      *time.Time
}

type AcademicStaffMeta struct {
	*StudentMeta
	ID         uuid.UUID
	Signature  []byte
	DateSigned *time.Time
}

type ITStaffMeta struct {
	*AcademicStaffMeta
	ID           uuid.UUID
	LaptopMeta   LaptopMeta
	DateApproved *time.Time
	Signature    []byte
}

type LaptopMeta struct {
	ID      uuid.UUID
	AssetID string
	Model   string
}
