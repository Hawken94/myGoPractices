package errors

import (
	"fmt"
	"myGoPractices/thirdPackage/errors/logic"
	"runtime"
	"testing"
)

func TestErrors(t *testing.T) {
	ret := logic.ReturnErr()
	t.Error(ret)
}

func TestGC(t *testing.T) {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ {
		}
	}()
	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")
	t.Error("--------")
}
