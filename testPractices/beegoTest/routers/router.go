package routers

import (
	"myGoPractices/testPractices/beegoTest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.MainController{})
}
