package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		log.Println("打开文件失败")
		return
	}
	sheetMap := xlsx.GetSheetMap()
	for name, index := range sheetMap {
		fmt.Println(name, "  ", index)
	}

	rows := xlsx.GetRows(sheetMap[1])
	for _, strs := range rows {
		for _, v := range strs {
			fmt.Printf("\t %s", v)
		}
		fmt.Println()
	}
	fmt.Println("sldkjflsk")
}
