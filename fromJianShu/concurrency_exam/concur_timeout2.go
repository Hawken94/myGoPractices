package main

import (
	"fmt"
	"time"
)

func main() {
	input := []int{3, 2, 1}
	chs := make([]chan string, len(input))
	startTime := time.Now()
	fmt.Println("Mutilrun start")
	for i, sleepTime := range input {
		chs[i] = make(chan string)
		go Run(i, sleepTime, chs[i])
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Mutilrun finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
}

// Run 运行 goroutine
func Run(taskID, sleepTime int, ch chan string) {
	ch_run := make(chan string)
	go run(taskID, sleepTime, ch_run)
	select {
	case re := <-ch_run:
		ch <- re
	case <-time.After(time.Duration(sleepTime) * time.Second):
		re := fmt.Sprintf("task id %d,timeout", taskID)
		ch <- re
	}
}

func run(taskID, sleepTime int, ch chan string) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	ch <- fmt.Sprintf("task id %d, sleep %d second", taskID, sleepTime)
	return
}
