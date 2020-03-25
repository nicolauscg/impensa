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
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{Handler: handler},
			),
		),
		beego.NSNamespace("/transaction",
			beego.NSBefore(controllers.AuthFilter),
			beego.NSInclude(
				&controllers.TransactionController{Handler: handler},
			),
		),
		beego.NSNamespace("/user",
			beego.NSBefore(controllers.AuthFilter),
			beego.NSInclude(
				&controllers.UserController{Handler: handler},
			),
		),
	)

	beego.AddNamespace(ns)
}
