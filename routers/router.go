package routers

import (
	"github.com/astaxie/beego"
	"github.com/nicolauscg/impensa/controllers"
	handlerPkg "github.com/nicolauscg/impensa/handlers"
)

func init() {
	handler, err := handlerPkg.NewHandler("test")
	if err != nil {
		panic(err)
	}

	ns := beego.NewNamespace("v1",
		beego.NSNamespace("/transaction",
			beego.NSInclude(
				&controllers.TransactionController{Handler: handler},
			),
		),
	)
	beego.AddNamespace(ns)
}
