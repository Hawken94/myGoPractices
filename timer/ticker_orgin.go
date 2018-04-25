package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 使用 time 包实现定时器
func main() {

	wg.Add(2)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go loopTimer(interrupt)
	go forceTimer()

	wg.Wait()
}

func loopTimer(interrupt chan os.Signal) {
	defer wg.Done()
	ticker := time.NewTicker(3 * time.Second)
	tick := time.Tick(10 * time.Second)
	boom := time.After(9 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now().Format("2006-01-02_15:04:05"))
		case <-boom:
			fmt.Println("到期了")
			return
		case <-tick:
			fmt.Println("结束了")
			return
		case <-interrupt:
			fmt.Println("系统强制退出")
			return
		default:
			fmt.Println("xxxxxxx")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func forceTimer() {
	defer wg.Done()
	for {
		time.Sleep(3 * time.Second)
		fmt.Println(time.Now().Format("2006-01-02_15:04:05"))
	}
}
