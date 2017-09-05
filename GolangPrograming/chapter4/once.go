package main

import (
	"fmt"
	"sync"
)

var a string
var once sync.Once

func setup() {
	a = "hello world"
}

func doprint() {
	once.Do(setup)
	fmt.Print(a)
}

func main() {
	go doprint()
	go doprint()
}
