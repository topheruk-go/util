package model

import (
	"io"
	"time"

	"github.com/topheruk/go/learn/fs/csv/serde"
)

type Enrollment struct {
	CourseID              string     `csv:"course_id"`
	RootAccount           string     `csv:"root_account"`
	StartDate             *time.Time `csv:"start_date"`
	EndDate               *time.Time `csv:"end_date"`
	UserID                string     `csv:"user_id"`
	UserIntergrationID    string     `csv:"user_integration_id"`
	Role                  string     `csv:"role"`
	RoleID                string     `csv:"role_id"`
	SectionID             string     `csv:"section_id"`
	Status                Status     `csv:"status"`
	AssociatedUserID      string     `csv:"associated_user_id"`
	LimitSectionPrivleges bool       `csv:"limit_section_privileges"`
	Notify                bool       `csv:"notify"`
}

type EnrollmentSerde struct {
	serde.Serde
}

func (sd EnrollmentSerde) Get() (v []Enrollment, err error) {
EOF:
	for {
		var u Enrollment
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

func (sd EnrollmentSerde) Set(v []Enrollment) error {
	for _, u := range v {
		if err := sd.Encode(u); err != nil {
			return err
		}
	}

	return sd.Flush()
}
