package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:transactionId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:userId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
