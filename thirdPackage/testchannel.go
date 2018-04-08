package main

import (
	"fmt"
)

func main() {
	/* for i := 0; i < 20; i++ {
		go printGood()
		time.Sleep(2 * 100000000)
	} */
	mySlice := make([]int, 100)
	copy(mySlice[3:5], []int{3, 4})
	fmt.Println(mySlice[3:5])

}
func printGood() {
	fmt.Println("hawken 终于集齐五福了!!!")
}
