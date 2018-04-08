package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	num := 0
	for {
		n, _ := fmt.Scan(&num)
		if n == 0 {
			break
		} else {
			fmt.Printf("请输入%d个数字\n", num)
			nums := make([]int, 0)
			for i := 0; i <= num; i++ {
				n, _ := fmt.Scan(&i)
				if n == 0 {
					break
				} else {
					nums = append(nums, i)
					fmt.Println(nums)
				}
			}
			// 列出所有排列组合
			slices := permute(nums)
			tmp := slicesTran(slices)
			// 为了方便排序
			var sortSlice sort.IntSlice
			sortSlice = tmp

			sort.Sort(sortSlice)
			fmt.Println(sortSlice)
		}

	}

}

// 排列组合TODO: 传说中的回溯法
func permute(nums []int) [][]int {
	slices := make([][]int, 0)
	backtrackPer(&slices, []int{}, nums)
	return slices
}

func backtrackPer(slices *[][]int, temp []int, nums []int) {
	if len(temp) == len(nums) {
		// slice是引用类型,这里用临时变量防止结果被修改
		tempSlice := make([]int, len(temp))
		copy(tempSlice, temp)
		*slices = append(*slices, tempSlice)
	} else {
		for i := 0; i < len(nums); i++ {
			// temp包含nums[i]就跳过循环
			if isContain(temp, nums[i]) {
				continue
			}
			temp = append(temp, nums[i])
			backtrackPer(slices, temp, nums)
			temp = temp[:len(temp)-1]
		}
	}
}
func isContain(nums []int, num int) bool {
	for _, v := range nums {
		if v == num {
			return true
		}
	}
	return false
}

// slicesTran 多维切片转换成一维切片
func slicesTran(slices [][]int) []int {
	result := make([]int, 0)

	for _, v := range slices {
		tmp := sliceToNum(v)

		result = append(result, tmp)
	}
	return result
}

// sliceToNum 数组转换成对应的数字
func sliceToNum(slice []int) int {
	reuslt := 0
	for i := len(slice) - 1; i >= 0; i-- {
		reuslt += slice[i] * int(math.Pow10(i))
	}
	return reuslt
}
