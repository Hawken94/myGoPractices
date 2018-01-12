package interest

import (
	"testing"
)

func TestFindMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
	}{
		{
			[][]int{
				{1, 2, 8, 9},
				{2, 4, 9, 12},
				{4, 7, 10, 13},
				{6, 8, 11, 15},
			},
			71,
		},
	}

	for _, test := range tests {
		t.Errorf("findMatrix :%v \n", findMatrix(test.matrix, test.target))
	}
}
