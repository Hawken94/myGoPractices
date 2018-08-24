package tests

import (
	"fmt"
	"testing"
)

var c = count()

func TestVar(t *testing.T) {
	s := c
	test()
	r := c
	t.Error(s, r)
}
func test() {
	i := c
	fmt.Println(i)
}

func count() int {
	fmt.Println("000000")
	return 1
}
