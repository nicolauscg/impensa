package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AccountController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AccountController"],
        beego.ControllerComments{
            Method: "CreateAccount",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("newAccount", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AccountController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AccountController"],
        beego.ControllerComments{
            Method: "GetAllAccounts",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AccountController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AccountController"],
        beego.ControllerComments{
            Method: "UpdateAccounts",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("accountUpdate", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AccountController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AccountController"],
        beego.ControllerComments{
            Method: "DeleteAccounts",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("accountDelete", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AccountController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AccountController"],
        beego.ControllerComments{
            Method: "GetAccount",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"],
        beego.ControllerComments{
            Method: "OauthGoogleCallback",
            Router: `/google/callback`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("state", param.IsRequired),
				param.New("code", param.IsRequired),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"],
        beego.ControllerComments{
            Method: "OauthGoogleLogin",
            Router: `/google/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("credential", param.IsRequired, param.InBody),
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

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"],
        beego.ControllerComments{
            Method: "RequestResetUserPassword",
            Router: `/requestreset`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("requestResetUserPasswordBody", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"],
        beego.ControllerComments{
            Method: "ResetUserPassword",
            Router: `/resetpassword`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("resetUserPasswordBody", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:AuthController"],
        beego.ControllerComments{
            Method: "VerifyUser",
            Router: `/verify`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("userId", param.IsRequired),
				param.New("verifyKey", param.IsRequired),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "CreateCategory",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("newCategory", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetAllCategories",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "UpdateCategories",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("categoryUpdate", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "DeleteCategories",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("categoryDelete", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetCategory",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:GraphController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:GraphController"],
        beego.ControllerComments{
            Method: "MailTransactionSummary",
            Router: `/mail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("dateTimeStart"),
				param.New("dateTimeEnd"),
				param.New("email"),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:GraphController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:GraphController"],
        beego.ControllerComments{
            Method: "GetTransactionAccountSummary",
            Router: `/transaction/account`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("dateTimeStart"),
				param.New("dateTimeEnd"),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:GraphController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:GraphController"],
        beego.ControllerComments{
            Method: "GetTransactionCategorySummary",
            Router: `/transaction/category`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("dateTimeStart"),
				param.New("dateTimeEnd"),
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
            Method: "GetAllTransactions",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("description"),
				param.New("account"),
				param.New("category"),
				param.New("dateTimeStart"),
				param.New("dateTimeEnd"),
				param.New("amountMoreThan"),
				param.New("amountLessThan"),
				param.New("limit"),
				param.New("afterCursor"),
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

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "GetAccountsAndCategories",
            Router: `/create`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "GetSomeDescriptionAutocomplete",
            Router: `/description/complete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("description"),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "GetTransactionWithAccountsCategoriesRecurrence",
            Router: `/edit/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "ExportTransactions",
            Router: `/export`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("dateTimeStart"),
				param.New("dateTimeEnd"),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "ImportTransactions",
            Router: `/import`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("transactionImport", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"] = append(beego.GlobalControllerRouter["github.com/nicolauscg/impensa/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "GetAllTransactionsForTable",
            Router: `/table`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("description"),
				param.New("account"),
				param.New("category"),
				param.New("dateTimeStart"),
				param.New("dateTimeEnd"),
				param.New("amountMoreThan"),
				param.New("amountLessThan"),
				param.New("limit"),
				param.New("afterCursor"),
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
