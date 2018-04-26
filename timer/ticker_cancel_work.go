package main

import (
	"fmt"
	"time"
)

// 能定时的去做某件事，并且在执行时间超时的时候，能把这个定时器关掉

func main() {
	res := make(chan interface{}, 1000)
	timeout := time.After(10 * time.Second)
	DoTickerWork2(res, timeout)
	for v := range res {
		fmt.Println(v)
	}
}

// DoTickerWork 定时器任务
func DoTickerWork(res chan interface{}, timeout <-chan time.Time) {
	t := time.NewTicker(3 * time.Second)
	defer t.Stop()

	go func() {
		defer close(res)
		i := 1
		for {
			<-t.C
			fmt.Printf("start %v th work \n", i)
			res <- i
			i++
		}
	}()
	<-timeout
	return
}

// DoTickerWork2 定时器任务 解决死锁
func DoTickerWork2(res chan interface{}, timeout <-chan time.Time) {
	t := time.NewTicker(3 * time.Second)
	done := make(chan bool, 1)
	defer t.Stop()

	go func() {
		defer close(res)
		i := 1
		for {
			select {
			case <-t.C:
				fmt.Printf("start %v th work \n", i)
				res <- i
				i++
			case <-timeout:
				close(done)
				return
			}
		}
	}()
	<-done // TODO:有什么作用?
	return
}
