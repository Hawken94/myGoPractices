package timingwheel

import (
	"fmt"
	"testing"
	"time"
)

func TestStartTimer(t *testing.T) {
	tw := NewTimingWheel(time.Millisecond, 20)
	tw.Start()
	defer tw.Stop()

	exitC := make(chan time.Time, 1)
	tw.AfterFunc(time.Second, func() {
		fmt.Println("The timer fires")
		exitC <- time.Now()
	})
	t.Error("------------------")
}
