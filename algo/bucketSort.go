package algo

/*
桶排序（Bucket sort）或所谓的箱排序，是一个排序算法，工作的原理是将数组分到有限数量的桶里。每个桶再个别排序（有可能再使用别的
排序算法或是以递归方式继续使用桶排序进行排序）。桶排序是鸽巢排序的一种归纳结果。当要被排序的数组内的数值是均匀分配的时候，
桶排序使用线性时间（O(n)}（大O符号））。但桶排序并不是比较排序，他不受到 O(nlog n)下限的影响。

桶排序以下列程序进行：
1.设置一个定量的数组当作空桶子。
2.寻访序列，并且把项目一个一个放到对应的桶子去。
3.对每个不是空的桶子进行排序。
4.从不是空的桶子里把项目再放回原来的序列中。
*/

// 把待排序的数组里面的数在与之相等的桶的索引值，对应的桶里面的数据加1
// 1、产生随机数
// 2、创建对应的桶
// 3、桶和数据对应起来
// 4、相对应的桶数值加一
// 5、打印桶的索引值

// BucketSort 桶排序
func BucketSort(data []int) []int {

	max := data[0]
	for _, d := range data {
		if max < d {
			max = d
		}
	}
	// 分配桶的个数,此排序范围和data的最大值有关
	sorted := make([]int, max+1)

	// 将待排序数组里面的数放到与sorted数组相对于的桶里面
	// 待排序的数等于桶排序的索引值，则对应的桶 sorted[data[i]]数值加1
	for i := 0; i < len(data); i++ {
		sorted[data[i]] = sorted[data[i]] + 1
	}
	data = make([]int, 0)
	for i := 0; i < len(sorted); i++ {
		// 打印出数值不为0的桶的索引值;缺陷:非负整数
		if sorted[i] > 0 {
			// 判断桶里面有多少个数，然后全部打印出来
			for j := 0; j < sorted[i]; j++ {
				// fmt.Print(i)
				data = append(data, i)
			}
		}
	}
	return data
}
