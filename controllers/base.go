package controllers

import (
	"github.com/astaxie/beego"
	dt "github.com/nicolauscg/impensa/datatransfers"
	handlerPkg "github.com/nicolauscg/impensa/handlers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseController struct {
	beego.Controller
	UserId          primitive.ObjectID
	ResponseBuilder dt.ResponseBuilder
	Handler         *handlerPkg.Handler
}

func (a *BaseController) Prepare() {
	if userId := a.Ctx.Input.Param("userId"); len(userId) > 0 {
		a.UserId, _ = primitive.ObjectIDFromHex(userId)
	}
	a.ResponseBuilder = dt.NewResponseBuilder(a.Ctx.ResponseWriter)
}
