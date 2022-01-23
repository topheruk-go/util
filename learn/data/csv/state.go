package main

import (
	"time"

	"github.com/topheruk/go/learn/data/csv/model"
)

// TODO: can I turn this into a stateFn?

// type stateFn func(field, col string, v interface{}) string

func state(field, col string, v interface{}) string {
	// fmt.Println(field, col, v)

	switch v.(type) {
	case string:
	case int:
	case time.Time:
	case model.Status:
		return stateStatus(field, col, v.(model.Status))
	}
	return field
}

func stateStatus(field, col string, s model.Status) string {
	field = model.Deleted.String()
	return field
}
