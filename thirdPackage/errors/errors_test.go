package errors

import (
	errs "errors"
	"fmt"
	"runtime"
	"testing"

	"github.com/pkg/errors"
)

func TestErrors(t *testing.T) {
	err := errs.New("this is a test")
	err2 := errs.New("this is a test")
	ret := errors.Wrap(err, "test errors")
	ret = errors.Wrap(err2, "2222222")
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
