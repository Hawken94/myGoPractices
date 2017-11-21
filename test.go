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
}
