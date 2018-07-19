package main

import (
	"fmt"
	"reflect"
)

// test

type handle = int
type MyStr string

func (m MyStr) Error() string {
	return string(m)
}
func test() error {
	var e *MyStr = nil

	return e
}
func main() {
	var var1 int = 1
	var var2 handle = 2
	types(var1)
	types(var2)

	var str *MyStr = nil
	fmt.Println(str, reflect.TypeOf(str))
	fmt.Println(test() == nil)
}

func types(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println(fmt.Sprintf("I am an int: %d", v))
	}
	switch v := val.(type) {
	case handle:
		fmt.Println(fmt.Sprintf("I am an handle: %d", v))
	}
}

// aaaaaaaaaaaa
