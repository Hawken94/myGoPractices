package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		go printGood()
		time.Sleep(2 * 100000000)
	}
}
func printGood() {
	fmt.Println("hawken 终于集齐五福了!!!")
}
