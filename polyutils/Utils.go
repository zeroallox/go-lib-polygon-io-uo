package polyutils

import (
	"time"
)

var yyyyMMddFormat = "2006-01-02"

// TimeToStringDate returns a string from date formatted as YYYY-MM-DD ("2006-01-02")
func TimeToStringDate(date time.Time) string {
	return date.Format(yyyyMMddFormat)
}

// TimeFromStringDate returns a time.Time from a string YYYY-MM-DD format ("2006-01-02")
func TimeFromStringDate(str string, loc *time.Location) (time.Time, error) {

	if loc == nil {
		loc = time.UTC
	}

	tm, err := time.ParseInLocation(yyyyMMddFormat, str, loc)
	if err != nil {
		return time.Time{}, err
	}

	return tm, nil
}
