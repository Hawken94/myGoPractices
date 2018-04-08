package main

import (
	"fmt"
)

func main1() {
	s1 := []int{1, 2, 3}
	fmt.Printf("s1 len:%v cap:%v \n", len(s1), cap(s1))
	s2 := s1[:0]
	fmt.Printf("s2 len:%v cap:%v \n", len(s2), cap(s2))
	fmt.Println(s1, s2)
	fmt.Printf("s1:%p s2:%p \n", &s1, &s2)
	s2 = append(s2, 4)
	fmt.Printf("s2 len:%v cap:%v \n", len(s2), cap(s2))

	fmt.Println(s1)
	fmt.Println(s2)
}

func main() {
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	fmt.Println(x)
	y := append(s, 12)
	fmt.Println(s, x, y)
}
