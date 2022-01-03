package model

import (
	"io"
	"time"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type Section struct {
	// [Requried] A unique identifier used to reference sections in the enrollments data. This identifier must not change for the section, and must be globally unique. In the user interface, this is called the SIS ID.
	ID string `csv:"id"`
	// [Requried, Sticky] The course identfier from course.csv
	CourseID string `csv:"course_id"`
	// [Requried, Sticky] The name of the section
	Name string `csv:"name"`
	// FIXME: [Requried] active, deleted
	Status string `csv:"status"`
	// Sets the integration_id of the section
	IntegrationID string `csv:"integration_id"`
	// [Sticky] The section start date The format should be in ISO 8601: YYYY-MM-DDTHH:MM:SSZ
	StartDate *time.Time `csv:"start_date"`
	// [Sticky] The section end date The format should be in ISO 8601: YYYY-MM-DDTHH:MM:SSZ
	EndDate *time.Time `csv:"end_date"`
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
