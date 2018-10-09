package main

import (
	"fmt"
	"sync"
	"time"
)

func dostuff(wg *sync.WaitGroup, i int) {
	fmt.Printf("goroutine id first: %d\n", i)
	time.Sleep(3 * time.Second)
	fmt.Printf("goroutine id end: %d\n", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	workers := 10

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go dostuff(&wg, i)
	}
	wg.Wait()
}
