package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var wg sync.WaitGroup
	temp := make([]int, 4)
	i := 0
	i++
	go func(i int) {
		wg.Add(1)
		defer wg.Done()
		temp[i] = i
		fmt.Println(temp, 1)
	}(i)
	i++
	go func(i int) {
		wg.Add(1)
		defer wg.Done()
		temp[i] = i
		fmt.Println(temp, 2)
	}(i)
	i++
	go func(i int) {
		wg.Add(1)
		defer wg.Done()
		temp[i] = i
		fmt.Println(temp, 3)
	}(i)
	wg.Wait()
	fmt.Println(temp)
}
