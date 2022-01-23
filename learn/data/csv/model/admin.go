package model

import (
	"io"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type Admin struct {
	UserID    string
	AccountID string
	RoleID    string
	Role      string
	// FIXME: enum:active,deleted
	Status      string
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
