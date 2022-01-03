package model

import (
	"io"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type XList struct {
	// [ required ] The course identifier from courses.csv
	CourseID string `csv:"xlist_course_id"`
	// [ required ] The section identifier from sections.csv
	SectionID string `csv:"section_id"`
	// FIXME: [ required ] active, deleted
	Status string `csv:"status"`
}

type XListSerde struct {
	serde.Serde
}

func (sd XListSerde) Get() (v []XList, err error) {
EOF:
	for {
		var u XList
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

func (sd XListSerde) Set(v []XList) error {
	for _, u := range v {
		if err := sd.Encode(u); err != nil {
			return err
		}
	}

	return sd.Flush()
}
