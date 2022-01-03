// course_id,short_name,long_name,account_id,term_id,status

package model

import (
	"io"
	"time"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type Course struct {
	// [ required ] A unique identifier used to reference courses in the enrollments data. This identifier must not change for the account, and must be globally unique. In the user interface, this is called the SIS ID.
	ID string `csv:"course_id"`
	// [ required, sticky ] A long name for the course. (This can be the same as the short name, but if both are available, it will provide a better user experience to provide both.)
	Name string `csv:"long_name"`
	// [ required, sticky ] A short name for the course
	ShortName string `csv:"short_name"`
	// [ sticky ] The account identifier from accounts.csv. New courses will be attached to the root account if not specified here
	AccountID string `csv:"account_id"`
	// [ sticky ] The term identifier from terms.csv, if no term_id is specified the default term for the account will be used
	TermID string `csv:"term_id"`
	// FIXME: [ required, sticky ] active, deleted, completed, published
	Status string `csv:"status"`
	// Sets the integration_id of the course
	IntegrationID string `csv:"integration_id"`
	// [ sticky ] he course start date. The format should be in ISO 8601: YYYY-MM-DDTHH:MM:SSZ. To remove the start date pass "<delete>"
	StartDate *time.Time `csv:"start_date"`
	// [ sticky ] he course start date. The format should be in ISO 8601: YYYY-MM-DDTHH:MM:SSZ. To remove the start date pass "<delete>"
	EndDate *time.Time `csv:"end_date"`
	// FIXME: on_campus, online, blended
	CourseFormat string `csv:"course_format"`
	// The SIS id of a pre-existing Blueprint course. When provided, the current course will be set up to receive updates from the blueprint course. Requires Blueprint Courses feature. To remove the Blueprint Course link you can pass 'dissociate' in place of the id.
	BlueprintCourseID string `csv:"blueprint_course_id"`
	// [ stricky ] nightly_sync, not_set
	GradePassbackSetting string `csv:"grade_passback_setting"`
	// Whether the course is a homeroom course. Requires the courses to be associated with a "Canvas for Elementary"-enabled account.
	HomeroomCourse bool `csv:"homeroom_course"`
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
