package zhuozhuo

import (
	"fmt"
)

// 1.给定的字符串，找出其中最长的连续没有重复的字符串，并给出它的长度
// 例如：
// 输入：“aabcceddabc” 输出：“dabc” 4

// getLongestSubstrings ... byte是go里面的字节
// 思路：利用map查询，string的[i]值为key，i为value，值相同则进行max比较，取最大
func getLongestSubstrings(s string) (string, int) {
	// TODO:该方法只能返回它的长度,不能返回正确的子字符串
	if len(s) == 0 {
		return "", 0
	}
	tempMap := make(map[byte]int)
	var j, maxNum int
	// var flag bool
	var index int
	for i := 0; i < len(s); i++ {
		if _, ok := tempMap[s[i]]; ok {
			j = max(j, tempMap[s[i]]+1)
		}
		// 将字符存到 tempMap 中,同时也有更新值的功能
		tempMap[s[i]] = i

		// 这里有问题!!!
		if maxNum >= i-j+1 {
			index++
		}

		maxNum = max(maxNum, i-j+1)

	}
	fmt.Println(index)
	return s[index : index+maxNum], maxNum
}

// 2、合并两个升序排序的int的slice
// 例如：
// 输入：[1,3,4]  [1,1,2,5] 输出：[1,1,1,2,3,5]

// merge 合并已排序数组
// 如果不返回一个引用就可以使用一个临时变量了
func merge(nums1 []int, nums2 []int) []int {
	i, j := len(nums1)-1, len(nums2)-1
	k := i + j + 1
	// 建立一个长度为 nums1与 nums2之和的切片
	result := make([]int, k+1, k+1)

	// 从后往前遍历切片
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			result[k] = nums1[i]
			i--
		} else {
			result[k] = nums2[j]
			j--
		}
		k--
	}
	// 如果 nums1有剩余,则补到 result 切片中
	for i >= 0 && k >= 0 {
		result[k] = nums1[i]
		k--
		i--
	}
	// 如果 nums2有剩余,则补到 result 切片中
	for j >= 0 && k >= 0 {
		result[k] = nums2[j]
		k--
		j--
	}

	return result
}

// 3、非负整数的slice，每一位上的数字代表它所能向前前进的最大步数，从第一位开始计算，判断是否能够到达最后一个节点
// 例如：
// 输入:[1,4,1,1,4]  输出：true
// 输入:[3,2,1,0,4]  输出：false
func canJump(nums []int) bool {
	reachable := 0
	for i := 0; i < len(nums); i++ {
		// 如果当前的序号大于之前的最大步数,说明不可到达
		if i > reachable {
			return false
		}
		// 最大可到达的位数
		reachable = max(reachable, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
