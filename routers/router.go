package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/topic",
			beego.NSInclude(
				&controllers.TopicController{},
			),
		),

		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),

		beego.NSNamespace("/label",
			beego.NSInclude(
				&controllers.LabelController{},
			),
		),

		beego.NSNamespace("/admin",
			beego.NSInclude(
				&controllers.AdminController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
