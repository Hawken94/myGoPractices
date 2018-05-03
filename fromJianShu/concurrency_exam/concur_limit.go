package main

import (
	"fmt"
	"time"
)

func main() {
	input := []int{3, 2, 1}
	timeout := 2
	chLimit := make(chan bool, 2)
	chs := make([]chan string, len(input))
	limitFunc := func(chLimit chan bool, ch chan string, taskID, sleepTime, timeout int) {
		Run(taskID, sleepTime, timeout, ch)
		<-chLimit
	}
	startTime := time.Now()
	fmt.Println("Mutilrun start")
	for i, sleepTime := range input {
		chs[i] = make(chan string, 1)
		chLimit <- true
		go limitFunc(chLimit, chs[i], i, sleepTime, timeout)
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Mutilrun finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
}

// Run 运行 goroutine
func Run(taskID, sleepTime, timeout int, ch chan string) {
	ch_run := make(chan string)
	go run(taskID, sleepTime, ch_run)
	select {
	case re := <-ch_run:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d,timeout", taskID)
		ch <- re
	}
}

func run(taskID, sleepTime int, ch chan string) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	ch <- fmt.Sprintf("task id %d, sleep %d second", taskID, sleepTime)
	return
}
