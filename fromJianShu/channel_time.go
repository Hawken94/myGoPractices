package main

import (
	"fmt"
	"time"
)

func timenow(ch chan string) {
	tn := time.Now().Format("2006年01月02日 15点04分05秒.0000000 时区 -0700")
	ch <- tn
}

func main() {
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go timenow(ch)
		fmt.Println(<-ch)
		time.Sleep(1000 * time.Millisecond)
	}
}
