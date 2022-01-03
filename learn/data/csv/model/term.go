package model

import (
	"io"
	"time"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type Term struct {
	// [ required ] A unique identifier used to reference terms in the enrollments data. This identifier must not change for the account, and must be globally unique. In the user interface, this is called the SIS ID.
	ID string `csv:"term_id"`
	// [ required, sticky ] The nname of the term
	Name string `csv:"name"`
	// FIXME: [ required ] active, deleted
	Status string `csv:"status"`
	// Sets the integration_id of the term
	IntegrationID string `csv:"integration_id"`
	// When set, all columns except term_id, status, start_date, and end_date will be ignored for this row. Can only be used for an existing term. If status is active, the term dates will be set to apply only to enrollments of the given type. If status is deleted, the currently set dates for the given enrollment type will be removed. Must be one of StudentEnrollment, TeacherEnrollment, TaEnrollment, or DesignerEnrollment.
	DateOverrideEnrollmentType string `csv:"date_override_enrollment_type"`
	// [ sticky ] The date the term starts. The format should be in ISO 8601: YYYY-MM-DDTHH:MM:SSZ
	StartDate *time.Time `csv:"start_date"`
	// [ sticky ] The date the term ends. The format should be in ISO 8601: YYYY-MM-DDTHH:MM:SSZ
	EndDate *time.Time `csv:"end_date"`
}

type TermSerde struct {
	serde.Serde
}

func (sd TermSerde) Get() (v []Term, err error) {
EOF:
	for {
		var u Term
		switch err := sd.Decode(&u); {
		case err == io.EOF:
			break EOF
		case err != nil:
			return nil, err
		}
		v = append(v, u)
	}

	return v, nil
}

func (sd TermSerde) Set(v []Term) error {
	for _, u := range v {
		if err := sd.Encode(u); err != nil {
			return err
		}
	}

	return sd.Flush()
}
