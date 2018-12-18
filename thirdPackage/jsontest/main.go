package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("0000000000")
	time.Sleep(time.Second * 5)
	go func() {
		fmt.Println("1111111")
	}()
}

type resp struct {
	Age int `json:"age"`
	// Version int `json:"version"`
}
type resp2 struct {
	Age     *int    `json:"age"`
	Version *string `json:"version"`
}
