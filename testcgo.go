package main

// #include <hello.c>
// void sayHello(const char* s);
import (
	"C"
	"fmt"
)

func main() {
	C.sayHello(C.CString("hello, world\n"))
	reply := "reply"
	fmt.Println(reply)
}
