package utils

import (
	"time"
)

func FormatFromUNIX(date int64) time.Time {
	format := time.Unix(date, 0)
	return format
}

func FormatPointerFromUNIX(date *int64) *time.Time {
	if date != nil {
		format := time.Unix(*date, 0)
		return &format
	}
	return nil
}

func FormatToUNIX(date time.Time) int64 {
	format := date.Unix()
	return format
}

func FormatPointerToUNIX(date *time.Time) *int64 {
	if date != nil {
		format := date.Unix()
		return &format
	}
	return nil
}
