package logic

import (
	"strconv"
	"time"
)

// StrToInt 字符串转数字
func StrToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// HandleGender 处理性别
func HandleGender(g string) int {
	switch g {
	case "男":
		return 0
	case "女":
		return 1
	default:
		return 127
	}
}

// HandlePoliticStatus 处理政治面貌
func HandlePoliticStatus(p string) int {
	switch p {
	case "共青团员":
		return 0
	case "群众":
		return 1
	case "中共预备党员":
		return 2
	case "中共党员":
		return 3
	default:
		return 127
	}
}

// HandleEducation 处理学历
func HandleEducation(e string) int {
	switch e {
	case "本科":
		return 0
	default:
		return 1
	}
}

// HandleUniversityMode 处理培养方式
func HandleUniversityMode(u string) int {
	switch u {
	case "普通全日制":
		return 0
	default:
		return 1
	}
}

// StrToTime 字符串转化成时间格式
func StrToTime(str string) (time.Time, error) {
	dateStyle := "20060102"
	return time.Parse(dateStyle, str)
}
