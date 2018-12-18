package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.gumpcome.com/common/tools/dbkit"
	"gitlab.gumpcome.com/common/tools/logkit"
)

func main() {
	dbkit.InitMysql("root", "1058471169", "127.0.0.1", "ncu", "ncu", 500, 800)

	gin.SetMode("debug")
	e := initEngine()
	e.Run(":2333")
}
func initEngine() *gin.Engine {
	r := gin.New()
	r.POST("/add", addInfo)

	return r
}

func addInfo(c *gin.Context) {
	ageStr := c.PostForm("age")
	age, _ := strconv.Atoi(ageStr)
	conn, err := dbkit.GetMysqlCon("ncu")
	if err != nil {
		logkit.Error(err)
	}
	conn.Exec("update tmp set age=age+? where id =1", age)
}
