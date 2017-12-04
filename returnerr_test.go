package main

import (
	"testing"
)

func TestReturnErr(t *testing.T) {
	t.Errorf("xhk:%v\n", doStuff())
}
