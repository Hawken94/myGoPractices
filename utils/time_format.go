package utils

import (
	"fmt"
	"time"
)

func formatTime(date time.Time) string {
	str := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(date)
	return str
}
