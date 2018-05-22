package main

import (
	"fmt"
	"log"
	"myGoPractices/gitrepo/exceldemo/model"
	"myGoPractices/utils"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const (
	mainDbUserName  = "root"
	mainDbUserPwd   = "1058471169"
	mainDbHost      = "127.0.0.1"
	mainDbName      = "ncu"
	mainDbCfgName   = "main"
	mainDbMaxIdle   = 10
	mainDbMaxActive = 20
)

var sheetName = "学信网毕业生"

func main() {
	model.InitMysql(mainDbUserName, mainDbUserPwd, mainDbHost, mainDbName, mainDbCfgName, mainDbMaxIdle, mainDbMaxActive)

	readExcel()
}

func readExcel() {
	excelFileName := "/Users/hawken/workspace/Go/Development/src/myGoPractices/gitrepo/exceldemo/2018届毕业生7032.xlsx"
	xlsx, err := excelize.OpenFile(excelFileName)
	if err != nil {
		log.Fatalln("打开文件错误", err)
	}

	rows := xlsx.GetRows("学信网毕业生")
	count := 0

	fmt.Println(xlsx.GetCellValue(sheetName, "A"+strconv.Itoa(12)))
	time.Sleep(10 * time.Second)

	var excelModels []*model.ExcelModel

	for i, strs := range rows {
		if utils.StrsNotBlank(strs) {
			count++

			excel := new(model.ExcelModel)

			for _, v := range strs {
				fmt.Printf("\t %s", v)
			}
			fmt.Println()
			if count%2000 == 0 {
				fmt.Println("读取了2000行数据,准备插入到数据库中!")
				time.Sleep(10 * time.Second)

				excelModels = make([]*model.ExcelModel, 0)
				continue
			}

		}
	}
}

func handleData(excel *model.ExcelModel, xlsx *excelize.File, i string) error {
	var err error
	excel.ExamineeNo = xlsx.GetCellValue(sheetName, "A"+i)
	excel.UnknowCode, err = strToInt(xlsx.GetCellValue(sheetName, "B"+i))
	checkErrWithReturn(err)
	excel.StudentNo = xlsx.GetCellValue(sheetName, "C"+i)
	excel.StuName = xlsx.GetCellValue(sheetName, "D"+i)

	gender := xlsx.GetCellValue(sheetName, "E"+i)
	excel.Gender = handleGender(gender)

	excel.Birthday = xlsx.GetCellValue(sheetName, "F"+i)
	excel.Identity = xlsx.GetCellValue(sheetName, "G"+i)
	excel.PoliticStatus = xlsx.GetCellValue(sheetName, "H"+i)
	excel.Nation = xlsx.GetCellValue(sheetName, "I"+i)
	excel.DegreeNo = xlsx.GetCellValue(sheetName, "J"+i)
	excel.School = xlsx.GetCellValue(sheetName, "K"+i)
	excel.MajorCode = xlsx.GetCellValue(sheetName, "L"+i)
	excel.MajorName = xlsx.GetCellValue(sheetName, "M"+i)
	excel.College = xlsx.GetCellValue(sheetName, "N"+i)
	excel.MajorDirection = xlsx.GetCellValue(sheetName, "O"+i)
	excel.ClassName = xlsx.GetCellValue(sheetName, "P"+i)
	excel.Education = xlsx.GetCellValue(sheetName, "Q"+i)
	excel.LengthOfSchooling = xlsx.GetCellValue(sheetName, "R"+i)
	excel.UniversityMode = xlsx.GetCellValue(sheetName, "S"+i)
	excel.SchoolBeginDate = xlsx.GetCellValue(sheetName, "T"+i)
	excel.SchoolYear = xlsx.GetCellValue(sheetName, "U"+i)
	excel.RegisterSchoolRoll = xlsx.GetCellValue(sheetName, "V"+i)
	excel.SchoolLeaveDate = xlsx.GetCellValue(sheetName, "W"+i)

	return nil
}
func strToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func checkErrWithReturn(err error) error {
	if err != nil {
		fmt.Println(err)
		return err
	}
}
func handleGender(g string) int {
	switch g {
	case "男":
		return 0
	case "女":
		return 1
	default:
		return 127
	}
}
