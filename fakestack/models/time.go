package models

import (
	"encoding/json"
	"time"
)

const TimeFormat = "2006-01-02T15:04:05"

type Time time.Time

func (t *Time) UnmarshalJSON(buf []byte) error {
	var timeString string
	err := json.Unmarshal(buf, &timeString)
	if err != nil { return err }

	tt, err := time.Parse(TimeFormat, timeString)
	if err != nil { return err }

	*t = Time(tt)
	return nil
}

func ZeroDate() Time {
	time, err := time.Parse("2006-01-02T15:04:05", "1900-01-01T00:00:00.000")
	if err != nil {
		panic("Failed to parse constant time")
	}

	return Time(time)
}

