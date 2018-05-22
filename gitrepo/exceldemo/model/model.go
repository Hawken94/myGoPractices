package model

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"gitlab.gumpcome.com/common/go_kit/logiccode"
	"gitlab.gumpcome.com/common/go_kit/strkit"
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
	Education          int       `desc:"学历 0:本科"`
	LengthOfSchooling  int       `desc:"学年制"`
	UniversityMode     int       `desc:"培养形式 0:普通全日制 1:其他"`
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
	beego.Info(fmt.Sprintf("%s 数据库初始化成功...", cfgName))
}

// @Title 获取MySQL连接
func GetMysqlCon(cfgName string) (*sql.DB, error) {
	if cfgName == "" {
		return nil, logiccode.DbConfigNameErrorCode()
	}
	return dbs[cfgName], nil
}

// CreateDeleteMysqlSQL INSERT INTO `user`(name,age,email,gender,height,interests) VALUES (?,?,?,?,?,?)
func CreateMysqlInsertSQL(tableName string, data map[string]interface{}) (string, []interface{}) {
	dataLen := len(data)
	if dataLen <= 0 {
		return "", nil
	}

	params := make([]interface{}, 0)

	//构建INSERT部分的SQL格式
	insertStrBuilder := strkit.StringBuilder{}
	insertStrBuilder.Append("INSERT INTO `").Append(tableName).Append("`(")

	//构建VALUES部分的SQL格式
	valuesStrBuilder := strkit.StringBuilder{}
	valuesStrBuilder.Append(") VALUES (")

	for k, v := range data {
		if len(params) > 0 {
			insertStrBuilder.Append(", ")
			valuesStrBuilder.Append(", ")
		}
		insertStrBuilder.Append("`").Append(k).Append("`")
		valuesStrBuilder.Append("?")
		params = append(params, v)
	}
	valuesStrBuilder.Append(")")

	sql := strkit.StrJoin(insertStrBuilder.ToString(), valuesStrBuilder.ToString())
	return sql, params
}

// SaveInMysql 保存数据
// @Description 	返回的int64类型的值,只有在表主键定义为"auto increment"情况下,才会有效,其他情况默认返回0
// @param myDbCon	数据库连接
// @param tableName	表名称
// @param data		需要保存的K-V键值对,K:字段名,V:字段值
func SaveInMysql(myDbCon *sql.DB, tableName string, data map[string]interface{}) (bool, int64, error) {
	if myDbCon == nil {
		return false, 0, logiccode.DbConErrorCode()
	}
	if tableName == "" || data == nil {
		return false, 0, logiccode.DbInsertErrorCode()
	}
	sql, params := CreateMysqlInsertSQL(tableName, data)

	beego.Debug(fmt.Sprintf("SQL %s VALS %#v", sql, params))

	result, err := myDbCon.Exec(sql, params...)
	if err != nil {
		beego.Error(fmt.Sprintf("%v", err))
		return false, 0, logiccode.DbInsertErrorCode()
	}

	id, err := result.LastInsertId()

	return true, id, err
}
