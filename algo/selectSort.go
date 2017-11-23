package algo

import (
	"sort"
)

/*
1.首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，
2.然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。

选择排序的主要优点与数据移动有关。如果某个元素位于正确的最终位置上，则它不会被移动。选择排序每次交换一对元素，
它们当中至少有一个将被移到其最终位置上，因此对 n个元素的表进行排序总共进行至多n-1次交换。
在所有的完全依靠交换去移动元素的排序方法中，选择排序属于非常好的一种
*/
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
func SelectSort(data []int) {
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
