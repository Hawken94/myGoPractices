package algo

import (
	"testing"
)

func TestSelectSort(t *testing.T) {
	var tests = []struct {
		data []int
	}{
		{[]int{1, 6, 2, 0, 6, 5, 8, 3, 5, 4}},
		{[]int{12, 324, 134, 3, 12, 233, 123, 34, 67, 45, 0}},
	}

	for _, test := range tests {
		t.Errorf("\n xhk origin:%v \n", test.data)
		// SelectSort(data)
		// InserSort(test.data)
		BubbleSort(test.data)

		t.Errorf("\n xhk sorted:%v \n", test.data)
	}

}
