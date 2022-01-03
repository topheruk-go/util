package model

import (
	"io"

	"github.com/topheruk/go/learn/data/csv/serde"
)

type UserObserver struct {
	// [ required ] The User identifier from users.csv for the observing user.
	ID string `csv:"observer_id"`
	// [ required ] The User identifier from users.csv for the student user.
	StudentID string `csv:"student_id"`
	// FIXME: [ required ] active, deleted
	Status string `csv:"status"`
}

type UserObserverSerde struct {
	serde.Serde
}

func (sd UserObserverSerde) Get() (v []UserObserver, err error) {
EOF:
	for {
		var u UserObserver
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

func (sd UserObserverSerde) Set(v []UserObserver) error {
	for _, u := range v {
		if err := sd.Encode(u); err != nil {
			return err
		}
	}

	return sd.Flush()
}
