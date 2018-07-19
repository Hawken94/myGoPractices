package main

import (
	"flag"
	"fmt"

	"gitlab.gumpcome.com/common/go_kit/timekit"
)

func main() {
	createTime := flag.String("d", "", "输入开始时间(格式为20160102)")
	flag.Parse()
	fmt.Println(*createTime)

	_, err := timekit.StringToTime("", "20060102")
	if err != nil {
		fmt.Println(err)
		return
	}
}
