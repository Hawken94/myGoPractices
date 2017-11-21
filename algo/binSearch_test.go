package algo

import (
	"testing"
)

func TestBinSearch(t *testing.T) {

	arrs := []int{1, 2, 3, 4, 4, 6, 7, 8, 9}
	t.Errorf("xhk :%v \n", binarySearch(arrs, 0, len(arrs), 5))
}
