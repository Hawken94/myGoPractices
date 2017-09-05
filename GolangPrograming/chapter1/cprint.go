package main

// need gcc ,so failed.

/*
#include <stdio.h>
*/
import "C"
import "unsafe"

func main() {
	cstr := C.CString("hello hawken")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
