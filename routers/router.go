package routers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	secrets "github.com/ijustfool/docker-secrets"
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

	var mgoConnString string
	if os.Getenv("APP_ENV") == "PROD" {
		beego.Info(fmt.Sprintf("load production environment"))
		if _, err := os.Stat(path.Join(projectDirPath, constants.EnvProdFileName)); err != nil {
			panic(err)
		}
		godotenv.Overload(constants.EnvProdFileName)
		dockerSecrets, err := secrets.NewDockerSecrets("")
		if err != nil {
			panic(err)
		}
		beego.BConfig.RunMode = "prod"
		mgoConnString, _ = dockerSecrets.Get("IMPENSA_BE_MGOCONNSTRING")
		apiSecret, _ := dockerSecrets.Get("IMPENSA_BE_API_SECRET")
		os.Setenv(constants.EnvApiSecret, apiSecret)
	} else {
		beego.Info(fmt.Sprintf("load development environment"))
		if _, err := os.Stat(path.Join(projectDirPath, constants.EnvDevLocalFileName)); err != nil {
			panic(err)
		}
		godotenv.Overload(constants.EnvDevLocalFileName)
		beego.BConfig.RunMode = "dev"
		mgoConnString = os.Getenv(constants.EnvMgoConnString)
	}
	dbName := os.Getenv(constants.EnvDBName)
	allowedOrigins := []string{os.Getenv(constants.EnvFrontendUrl)}

	handler, err := handlerPkg.NewHandler(dbName, mgoConnString)
	if err != nil {
		panic(err)
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
