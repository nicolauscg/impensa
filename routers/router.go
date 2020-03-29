package routers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/joho/godotenv"
	"github.com/nicolauscg/impensa/constants"
	"github.com/nicolauscg/impensa/controllers"
	handlerPkg "github.com/nicolauscg/impensa/handlers"
)

func init() {
	projectDirPath, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	if _, errProd := os.Stat(path.Join(projectDirPath, constants.EnvProdFileName)); errProd == nil {
		godotenv.Load(constants.EnvProdFileName)
	} else if _, errDev := os.Stat(path.Join(projectDirPath, constants.EnvDevFileName)); os.IsNotExist(errProd) && errDev == nil {
		godotenv.Load(constants.EnvDevFileName)
	} else {
		panic("failed to read env file")
	}

	if os.Getenv(constants.EnvRunMode) == "prod" {
		beego.BConfig.RunMode = "prod"
	} else {
		beego.BConfig.RunMode = "dev"
	}
	beego.Info(fmt.Sprintf("Running beego in %v mode", beego.BConfig.RunMode))

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
