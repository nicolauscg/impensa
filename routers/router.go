package routers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/nicolauscg/impensa/controllers"
	handlerPkg "github.com/nicolauscg/impensa/handler"
)

func init() {
	handler, err := handlerPkg.NewHandler()
	if err != nil {
		panic(err)
	}
	fmt.Printf("handler in router %v\n", handler)

	ns := beego.NewNamespace("v1",
		beego.NSNamespace("/transaction",
			beego.NSInclude(
				&controllers.TransactionController{Handler: handler},
			),
		),
	)
	beego.AddNamespace(ns)
}
