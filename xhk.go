package main

import (
	"fmt"
)

func main() {
	str := "    "
	b := []byte(str)

	fmt.Println(len(b))
	fmt.Println(str == "")

}
