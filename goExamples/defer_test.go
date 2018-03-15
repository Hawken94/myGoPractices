package goExamples

import (
	"testing"
)

func TestDefer(t *testing.T) {
	t.Error(deferFunc1(1))
}
