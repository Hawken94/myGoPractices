package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

// draw 随机设置字符单元的属性
func draw() {
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.Attribute(rand.Int()%8)+1)
		}
	}
	termbox.Flush() // 刷新后台缓存到界面中
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent() // 开始监听键盘事件
		}
	}()

	draw()
	for {
		select {
		case ev := <-eventQueue: // 获取键盘事件
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				return // 如果是 ESC 键,则退出程序
			}
		default:
			draw()
			time.Sleep(time.Millisecond * 1000)
		}
	}
}
