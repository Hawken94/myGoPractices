package main

import (
	"fmt"
	"math"
)

func main() {
	n := float64(1)
	var left, right float64
	for {
		left = 100 * n * n
		right = math.Pow(2, n)
		// x的最小值是 15
		if left <= right {
			fmt.Println(n)
			break
		}
		n++
	}
}
