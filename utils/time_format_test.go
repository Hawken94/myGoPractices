package utils

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	now := time.Now()

	after14days := now.AddDate(0, -2, 0)
	t.Errorf("xhk format_time :%v %v \n", formatTime(now), formatTime(after14days))

}
