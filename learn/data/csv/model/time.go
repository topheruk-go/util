package model

import "time"

// The time format is stated as "2006-01-02 15:04:05" but can change.
// Canvas requires "2006-01-02T15:04:05Z" so this wrapper is not really needed
// using "MarshalText()" & "UnmarshalText()" for continuity with std lib encoding package
// but could also use "csvutil" packages CSV equivelant
type Time struct {
	time.Time
}

const format = "2006-01-02 15:04:05"

func (t Time) MarshalText() ([]byte, error) {
	var b [len(format)]byte
	return t.AppendFormat(b[:0], format), nil
}

func (t *Time) UnmarshalText(data []byte) error {
	tt, err := time.Parse(format, string(data))
	if err != nil {
		return err
	}
	*t = Time{Time: tt}
	return nil
}
