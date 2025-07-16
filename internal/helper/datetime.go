package helper

import "time"

func ParseDate(dateStr, format string) (time.Time, error) {
	parsedTime, err := time.Parse(format, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
