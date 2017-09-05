package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(500 * time.Millisecond)
	boom := time.After(2000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("bomm!")
			return
		default:
			fmt.Println("	.")
			// time.Sleep(250 * time.Millisecond)
		}
	}
}
