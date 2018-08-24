package tests

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	s := "2018-08-10 10:23:27"
	formatTime(s)
	date, err := time.Parse("2006-01-02 15:04:05", s)
	t.Errorf("date:%v err:%v", date, err)

}

func formatTime(str string) time.Time {

	t1 := make([]int, 0, 0)
	var t2 []int
	t3 := []int{}
	fmt.Println(t1 == nil, t2 == nil, t3 == nil)

	strs := "你好"
	for _, s := range strs {
		fmt.Println("-----", s)
	}

	if str == "" {
		str = "0000-01-01 00:00:00"
	}
	loc, _ := time.LoadLocation("Asia/Chongqing")

	date, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	if err != nil {
		log.Printf("时间转化错误:%v", err)
	}
	return date
}
