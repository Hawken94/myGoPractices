package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	loopTimer()

	wg.Wait()
}

func loopTimer() {
	defer wg.Done()
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()

	for {
		<-timer.C
		fmt.Println(time.Now().Format("2006-01-02_15:04:05"))

	}

}
