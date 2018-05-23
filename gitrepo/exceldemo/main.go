package main

import (
	"fmt"
	"log"
	"myGoPractices/gitrepo/exceldemo/logic"
	"myGoPractices/gitrepo/exceldemo/model"
	"myGoPractices/utils"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const (
	mainDbUserName  = "root"
	mainDbUserPwd   = "1058471169"
	mainDbHost      = "127.0.0.1:3306"
	mainDbName      = "ncu"
	mainDbCfgName   = "main"
	mainDbMaxIdle   = 10
	mainDbMaxActive = 20
)

var sheetName = "学信网毕业生"

func main() {
	begin := time.Now()
	model.InitMysql(mainDbUserName, mainDbUserPwd, mainDbHost, mainDbName, mainDbCfgName, mainDbMaxIdle, mainDbMaxActive)

	if err := readExcel(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("程序运行时间:", time.Since(begin))
}

func readExcel() error {
	excelFileName := "./2018届毕业生7032.xlsx"
	xlsx, err := excelize.OpenFile(excelFileName)
	if err != nil {
		log.Fatalln("打开文件错误", err)
	}

	rows := xlsx.GetRows("学信网毕业生")
	count := 0
	canTimes := 0 // 处理剩余的
	rowsLen := len(rows)
	// totalTimes := len(rows) / 2000

	fmt.Println(xlsx.GetCellValue(sheetName, "A"+strconv.Itoa(12)))
	// time.Sleep(1 * time.Second)

	var excelModels []*model.ExcelModel

	for i, strs := range rows {
		if utils.StrsNotBlank(strs) {
			count++
			excel := new(model.ExcelModel)
			handleData(excel, xlsx, strconv.Itoa(i+1))
			excelModels = append(excelModels, excel)

			if count%2000 == 0 {
				canTimes++
				fmt.Println("读取了2000行数据,准备插入到数据库中!")
				tmp := time.Now()
				model.InfoAdd(excelModels)
				fmt.Println("插入耗时:", time.Since(tmp))

				excelModels = make([]*model.ExcelModel, 0)
				continue
			}
			// 将剩余的插入到数据库中
			if count == rowsLen {
				fmt.Println("读取剩余数据,准备插入到数据库中!")
				tmp := time.Now()
				model.InfoAdd(excelModels)
				fmt.Println("插入耗时:", time.Since(tmp))
			}

		}
	}
	fmt.Println(count, "----->", len(rows))
	return nil
}

func handleData(excel *model.ExcelModel, xlsx *excelize.File, i string) error {
	var err error
	excel.ExamineeNo = xlsx.GetCellValue(sheetName, "A"+i)
	excel.UnknowCode, err = logic.StrToInt(xlsx.GetCellValue(sheetName, "B"+i))
	if err != nil {
		fmt.Println(err)
		return err
	}

	excel.StudentNo = xlsx.GetCellValue(sheetName, "C"+i)
	excel.StuName = xlsx.GetCellValue(sheetName, "D"+i)

	gender := xlsx.GetCellValue(sheetName, "E"+i)
	excel.Gender = logic.HandleGender(gender)

	excel.Birthday, err = logic.StrToTime(xlsx.GetCellValue(sheetName, "F"+i))
	if err != nil {
		fmt.Println(err)
		return err
	}

	excel.Identity = xlsx.GetCellValue(sheetName, "G"+i)
	pStatus := xlsx.GetCellValue(sheetName, "H"+i)
	excel.PoliticStatus = logic.HandlePoliticStatus(pStatus)

	excel.Nation = xlsx.GetCellValue(sheetName, "I"+i)
	excel.DegreeNo, err = logic.StrToInt(xlsx.GetCellValue(sheetName, "J"+i))
	if err != nil {
		fmt.Println(err)
		return err
	}

	excel.School = xlsx.GetCellValue(sheetName, "K"+i)
	excel.MajorCode, err = logic.StrToInt(xlsx.GetCellValue(sheetName, "L"+i))
	if err != nil {
		fmt.Println(err)
		return err
	}

	excel.MajorName = xlsx.GetCellValue(sheetName, "M"+i)
	excel.College = xlsx.GetCellValue(sheetName, "N"+i)
	excel.MajorDirection = xlsx.GetCellValue(sheetName, "O"+i)
	excel.ClassName = xlsx.GetCellValue(sheetName, "P"+i)
	excel.Education = logic.HandleEducation(xlsx.GetCellValue(sheetName, "Q"+i))
	excel.LengthOfSchooling, err = logic.StrToInt(xlsx.GetCellValue(sheetName, "R"+i))
	if err != nil {
		fmt.Println(err)
		return err
	}

	excel.UniversityMode = logic.HandleUniversityMode(xlsx.GetCellValue(sheetName, "S"+i))
	excel.SchoolBeginDate, err = logic.StrToTime(xlsx.GetCellValue(sheetName, "T"+i))
	if err != nil {
		fmt.Println(err)
		return err
	}

	excel.SchoolYear, err = logic.StrToInt(xlsx.GetCellValue(sheetName, "U"+i))
	excel.RegisterSchoolRoll = xlsx.GetCellValue(sheetName, "V"+i)
	excel.SchoolLeaveDate, err = logic.StrToTime(xlsx.GetCellValue(sheetName, "W"+i))

	return nil
}
