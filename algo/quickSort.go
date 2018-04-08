package algo

/*
1.从数列中挑出一个元素，称为"基准"（pivot），
2.重新排序数列，所有比基准值小的元素摆放在基准前面，所有比基准值大的元素摆在基准后面（相同的数可以到任何一边）。在这个分
割结束之后，该基准就处于数列的中间位置。这个称为分割（partition）操作。
3.递归地（recursively）把小于基准值元素的子数列和大于基准值元素的子数列排序。

递归到最底部时，数列的大小是零或一，也就是已经排序好了。这个演算法一定会结束，因为在每次的迭代（iteration）中，它至少
会把一个元素摆到它最后的位置去。
*/
/*
最坏时间复杂度	\Theta (n^{2})
最优时间复杂度	\Theta (n\log n)
平均时间复杂度	\Theta (n\log n)
空间复杂度	根据实现的方式不同而不同
*/
func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivotPos := partition2(arr, left, right)
	quickSort(arr, left, pivotPos-1)
	quickSort(arr, pivotPos+1, right)
}

func partition(arr []int, left, right int) int {
	pivotKey := arr[left]
	pivotPoint := left

	for left < right {
		for left < right && arr[right] >= pivotKey {
			right--
		}

		for left < right && arr[left] <= pivotKey {
			left++
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[left], arr[pivotPoint] = arr[pivotPoint], arr[left]
	return left
}

// 优化过的
func partition2(arr []int, left, right int) int {
	pivotKey := arr[left]
	// pivotPoint := left

	for left < right {
		for left < right && arr[right] >= pivotKey {
			right--
		}
		arr[left] = arr[right] // 把小的移动到左边
		for left < right && arr[left] <= pivotKey {
			left++
		}
		arr[right] = arr[left] // 把大的移动到右边
	}
	arr[left] = pivotKey
	return left
}
