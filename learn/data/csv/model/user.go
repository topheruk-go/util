package model

import (
	"io"

	"github.com/topheruk/go/learn/data/csv/serde"
)

// canvas_user_id,
// user_id,
// integration_id
// authentication_provider_id
// login_id
// first_name
// last_name,
// full_name,
// sortable_name,
// short_name,
// email,status,
// created_by_sis

// canvas_user_id,user_id,integration_id,authentication_provider_id,login_id,first_name,last_name,full_name,sortable_name,short_name,email,status,created_by_sis
// 2572,,,,Chelsea.Aagaard@stu.fra.ac.uk,Chelsea,Aagaard,Chelsea Aagaard,"Aagaard, Chelsea",Chelsea,Chelsea.Aagaard@stu.fra.ac.uk,active,false

type User struct {
	ID                       string `csv:"user_id"`
	IntegrationID            string `csv:"intergration_id"`
	LoginID                  string `csv:"login_id"`
	Password                 string `csv:"password"`
	SSHA                     string `csv:"ssha_password"`
	AuthenticationProviderID string `csv:"authentication_provider_id"`
	FirstName                string `csv:"first_name"`
	LastName                 string `csv:"last_name"`
	FullName                 string `csv:"full_name"`
	SortableName             string `csv:"sortable_name"`
	ShortName                string `csv:"short_name"`
	Email                    string `csv:"email"`
	Pronouns                 string `csv:"pronouns"`
	DeclaredUserType         Status `csv:"declared_user_type"`
	Status                   Status `csv:"status"`
}

type UserSerde struct {
	serde.Serde
}

func (sd UserSerde) Get() (v []User, err error) {
EOF:
	for {
		var u User
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

func (sd UserSerde) Set(v []User) error {
	for _, u := range v {
		if err := sd.Encode(u); err != nil {
			return err
		}
	}

	return sd.Flush()
}
