package tests

import (
	"regexp"
	"testing"
)

func TestReg(t *testing.T) {
	t.Error(reg())
}

func reg() []string {
	str := "根据客户号码4295270493所属地市592无权限操作！"
	pat := `\d+`
	reg := regexp.MustCompile(pat)
	result := reg.FindAllString(str, -1)
	return result
}
