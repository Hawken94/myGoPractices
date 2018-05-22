package utils

import (
	"time"
)

func formatTime(date time.Time) string {
	str := date.Format("2006-01-02 15:04:05")
	// fmt.Println(date)
	return str
}
