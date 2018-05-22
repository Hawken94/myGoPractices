package controllers

import (
	"fmt"
	"myGoPractices/testPractices/beegoTest/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// HandleGet ...
// @router /test [get]
func (c *MainController) HandleGet() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	c.TplName = "index.tpl"
}

// HandlePost ...
// @router /test [post]
func (c *MainController) HandlePost() {
	params := new(models.QueryParams)

	if err := c.ParseForm(params); err != nil {
		beego.Error("参数解析错误", fmt.Sprintf("%#v", err))
		c.Ctx.Output.JSON("参数解析错误", true, false)
		return
	}
	/*
		 if err := json.Unmarshal(c.Ctx.Input.RequestBody, params); err != nil {
			beego.Error("参数解析错误", fmt.Sprintf("%#v", err))
			c.Ctx.Output.JSON("参数解析错误", true, false)
			return
		} */
	name := c.GetString("name")

	fmt.Printf("params :%v name:%s \n", params, name)
	c.Ctx.Output.JSON(params, true, false)
}
