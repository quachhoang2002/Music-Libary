package util

import (
	"time"
)

const (
	DateTimeFormat = "2006-01-02 15:04:05"
)

func StrToDateTime(str string) (time.Time, error) {
	t, err := time.Parse(DateTimeFormat, str)
	if err != nil {
		return time.Time{}, err
	}
	t = t.Add(-time.Duration(7) * time.Hour)
	return t.In(GetDefaultTimezone()), nil
}

func DateTimeToStr(dt time.Time, ft *string) string {
	if ft == nil {
		return dt.Format(DateTimeFormat)
	} else {
		return dt.Format(*ft)
	}
}

func Now() time.Time {
	return time.Now().In(GetDefaultTimezone())
}

func GetDefaultTimezone() *time.Location {
	localTimeZone, _ := time.LoadLocation("Local")
	return localTimeZone
}
