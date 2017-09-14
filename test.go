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
}
