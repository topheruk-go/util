package model

import (
	"io"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type Group struct {
	ID               string `csv:"group_id"`
	GroupCatergoryID string `csv:"group_category_id"`
	AccountID        string `csv:"account_id"`
	CourseID         string `csv:"course_id"`
	Name             string `csv:"name"`
	Status           Status `csv:"status"`
}

type GroupSerder struct {
	serde.Serde
}

func (sd GroupSerder) Get() (v []Group, err error) {
EOF:
	for {
		var u Group
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

func (sd GroupSerder) Set(v []Group) error {
	for _, u := range v {
		if err := sd.Encode(u); err != nil {
			return err
		}
	}

	return sd.Flush()
}

type GroupMembership struct {
	GroupID string `csv:"group_id"`
	UserID  string `csv:"user_id"`
	Status  Status `csv:"status"`
}

type GroupMembershipSerder struct {
	serde.Serde
}

func (sd GroupMembershipSerder) Get() (v []GroupMembership, err error) {
EOF:
	for {
		var u GroupMembership
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

func (sd GroupMembershipSerder) Set(v []GroupMembership) error {
	for _, u := range v {
		if err := sd.Encode(u); err != nil {
			return err
		}
	}

	return sd.Flush()
}

type GroupCategory struct {
	ID        string `csv:"group_category_id"`
	AccountID string `csv:"account_id"`
	CourseID  string `csv:"course_id"`
	Name      string `csv:"category_name"`
	Status    Status `csv:"status"`
}

type GroupCategorySerder struct {
	serde.Serde
}

func (sd GroupCategorySerder) Get() (v []GroupCategory, err error) {
EOF:
	for {
		var u GroupCategory
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

func (sd GroupCategorySerder) Set(v []GroupCategory) error {
	for _, u := range v {
		if err := sd.Encode(u); err != nil {
			return err
		}
	}

	return sd.Flush()
}
