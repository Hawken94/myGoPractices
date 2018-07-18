package main

import (
	"flag"
	"fmt"
)

func main() {
	createTime := flag.String("d", "", "输入开始时间(格式为20160102)")
	flag.Parse()
	fmt.Println(*createTime)
}
