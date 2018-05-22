package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["myGoPractices/testPractices/beegoTest/controllers:MainController"] = append(beego.GlobalControllerRouter["myGoPractices/testPractices/beegoTest/controllers:MainController"],
		beego.ControllerComments{
			Method: "HandleGet",
			Router: `/test`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myGoPractices/testPractices/beegoTest/controllers:MainController"] = append(beego.GlobalControllerRouter["myGoPractices/testPractices/beegoTest/controllers:MainController"],
		beego.ControllerComments{
			Method: "HandlePost",
			Router: `/test`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
