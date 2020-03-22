package routers

import (
	"github.com/astaxie/beego"
	"github.com/nicolauscg/impensa/controllers"
)

func init() {
	ns := beego.NewNamespace("v1",
		beego.NSNamespace("/transaction",
			beego.NSInclude(
				&controllers.TransactionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
