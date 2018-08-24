package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testRange()
	// b1 := []byte{0x0A, 12}
	// fmt.Println(string(b1))
	// fmt.Println(convert(b1), len(convert(b1)))
	// var a int
	// fmt.Println(add(a))

	intSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(intSlice[1:], intSlice[3:])

	ip := map[string]bool{}
	ip["xhk"] = true
	delete(ip, "xhk")
	fmt.Println(ip)

	var big int32 = 0x12345678
	var small = int8(big)
	fmt.Printf("big=%#x,small=%#x \n", big, small)

	var test = func(i int) { fmt.Println(i) }

	test(12)

	fmt.Println(byte('Q'))
	matrix := make([][]int, 4)
	matrix[0][0] = 3
	fmt.Println(matrix)

}
func testRange() int {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		return i
	}
	return 0
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

func xhk(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	for i := 1; i < n; i++ {
		if i*i <= n && (i+1)*(i+1) > n {
			return i
		}
	}
	return 0
}
func mySqrt(x int) int {
	for i := 0; i <= x/2+1; i++ {
		if i*i > x {
			return i - 1
		} else if i*i == x {
			return i
		}
	}
	return 0
}
