package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(fib(1))

	st := time.Now().UnixNano()
	fib(1000000000)
	en := time.Now().UnixNano()
	fmt.Println(en - st)

	start := time.Now().UnixNano()
	fibByRecursion(300)
	end := time.Now().UnixNano()
	fmt.Println(end, start)
	fmt.Println(end - start)
}

// 自下而上的,非递归，厉害了
func fib(n int) int {
	f, g := 1, 0
	for n > 0 {
		g = g + f
		f = g - f
		n--
	}
	return g
}

// 普通的递归
func fibByRecursion(n int) int {
	if n <= 1 {
		return 1
	}
	return fibByRecursion(n-1) + fibByRecursion(n-2)
}
