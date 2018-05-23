package model

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 数据源
var dbs = make(map[string]*sql.DB)

// ExcelModel 要导入 excel model
type ExcelModel struct {
	ExamineeNo         string    `json:"examinee_no" desc:"考生号"`
	UnknowCode         int       `json:"unknow_code" desc:"不知道的代号"`
	StudentNo          string    `desc:"学号"`
	StuName            string    `desc:"姓名"`
	Gender             int       `desc:"性别 0:男 1:女"`
	Birthday           time.Time `desc:"出生年月日"`
	Identity           string    `desc:"身份证"`
	PoliticStatus      int       `desc:"政治面貌 0:共青团员 1:群众 2:中共预备党员 3:中共党员"`
	Nation             string    `desc:"民族"`
	DegreeNo           int       `desc:"学位证编号"`
	School             string    `desc:"学校名称"`
	MajorCode          int       `desc:"专业代码"`
	MajorName          string    `desc:"专业名称"`
	College            string    `desc:"学院名称"`
	MajorDirection     string    `desc:"专业方向"`
	ClassName          string    `desc:"班级名称"`
	Education          int       `desc:"学历 0:本科 1:其他"`
	LengthOfSchooling  int       `desc:"学年制"`
	UniversityMode     int       `desc:"培养方式 0:普通全日制 1:其他"`
	SchoolBeginDate    time.Time `desc:"入学时间"`
	SchoolYear         int       `desc:"入学年份"`
	RegisterSchoolRoll string    `desc:"注册学籍"`
	SchoolLeaveDate    time.Time `desc:"离校时间"`
}

func addDbCfg(cfgName string, db *sql.DB) error {
	if cfgName == "" {
		return errors.New("mysql config name is nil")
	}

	if db == nil {
		return errors.New("mysql conn is nil")
	}

	if dbs[cfgName] != nil {
		return errors.New("config name is existed")
	}

	dbs[cfgName] = db
	return nil
}

// InitMysql 初始化数据库
// @param userName 	用户名
// @param userPwd 	密码
// @param host 		地址
// @param dbName 	数据库名称
// @param cfgName	数据库连接池配置名称
// @param maxIdle 	最大活跃连接数
// @param maxActive	最大连接数
func InitMysql(userName string, userPwd string, host string, dbName string, cfgName string, maxIdle int, maxActive int) {
	if userName == "" || userPwd == "" || host == "" || dbName == "" {
		panic(fmt.Sprintf("%s mysql connection info is empty!", cfgName))
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local", userName, userPwd, host, dbName))

	if err != nil {
		panic(err.Error())
	}
	if maxIdle <= 0 {
		maxIdle = 10
	}
	if maxActive <= 0 {
		maxActive = 20
	}
	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxActive)
	//设置连接的存活时间,不是指sleep时间,而是从创建连接时算起。
	db.SetConnMaxLifetime(3 * time.Minute)
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	addDbCfg(cfgName, db)
	fmt.Printf("%s 数据库初始化成功...", cfgName)
}

// GetMysqlCon @Title 获取MySQL连接
func GetMysqlCon(cfgName string) (*sql.DB, error) {
	if cfgName == "" {
		return nil, errors.New("连接池名称为空")
	}
	return dbs[cfgName], nil
}

// InfoAdd 信息添加
func InfoAdd(model []*ExcelModel) error {
	db, err := GetMysqlCon("main")
	if err != nil {
		fmt.Println("获取数据库连接失败:", err)
		return err
	}

	sqlExec := StringBuilder{}
	sqlExec.Append("INSERT INTO student_info")
	sqlExec.Append(" (`examinee_no`,`unknown_code`,`student_no`,`stu_name`,`gender`,")
	sqlExec.Append("`birthday`,`indentity`,`politic_status`,`nation`,`degree_no`,")
	sqlExec.Append("`school`,`major_code`,`major_name`,`college`,`major_direction`,")
	sqlExec.Append("`class_name`,`length_of_schooling`,`university_mode`,`school_begin_date`,`school_year`,")
	sqlExec.Append("`register_school_roll`,`school_leave_date`) VALUES ")

	params := make([]interface{}, 0, 10*len(model))
	n := 1
	m := len(model)

	for _, v := range model {
		sqlExec.Append("(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)") // 22个参数
		if n < m {
			sqlExec.Append(",")
		}

		params = append(params, v.ExamineeNo, v.UnknowCode, v.StudentNo, v.StuName, v.Gender,
			v.Birthday, v.Identity, v.PoliticStatus, v.Nation, v.DegreeNo,
			v.School, v.MajorCode, v.MajorName, v.College, v.MajorDirection,
			v.ClassName, v.LengthOfSchooling, v.UniversityMode, v.SchoolBeginDate, v.SchoolYear,
			v.RegisterSchoolRoll, v.SchoolLeaveDate)
		n++
	}
	result, err := db.Exec(sqlExec.ToString(), params...)
	if err != nil {
		fmt.Println("插入信息失败:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("获取受影响的行数失败:", err)
		return err
	}
	fmt.Println("受影响的行数为:", rowsAffected)

	return nil
}

// StringBuilder 实例构造方法
type StringBuilder struct {
	buf bytes.Buffer
}

// Append 添加字符串到字符串构建实例里面
//	实例构造方法: sb.Append("hello").Append(" world")
func (sb *StringBuilder) Append(str string) *StringBuilder {
	if str != "" {
		sb.buf.WriteString(str)
	}
	return sb
}

// ToString 输出字符串构建实例里面的所有字符串
//	实例构造方法: sb.Append("hello").Append(" world").ToString()
func (sb *StringBuilder) ToString() string {
	return sb.buf.String()
}
