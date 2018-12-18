package thirdPackage

import (
	"testing"
)

func UseMutex(errs string) {
	mutex.Lock()
	errs = "true"
	mutex.Unlock()
}
func UseChan(ch chan string) {
	ch <- "true"
	<-ch
}

var errs string
var ch chan string

func BenchmarkUseMutex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UseMutex()
	}
}
func BenchmarkUseChan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UseChan()
	}
}
