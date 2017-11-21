package main

import (
	"fmt"
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
<<<<<<< HEAD
	fmt.Println(isZiCard(0x78))
	v := []interface{}{nil, 0}
	fmt.Println(len(v))
	fmt.Println(string(append([]byte("hello"), " world"...)))

	slice := []int{1, 2, 3, 4, 5}
	sliceModify(slice)
	fmt.Println(slice)
}

func sliceModify(slice []int) []int {
	slice = append(slice, 6)
	fmt.Println(slice)
	return slice
=======
	// fmt.Println(isZiCard(0x78))
	// fmt.Println(chooseKingCard([]int{2, 3}))

	n := int32(3)
	if n%4 == 0 {
		n = n / 4
	} else {
		n = n/4 + 1
	}
	fmt.Println(n)

	s := []string([]string{"xhk", "123"})
	for _, v := range s {
		fmt.Println(v)
	}

	var test Test
	test.age = 12
	fmt.Println(test)
	zz := NewTest()
	fmt.Println(*zz)

	for i := 4; i > 0; i >>= 1 {
		fmt.Println(i)
	}
	fmt.Println(1 >> 1)

	qsort([]int{3, 1, 2, 4})
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
>>>>>>> 326c1e270ed2d8833d620094c5e0fe035a70f939
}
