package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/nicolauscg/impensa/controllers"
	handlerPkg "github.com/nicolauscg/impensa/handlers"
)

func init() {
	handler, err := handlerPkg.NewHandler("test")
	if err != nil {
		panic(err)
	}

	allowedOrigins := make([]string, 0)
	if beego.BConfig.RunMode == "dev" {
		allowedOrigins = append(allowedOrigins, "http://localhost:3000")
	}

	ns := beego.NewNamespace("v1",
		beego.NSBefore(cors.Allow(&cors.Options{
			AllowOrigins:     allowedOrigins,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
			AllowCredentials: true,
		})),
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
