package types

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"
)

// Time ...
type Time struct {
	time.Time
}

// NewTime ...
func NewTime(t time.Time) Time {
	return Time{t}
}

// MarshalJSON ...
func (t *Time) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(t.Time)
	if err != nil {
		return nil, err
	}

	data = bytes.TrimRight(data, `"`)
	data = bytes.TrimRight(data, "Z")
	data = append(data, '"')

	return data, nil
}

// UnmarshalJSON ...
func (t *Time) UnmarshalJSON(data []byte) error {
	d := string(data)
	d = strings.Trim(d, "\"")
	d = strings.TrimRight(d, "Z")

	format := "2006-01-02T15:04:05.000"
	tTime, err := time.Parse(format, d)
	if err == nil {
		t.Time = tTime

		return nil
	}

	format = "2006-01-02T15:04:05"
	tTime, err = time.Parse(format, d)
	if err != nil {
		return err
	}

	t.Time = tTime

	return nil
}
