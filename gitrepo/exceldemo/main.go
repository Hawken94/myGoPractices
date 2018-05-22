package main

import (
	"fmt"
	"log"
	"myGoPractices/utils"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	readExcel()
}

func readExcel() {
	excelFileName := "/Users/majing/workspace/Go/Development/src/myGoPractices/gitrepo/exceldemo/2018届毕业生7032.xlsx"
	xlsx, err := excelize.OpenFile(excelFileName)
	if err != nil {
		log.Fatalln("打开文件错误", err)
	}

	rows := xlsx.GetRows("学信网毕业生")
	count := 0

	for _, strs := range rows {
		if utils.StrsNotBlank(strs) {
			count++

			for _, v := range strs {
				fmt.Printf("\t %s", v)
			}
			fmt.Println()
			if count%2000 == 0 {
				fmt.Println("读取了2000行数据,准备插入到数据库中!")
				time.Sleep(10 * time.Second)
				continue
			}

		}
	}
}

type excelModel struct {
	ExamineeNo string `json:"examinee_no" desc:"考生号"`
	UnknowCode int    `json:"unknow_code" desc:"不知道的代号"`
}
