// Created by cgo - DO NOT EDIT

//line hello.go:1
package main

// void SayHello(_GoString_ s);
import _ "unsafe"

import "fmt"

func main() {
	(_Cfunc_SayHello)("Hello world \n")
}

//export SayHello
func SayHello(s string) {
	fmt.Print(s)
}
