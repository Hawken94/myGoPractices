package main

import (
	"fmt"
)

func main() {
	s := make([]int, 2, 3)
	s[0] = 1
	s[1] = 2
	s1 := s
	fmt.Println(s1, " ", s)
	s1 = append(s1, 3)
	s = append(s, 123)
	s[0] = 11
	fmt.Printf("s1 :%v %p s:%v %p", s1, &s1, s, &s)

}
