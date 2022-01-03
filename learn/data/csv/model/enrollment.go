package model

import (
	"io"
	"time"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type Enrollment struct {
	// [ required (or section_id) ] The course identifier from courses.csv
	CourseID string `csv:"course_id"`
	// The domain of the account to search for the user.
	RootAccount string `csv:"root_account"`
	// [ sticky ] The enrollment start date. For start_date to take effect the end_date also needs to be populated. The format should be in ISO 8601: YYYY-MM-DDTHH:MM:SSZ
	StartDate *time.Time `csv:"start_date"`
	// [ sticky ] The enrollment end date. For end_date to take effect the start_date also needs to be populated. The format should be in ISO 8601: YYYY-MM-DDTHH:MM:SSZ
	EndDate *time.Time `csv:"end_date"`
	// [ required (or user_integration_id) ] The User identifier from users.csv, required to identify user. If the user_integration_id is present, this field will be ignored.
	UserID string `csv:"user_id"`
	// [ required (or user_id) ]The integration_id of the user from users.csv required to identify user if the user_id is not present.
	UserIntergrationID string `csv:"user_integration_id"`
	// [ required (or role_id) ] student, teacher, ta, observer, designer, or a custom role defined by the account. When using a custom role, the name is case sensitive.
	Role string `csv:"role"`
	// [ required (or role) ]Uses a role id, either built-in or defined by the account
	RoleID string `csv:"role_id"`
	// [ required (or course_id) ]The section identifier from sections.csv, if none is specified the default section for the course will be used
	SectionID string `csv:"section_id"`
	// active, deleted, completed, inactive, deleted_last_completed
	Status string `csv:"status"`
	// For observers, the user identifier from users.csv of a student in the same course that this observer should be able to see grades for. Ignored for any role other than observer
	AssociatedUserID string `csv:"associated_user_id"`
	// Defaults to false. When true, the enrollment will only allow the user to see and interact with users enrolled in the section given by course_section_id.
	LimitSectionPrivleges bool `csv:"limit_section_privileges"`
	// If true, a notification will be sent to the enrolled user. Notifications are not sent by default.
	Notify bool `csv:"notify"`
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
