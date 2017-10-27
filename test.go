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
	fmt.Println(isZiCard(0x78))
	fmt.Println(chooseKingCard([]int{2, 3}))

	n := int32(3)
	if n%4 == 0 {
		n = n / 4
	} else {
		n = n/4 + 1
	}
	fmt.Println(n)

	var test Test
	test.age = 12
	fmt.Println(test)
	zz := NewTest()
	fmt.Println(*zz)
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
