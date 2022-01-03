package model

import (
	"io"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type Account struct {
	// [ required ] A unique identifier used to reference accounts in the enrollments data. This identifier must not change for the account, and must be globally unique. In the user interface, this is called the SIS ID.
	ID string `csv:"account_id"`
	// [ required, sticky ] The account identifier of the parent account. If this is blank the parent account will be the root account. Note that even if all values are blank, the column must be included to differentiate the file from a group import.
	ParentID string `csv:"parent_account_id"`
	// [ required, sticky ] The name of the account
	Name string `csv:"name"`
	// [ required ] active, deleted
	Status string `csv:"status"`
	// Sets the integration_id of the account
	IntegrationID string `csv:"integration_id"`
}

type AccountSerde struct {
	serde.Serde
}

func (sd AccountSerde) Get() (v []Account, err error) {
EOF:
	for {
		var u Account
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

func (sd AccountSerde) Set(v []Account) error {
	for _, u := range v {
		if err := sd.Encode(u); err != nil {
			return err
		}
	}

	return sd.Flush()
}
