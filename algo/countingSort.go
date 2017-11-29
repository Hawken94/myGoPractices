package algo

/*
1.找出待排序的数组中最大和最小的元素
2.统计数组中每个值为i的元素出现的次数，存入数组 C 的第 i 项
3.对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）
4.反向填充目标数组：将每个元素i放在新数组的第C(i)项，每放一个元素就将C(i)减去1
*/

// CountingSort 计数排序
func CountingSort(data []int) []int {
	b := make([]int, len(data))
	max, min := data[0], data[0]
	for _, d := range data {
		if max < d {
			max = d
		}
		if min > d {
			min = d
		}
	}

	// 这里k的大小是要排序的数组中，元素大小的极值差+1
	k := max - min + 1
	c := make([]int, k)
	for i := 0; i < len(data); i++ {
		c[data[i]-min]++ // 优化过的地方，减小了数组c的大小
	}
	for i := 1; i < len(c); i++ {
		c[i] = c[i] + c[i-1]
	}
	for i := len(data) - 1; i >= 0; i-- {
		c[data[i]-min]--            // 考虑重复
		b[c[data[i]-min]] = data[i] // 按存取的方式取出c的元素
	}
	return b
}
