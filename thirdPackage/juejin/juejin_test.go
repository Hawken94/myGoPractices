package juejin

import (
	"testing"
)

func TestHappyNum(t *testing.T) {
	var tests = []struct {
		num int
	}{
		{1000000},
	}

	for _, test := range tests {
		t.Errorf("xhk happyNum: %v \n", getNHappyNum(test.num))
	}
}
