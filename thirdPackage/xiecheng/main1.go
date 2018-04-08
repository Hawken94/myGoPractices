package main

import (
	"fmt"
	"io"
)

func main() {
	input := 0
	nums := make([]int, 0)

	n, _ := fmt.Scan(&input)
	if n == 0 {
		return
	} else {
		// fmt.Printf("请输入%d个数字\n", input)
		// nums = make([]int, 0)
		for i := 0; i <= input; i++ {
			_, err := fmt.Scan(&i)
			if err == io.EOF {
				break
			} else {
				nums = append(nums, i)
			}
		}
	}
	fmt.Println(nums)

	temp := make([]int, 0)
	for _, num := range nums {
		if num != 0 {
			temp = append(temp, num)
		}
	}

	result := make([]int, 0)
	if len(nums) > len(temp) {
		for i := 0; i < len(nums)-len(temp); i++ {
			result = append(result, 0)
		}
		result = append(result, temp...)
	} else {
		result = nums
	}
	for _, v := range result {
		fmt.Println(v)
	}
}
