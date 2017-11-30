package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	b1 := []byte{0x0A, 12}
	fmt.Println(string(b1))
	fmt.Println(convert(b1), len(convert(b1)))
	var a int
	fmt.Println(add(a))

	intSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(intSlice[1:], intSlice[3:])

	x := []int{1, 2, 3, 4, 5, 6}
}

func convert(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, "")
}

func add(a int) int {
	a = 6
	return a
}
