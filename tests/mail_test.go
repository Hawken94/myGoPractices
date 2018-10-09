package tests

import (
	"fmt"
	"net/smtp"
	"strings"
	"testing"

	"gitlab.gumpcome.com/common/tools/idkit"
)

func TestEmail(t *testing.T) {
	// if err := email(); err != nil {
	// 	t.Errorf("发送邮件失败:%v", err)
	// }
	t.Error(idkit.CreateMd5("xionghekuan"))
}

// auth := smtp.PlainAuth("", "1058471169@qq.com", "yhaqivmvlhjebfgh", "smtp.qq.com")
func email() error {
	auth := smtp.PlainAuth("", "1058471169@qq.com", "yhaqivmvlhjebfgh", "smtp.qq.com")

	sendContent := sendContent{}
	sendContent.Sender = "1058471169@qq.com"
	sendContent.ToWhos = []string{"1058471169@qq.com"}
	sendContent.Nickname = "无处安放的青春"
	sendContent.Subject = "约吗"
	sendContent.Body = "小哥哥,小哥哥,你好帅啊!"

	contentType := "Content-Type: text/plain; charset=UTF-8"

	sendContent.Msg = []byte(fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n%s",
		strings.Join(sendContent.ToWhos, ","), sendContent.Nickname, sendContent.Sender, sendContent.Subject,
		contentType, sendContent.Body))
	// sendContent.Msg = []byte("To:" + strings.Join(sendContent.ToWhos, ",") + "\r\nFrom:" + sendContent.Nickname +
	// 	"<" + sendContent.Sender + ">\r\nSubject:" + sendContent.Subject + "\r\n" + contentType + "\r\n\r\n" +
	// 	sendContent.Body)

	fmt.Println(string(sendContent.Msg))
	err := smtp.SendMail("smtp.qq.com:25", auth, sendContent.Sender, sendContent.ToWhos, sendContent.Msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
	return err
}

type sendContent struct {
	Sender   string
	ToWhos   []string
	Nickname string
	Subject  string
	Body     string
	Msg      []byte
}
