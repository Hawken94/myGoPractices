package algo

import (
	"sort"
)

type IntSlice []int

func selectSort(data sort.Interface) {
	for i := 0; i < data.Len()-1; i++ {
		min := i
		for j := min + i; j < data.Len(); j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(min, i)
	}
}

// SelectSort 选择排序
func (data IntSlice) SelectSort() {
	min := 0
	for i := 0; i < len(data)-1; i++ {
		min = i
		for j := min + 1; j < len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		data[min], data[i] = data[i], data[min]
	}
}
