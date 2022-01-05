package model

import (
	"io"

	"github.com/topheruk/go/learn/fs/csv/serde"
)

type Account struct {
	ID       string `csv:"account_id"`
	ParentID string `csv:"parent_account_id"`
	Name     string `csv:"name"`
	// FIXME: enum:active,deleted
	Status        Status `csv:"status"`
	IntegrationID string `csv:"integration_id"`
}

type AccountSerde struct {
	serde.Serde
}

func (sd AccountSerde) Get() (v []Account, err error) {
	// sd.Map(func(field, col string, v interface{}) string {
	// 	return field
	// })
	// var x int

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
