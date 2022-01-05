package model

import (
	"io"
	"time"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type Term struct {
	ID                         string     `csv:"term_id"`
	Name                       string     `csv:"name"`
	Status                     Status     `csv:"status"`
	IntegrationID              string     `csv:"integration_id"`
	DateOverrideEnrollmentType string     `csv:"date_override_enrollment_type"`
	StartDate                  *time.Time `csv:"start_date"`
	EndDate                    *time.Time `csv:"end_date"`
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
