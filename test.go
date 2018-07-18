package main

import (
	"fmt"
	"time"
)

func isZiCard(ziCard int32) bool {
	ziCards := []int32{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	for _, card := range ziCards {
		if card == ziCard {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(false || true && false)
	nums := []int{1, 1, 3, 4}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			fmt.Println(i, nums[i])
		}
	}

	now := time.Now()
	fmt.Println(now.Unix() < now.AddDate(0, 1, -1).Unix())
}

func chooseKingCard(dicePoints []int) int {
	if len(dicePoints) == 2 {
		sum := dicePoints[0] + dicePoints[1]
		if sum > 1 && sum <= 4 {
			return 11
		} else if sum >= 5 && sum <= 8 {
			return 22
		} else if sum >= 9 && sum <= 12 {
			return 33
		}
	}
	return 0
}

type Test struct {
	name string
	age  int
}

func NewTest() *Test {
	return &Test{
		name: "123",
	}
}

var err error

func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid, i := data[0], 1
	head, tail := 0, len(data)-1
	for i = 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	data[head] = mid
	qsort(data[:head])
	qsort(data[head+1:])
}
