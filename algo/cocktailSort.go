package algo

/*
鸡尾酒排序等于是冒泡排序的轻微变形。不同的地方在于从低到高然后从高到低，而冒泡排序则仅从低到高去比较序列里的每个元素。
他可以得到比冒泡排序稍微好一点的效能，原因是冒泡排序只从一个方向进行比对（由低到高），每次循环只移动一个项目。
以序列(2,3,4,5,1)为例，鸡尾酒排序只需要访问一次序列就可以完成排序，但如果使用冒泡排序则需要四次。但是在乱数序列的状态下，鸡尾酒排序与冒泡排序的效率都很差劲.
*/

// CocktailSort 鸡尾酒排序
func CocktailSort(data []int) {
	i, left := 0, 0
	right := len(data) - 1
	for left < right {
		for i = left; i < right; i++ {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
			}
		}
		right--
		for i = right; i > left; i-- {
			if data[i-1] > data[i] {
				data[i-1], data[i] = data[i], data[i-1]
			}
		}
		left++
	}
}
