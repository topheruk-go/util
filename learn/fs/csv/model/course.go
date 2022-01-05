// course_id,short_name,long_name,account_id,term_id,status

package model

import (
	"io"
	"time"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type Course struct {
	ID                   string     `csv:"course_id"`
	Name                 string     `csv:"long_name"`
	ShortName            string     `csv:"short_name"`
	AccountID            string     `csv:"account_id"`
	TermID               string     `csv:"term_id"`
	Status               Status     `csv:"status"`
	IntegrationID        string     `csv:"integration_id"`
	StartDate            *time.Time `csv:"start_date"`
	EndDate              *time.Time `csv:"end_date"`
	CourseFormat         Status     `csv:"course_format"`
	BlueprintCourseID    string     `csv:"blueprint_course_id"`
	GradePassbackSetting string     `csv:"grade_passback_setting"`
	HomeroomCourse       bool       `csv:"homeroom_course"`
}

type CourseSerde struct {
	serde.Serde
}

func (sd CourseSerde) Get() (v []Course, err error) {
EOF:
	for {
		var u Course
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

func (sd CourseSerde) Set(v []Account) error {
	for _, u := range v {
		if err := sd.Encode(u); err != nil {
			return err
		}
	}

	return sd.Flush()
}
