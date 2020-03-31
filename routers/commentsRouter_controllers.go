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
            MethodParams: param.Make(
				param.New("credential", param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("newUser", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "CreateTransaction",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("newTransaction", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "GetAllTransactions",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "UpdateTransactions",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("transactionUpdate", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "DeleteTransactions",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("transactionDelete", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "GetTransaction",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("userUpdate", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("userDelete", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

}
