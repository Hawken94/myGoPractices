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

	ip := map[string]bool{}
	ip["xhk"] = true
	delete(ip, "xhk")
	fmt.Println(ip)

	var big int32 = 0x12345678
	var small = int8(big)
	fmt.Printf("big=%#x,small=%#x \n", big, small)
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
