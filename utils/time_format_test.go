package utils

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {

	t.Errorf("xhk format_time :%v \n", formatTime(time.Now()))
}
