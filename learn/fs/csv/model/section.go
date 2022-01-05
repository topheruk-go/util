package model

import (
	"io"
	"time"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type Section struct {
	ID            string     `csv:"id"`
	CourseID      string     `csv:"course_id"`
	Name          string     `csv:"name"`
	Status        Status     `csv:"status"`
	IntegrationID string     `csv:"integration_id"`
	StartDate     *time.Time `csv:"start_date"`
	EndDate       *time.Time `csv:"end_date"`
}

type SectionSerde struct {
	serde.Serde
}

func (sd SectionSerde) Get() (v []Section, err error) {
EOF:
	for {
		var u Section
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

func (sd SectionSerde) Set(v []Section) error {
	for _, u := range v {
		if err := sd.Encode(u); err != nil {
			return err
		}
	}

	return sd.Flush()
}
