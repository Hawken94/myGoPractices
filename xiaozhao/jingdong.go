package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(MagicNum(1000, 1111))
}

func MagicNum(low, hight int) int {
	tempStr, count := "", 0
	var tempInt []int

	for i := low; i <= hight; i++ {
		if i < 10 {
			continue
		}
		tempStr = strconv.Itoa(i)
		for k := range tempStr {
			tempInt = append(tempInt, int(tempStr[k]-'0'))
		}

		if solute(tempInt) {
			fmt.Println(tempStr)
			count++
		}
		tempInt = make([]int, 0)
	}
	return count
}

func solute(tempInt []int) bool {
	// 排序
	QuickSort(tempInt, 0, len(tempInt)-1)
	max := tempInt[len(tempInt)-1]
	sum := 0
	for _, v := range tempInt {
		sum += v
	}
	// 通过平均值来判断是否是神奇数
	if sum%2 != 0 {
		return false
	}
	average := sum / 2
	if average < max {
		return false
	}

	if average == max {
		return true
	} else if average > max {
		for i := 0; i < len(tempInt)-2; i++ {
			if (max + tempInt[i]) == average {
				return true
			} else if (max + tempInt[i]) < average {
				max = (max + tempInt[i])
			}
		}
	}

	return false
}

// QuickSort 快速排序算法
func QuickSort(values []int, left, right int) {

	temp := values[left]

	p := left

	i, j := left, right

	for i <= j {

		for j >= p && values[j] >= temp {
			j--
		}

		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= temp && i <= p {
			i++
		}

		if i <= p {
			values[p] = values[i]
			p = i
		}
	}

	values[p] = temp

	if p-left > 1 {
		QuickSort(values, left, p-1)
	}

	if right-p > 1 {
		QuickSort(values, p+1, right)
	}
}
