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

func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, GetDefaultTimezone())
}

func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, GetDefaultTimezone())
}

func SetHour(t time.Time, hour int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hour, t.Minute(), t.Second(), 0, GetDefaultTimezone())
}

func SetMinute(t time.Time, minute int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minute, t.Second(), 0, GetDefaultTimezone())
}

func DateTimeToInt(dt time.Time) int {
	return int(dt.Unix())
}

func MillisecondsToTime(ms int64) time.Time {
	seconds := ms / 1000
	nanoseconds := (ms % 1000) * 1000000
	return time.Unix(seconds, nanoseconds)
}

func MicrosecondsToTime(ms int64) time.Time {
	seconds := ms / 1000000
	nanoseconds := (ms % 1000000) * 1000
	return time.Unix(seconds, nanoseconds)
}
