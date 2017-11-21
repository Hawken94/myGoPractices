package algo

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	data := IntSlice{1, 6, 2, 0, 6, 5, 8, 3, 5, 4}
	fmt.Println(data)
	data.SelectSort()

	t.Errorf("\n xhk sorted:%v \n", data)
}
