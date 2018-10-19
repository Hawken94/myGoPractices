package main

// void SayHello(_GoString_ s);
import "C"

import "fmt"

func main() {
	C.SayHello("Hello world \n")
	reply := "hello"
}

//export SayHello
func SayHello(s string) {
	fmt.Print(s)
}
