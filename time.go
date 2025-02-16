package common

import "time"

const TimeLayout = time.RFC3339

func ParseTime(value string) time.Time {
	parsedTime, err := time.Parse(TimeLayout, value)
	if err != nil {
		return time.Now()
	}
	return parsedTime
}

func FormatTime(t time.Time) string {
	return t.Format(TimeLayout)
}
