package tests

import (
	"fmt"
	"testing"
	"time"

	"gitlab.gumpcome.com/common/tools/logkit"
)

func init() {
	logkit.InitGraylog("", "")
}
func BenchmarkError(b *testing.B) {
	var error1, error2, errorf int64
	now1 := time.Now()
	for i := 0; i < b.N; i++ {
		logkit.Error("This is a test" + "for something")
	}
	error1 = time.Since(now1).Nanoseconds()

	now2 := time.Now()
	for i := 0; i < b.N; i++ {
		logkit.Error("This is a test", "for something")
	}
	error2 = time.Since(now2).Nanoseconds()

	now3 := time.Now()
	for i := 0; i < b.N; i++ {
		logkit.Errorf("This is a test %v", "for something")
	}
	errorf = time.Since(now3).Nanoseconds()

	fmt.Println("err1:", error1)
	fmt.Println("err2:", error2)
	fmt.Println("errf:", errorf)
}
func BenchmarkError2(b *testing.B) {
	for i := 0; i < 100000; i++ {
		logkit.Error("This is a test", "for something")
	}
}

func BenchmarkErrorf(b *testing.B) {
	for i := 0; i < 100000; i++ {
		logkit.Errorf("This is a test %s", "for something")
	}
}
