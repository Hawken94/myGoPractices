package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"

	"baliance.com/gooxml/document"
)

func main() {
	result := make([]string, 0)
	doc, err := document.Open("./系统报错截图.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	for _, para := range doc.Paragraphs() {
		for _, run := range para.Runs() {
			result = append(result, run.Text())
		}
	}
	docStr := strings.Join(result, "")
	docStr = strings.TrimSpace(docStr)
	docStr = docStr[1:]
	docStr = docStr[:len(docStr)-1]
	strs := strings.Split(docStr, ",")

	// strMap := make(map[string]string)
	info := []*Info{}

	for _, str := range strs {
		str = strings.TrimSpace(str)

		strSlice := reg(str)
		// strMap[strSlice[0]] = strSlice[1]
		i := &Info{}
		i.Code = strSlice[0]
		i.City = strSlice[1]
		info = append(info, i)
	}
	fmt.Println(len(info))

	saveToExcel(info)

}

// Info 信息体
type Info struct {
	Code string
	City string
}

// 正则匹配字符串
// 根据客户号码4295270493所属地市592无权限操作！
func reg(str string) []string {
	pat := `\d+`
	reg := regexp.MustCompile(pat)
	result := reg.FindAllString(str, -1)
	return result
}

func saveToExcel(data []*Info) error {
	xlsx := excelize.NewFile()
	// 创建一个工作表
	sheetName := "Sheet1"
	index := xlsx.NewSheet(sheetName)
	// 设置工作簿的默认工作表
	xlsx.SetActiveSheet(index)

	xlsx.SetCellValue(sheetName, "A1", "客户号码")
	xlsx.SetCellValue(sheetName, "B1", "所属地市")

	// 插入数据
	count := 1
	for _, v := range data {
		xlsx.SetCellValue(sheetName, "A"+strconv.Itoa(count+1), v.Code)
		xlsx.SetCellValue(sheetName, "B"+strconv.Itoa(count+1), v.City)
		count++
	}

	err := xlsx.SaveAs("./系统报错截图.xlsx")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
