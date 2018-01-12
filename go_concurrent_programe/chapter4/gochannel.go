package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		go func(no int) {
			fmt.Printf("num is %v \n", no)
		}(num)
		time.Sleep(time.Millisecond)
	}
}
