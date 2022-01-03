package model

import (
	"io"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type Admin struct {
	// [ required ] The User identifier from users.csv
	UserID string
	// [ required ] The account identifier from accounts.csv. Uses the root_account if left blank. The collumn is required even when importing for the root_account and the value is blank.
	AccountID string
	// [ required (or "role") ] Uses a role id, either built-in or defined by the account.
	RoleID string
	// [ required (or "role_id") ] AccountAdmin, or a custom role defined by the account. When using a custom role, the name is case sensitive.
	Role string
	// FIXME: [ requried ] active, deleted
	Status string
	// The domain of the account to search for the user.
	RootAccount string
}

type AdminSerde struct {
	serde.Serde
}

func (sd AdminSerde) Get() (v []Admin, err error) {
EOF:
	for {
		var u Admin
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

func (sd AdminSerde) Set(v []Admin) error {
	for _, u := range v {
		if err := sd.Encode(u); err != nil {
			return err
		}
	}

	return sd.Flush()
}
