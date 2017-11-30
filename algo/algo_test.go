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
		// BubbleSort(test.data)
		test.data = MergeSort(test.data)
		// CocktailSort(test.data)
		// test.data = BucketSort(test.data)
		// test.data = CountingSort(test.data)

		t.Errorf("\n xhk sorted:%v \n", test.data)
	}

}

func TestBinSearch(t *testing.T) {

	arrs := []int{1, 2, 3, 4, 6, 6, 7, 8, 9}
	t.Errorf("xhk :%v \n", binarySearch(arrs, 0, len(arrs), 10))
}
