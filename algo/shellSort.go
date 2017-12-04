package algo

/*
步长的选择是希尔排序的重要部分。只要最终步长为1任何步长序列都可以工作。算法最开始以一定的步长进行排序。
然后会继续以一定步长进行排序，最终算法以步长为1进行排序。当步长为1时，算法变为插入排序，这就保证了数据一定会被排序。
*/

// ShellSort 希尔排序TODO: 有点复杂
func ShellSort(data []int) {
	n := len(data)
	if n < 2 {
		return
	}
	key := n / 2
	for ; key > 0; key = key / 2 {
		for i := key; i < n; i++ {
			j := i
			for j >= key && data[j] < data[j-key] {
				data[j], data[j-key] = data[j-key], data[j]
				j = j - key
			}
		}

		// 已知的最好步长序列是由Sedgewick提出的(1, 5, 19, 41, 109,...)
		// “比较在希尔排序中是最主要的操作，而不是交换。”
	}
}
