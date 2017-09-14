package main

import (
	"fmt"
	"runtime"
)

// 使用两个 goroutine 交替打印序列，一个 goroutinue 打印数字， 另外一个goroutine打印字母，
// 最终效果如下 12AB34CD56EF78GH910IJ 。

// 思路： 问题很简单，使用 channel 来控制打印的进度。使用两个 channel ，来分别控制数字和字母的打印序列，
//  数字打印完成后通过 channel 通知字母打印, 字母打印完成后通知数字打印，然后周而复始的工作。

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	chanN := make(chan bool)
	chanC := make(chan bool, 1)
	done := make(chan struct{})

	go func() {
		for i := 1; i < 11; i += 2 {
			<-chanC
			fmt.Print(i)
			fmt.Print(i + 1)
			chanN <- true
		}
	}()

	go func() {
		charSeq := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
		for i := 0; i < 10; i += 2 {
			<-chanN
			fmt.Print(charSeq[i])
			fmt.Print(charSeq[i+1])
			chanC <- true
		}
		done <- struct{}{}
	}()

	chanC <- true
	<-done

}
