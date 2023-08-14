package helper

import "time"

type EpochSeconds int64

const minutesInAnHour = 9 * 60 * 60

// ConvertToISO8601 converts EpochSeconds to ISO8601 format.
func (e EpochSeconds) ConvertToISO8601() string {
	jstLocation := time.FixedZone("JST", minutesInAnHour)
	jstTime := time.Unix(int64(e), 0).In(jstLocation)
	iso8601Time := jstTime.Format("2006-01-02 15:04:05")

	return iso8601Time
}
