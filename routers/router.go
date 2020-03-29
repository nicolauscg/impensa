package routers

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/nicolauscg/impensa/constants"
	"github.com/nicolauscg/impensa/controllers"
	handlerPkg "github.com/nicolauscg/impensa/handlers"
)

func init() {
	if runmode := os.Getenv(constants.EnvRunMode); runmode == "prod" {
		beego.BConfig.RunMode = "prod"
	} else {
		beego.BConfig.RunMode = "dev"
	}

	beego.Info(fmt.Sprintf("router init: runmode %v", beego.BConfig.RunMode))

	handler, err := handlerPkg.NewHandler(os.Getenv(constants.EnvDBName), os.Getenv(constants.EnvMgoConnString))
	if err != nil {
		panic(err)
	}

	allowedOrigins := make([]string, 0)
	if beego.BConfig.RunMode == "dev" {
		allowedOrigins = append(allowedOrigins, os.Getenv(constants.EnvFrontendUrl))
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
				&controllers.AuthController{controllers.BaseController{Handler: handler}},
			),
		),
		beego.NSNamespace("/transaction",
			beego.NSBefore(controllers.AuthFilter),
			beego.NSInclude(
				&controllers.TransactionController{controllers.BaseController{Handler: handler}},
			),
		),
		beego.NSNamespace("/user",
			beego.NSBefore(controllers.AuthFilter),
			beego.NSInclude(
				&controllers.UserController{controllers.BaseController{Handler: handler}},
			),
		),
	)

	beego.AddNamespace(ns)
}
