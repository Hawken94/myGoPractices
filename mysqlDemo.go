package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"

	"gitlab.gumpcome.com/common/tools/strkit"

	"github.com/astaxie/beego"
	"gitlab.gumpcome.com/common/tools/dbkit"
)

func main() {
	mainDbUserName := "dev_gumpcome"
	mainDbUserPwd := "GumpDev2017"
	mainDbHost := "582eb0f46d463.bj.cdb.myqcloud.com:12637"
	mainDbName := "goods_mgr"
	mainDbCfgName := "goods_mgr"
	mainDbMaxIdle := 10
	mainDbMaxActive := 20
	dbkit.InitMysql(mainDbUserName, mainDbUserPwd, mainDbHost, mainDbName, mainDbCfgName, mainDbMaxIdle, mainDbMaxActive)

	save()
}

func save() {
	conn, err := dbkit.GetMysqlCon("goods_mgr")
	if err != nil {
		beego.Error(fmt.Sprintf("获取数据库连接失败:%v", err))
		return
	}
	sql := strkit.StringBuilder{}
	sql.Append("select * from goods_category_map_label_list where category_code IN (?)")

	data, err := dbkit.FindInMysql(conn, sql.ToString(), []string{""}, "'010100','100102'")
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(data)
}
