package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"gitlab.gumpcome.com/common/go_kit/timekit"
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
	/*
		fmt.Println(false || true && false)
		nums := []int{1, 1, 3, 4}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] {
				fmt.Println(i, nums[i])
			}
		}
	*/

	slice := make([]int, 10, 10)
	slice = append(slice, 1)
	fmt.Println(slice)

	yesterday := time.Now().Add(-1e9 * 60 * 60 * 24)
	fmt.Println(yesterday)

	fmt.Println(1e9)

	startDate, err := timekit.StringToTime(strconv.Itoa(20160310), "20060102")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(startDate)

	fmt.Println("20170704"[6:8])

	var buf bytes.Buffer
	buf.WriteString("20180222"[:6])
	buf.WriteString("01")
	fmt.Println(buf.String())

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
