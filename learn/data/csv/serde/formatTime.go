package serde

import (
	"time"
)

func (c *CSV) formatTime(timeFormat string) {
	c.e.Register(func(t time.Time) ([]byte, error) {
		return t.AppendFormat(nil, timeFormat), nil
	})

	c.d.Register(func(data []byte, t *time.Time) error {
		tt, err := time.Parse(timeFormat, string(data))
		if err != nil {
			return err
		}
		*t = tt
		return nil
	})

	c.d.Map = func(field, column string, v interface{}) string {
		if _, ok := v.(time.Time); ok && field == "" {
			return time.Time{}.Format(timeFormat)
		}
		return field
	}
}
