package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	var str string
	fmt.Println("Please input numbers:")
	// 获取命令行的输入
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadBytes('\n')
	str = string(input[0 : len(input)-1])

	if len(str) < 1 || len(str) > 1000 {
		fmt.Println("输入的长度非法!")
		return
	}

	// 将字符串转换成数组
	intSlice, err := strToSlice(str)
	if err != nil {
		fmt.Println("输入的字符含有非数字!")
		return
	}

	result := 0

	// 对数组进行排序
	quickSort(intSlice, 0, len(intSlice)-1)

	// 判断字符串是否完全包含了0-9
	// 如果完全包含了0-9,则最小值为数组元素组成的最大值加1
	// 如果没有,则为缺少的那个最小数
	tmp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if arrFullContain(intSlice, tmp) {
		// 完全包含了0-9
		result = sliceToNumber(intSlice) + 1
	} else {
		// 不完全包含0-9
		tmp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		tmpNum := arrNotFullContain(intSlice, tmp)
		if arrFullContainNine(intSlice) {
			result = 10
		} else {
			result = tmpNum
		}
	}

	fmt.Printf("the result is %v \n", result)
}

func strToSlice(str string) ([]int, error) {
	result := make([]int, 0)
	for _, s := range str {
		// 判断字符串是否含有其他非数字字符
		tmp := int(s - '0')
		if tmp < 0 || tmp > 9 {
			return nil, errors.New("number is invalid!")
		}
		result = append(result, int(s-'0'))
	}
	return result, nil
}

// 完全包含
func arrFullContain(arr, tmp []int) bool {
	for _, v := range tmp {
		if !isContained(arr, v) {
			return false
		}
	}
	return true
}

// 完全包含1-9
func arrFullContainNine(arr []int) bool {
	tmp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range tmp {
		if !isContained(arr, v) {
			return false
		}
	}
	return true
}

// 不完全包含
func arrNotFullContain(arr, tmp []int) int {
	for _, v := range tmp {
		if !isContained(arr, v) {
			return v
		}
	}
	return 0
}

// 数组是否包含某个数字
func isContained(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func sliceToNumber(slice []int) int {
	sum := 0
	for i := len(slice) - 1; i >= 0; i-- {
		sum += int(math.Pow10(i))
	}
	return sum
}

// quickSort 快速排序
func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivotPos := partition(arr, left, right)
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
