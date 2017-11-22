package algo

/*
冒泡排序算法：
比较相邻的元素。如果第一个比第二个大，就交换他们两个。
对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
针对所有的元素重复以上的步骤，除了最后一个。
持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
*/

// {1, 6, 2, 0, 6, 5, 8, 3, 5, 4}

// BubbleSort 冒泡排序
func BubbleSort(data []int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		isChange := false

		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}
}
