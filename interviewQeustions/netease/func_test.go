package netease

import (
	"testing"
)

func TestOppositNum(t *testing.T) {
	var tests = []struct {
		num int
	}{
		{1325},
		{100},
	}

	for _, test := range tests {
		t.Errorf("OppositNum %v", OppositNum(test.num))
	}

}
