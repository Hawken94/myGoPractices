package chatbot

import (
	"fmt"
	"strings"
)

// simpleCN 代表对中文的演示机聊天机器人
type simpleCN struct {
	name string
	talk Talk
}

// NewSimpleCN 用于创建对中文的演示机聊天机器人
func NewSimpleCN(name string, talk Talk) Chatbot {
	return &simpleCN{name, talk}
}

func (robot *simpleCN) Name() string {
	return robot.name
}

func (robot *simpleCN) Begin() (string, error) {
	return "请输入你的名字:", nil
}

func (robot *simpleCN) Hello(username string) string {
	username = strings.TrimSpace(username)
	if robot.talk != nil {
		return robot.talk.Hello(username)
	}
	return fmt.Sprintf("你好,%s!我可以为你做些什么呢?", username)
}

func (robot *simpleCN) Talk(heard string) (saying string, end bool, err error) {
	heard = strings.TrimSpace(heard)
	if robot.talk != nil {
		return robot.talk.Talk(heard)
	}
	switch heard {
	case "":
		return
	case "没有", "再见":
		saying = "再见"
		end = true
		return
	default:
		saying = "对不起,我没听懂你说的."
		return
	}
}

func (robot *simpleCN) ReportError(err error) string {
	return fmt.Sprintf("发生了一个错误:%s \n", err)
}

func (robot *simpleCN) End() error {
	return nil
}
