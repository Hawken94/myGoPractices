package main

import (
	"fmt"
)

func fibonacci2(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go fibonacci2(cap(ch), ch)
	for i := range ch {
		fmt.Println(i)
	}
}
