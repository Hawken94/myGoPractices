package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	i := 0
	c := cron.New()
	spec := "*/3 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		fmt.Printf("cron running:%v time:%v \n", i, time.Now().Format("2006-01-02 15:04:05"))
	})
	c.Start()
	select {}
}
