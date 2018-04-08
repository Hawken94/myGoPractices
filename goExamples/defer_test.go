package goExamples

import (
	"testing"
)

func TestDefer(t1 *testing.T) {
	t1.Error(deferFunc3(1))
}
