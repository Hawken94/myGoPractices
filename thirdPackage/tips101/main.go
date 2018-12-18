package main

import (
	"fmt"
	"myGoPractices/thirdPackage/tips101/foo"
	"unsafe"
)

func main() {
	// conf := foo.Config{1, "haha", 123}         // doesn't compile
	conf := foo.Config{Name: "bar", Size: 123} // compile okay
	// spew.Dump(conf)
	fmt.Println(conf)
	// fmt.Println(runtime.Version())
	fmt.Println("size:", unsafe.Sizeof(conf))
	foo.P()
}
