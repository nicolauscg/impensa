package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	dt "github.com/nicolauscg/impensa/datatransfers"
	handlerPkg "github.com/nicolauscg/impensa/handlers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	beego.Controller
	Handler *handlerPkg.Handler
}

// @Title get a user by id
// @router /:userId [get]
func (o *UserController) Get() {
	userId, err := primitive.ObjectIDFromHex(o.Ctx.Input.Param(":userId"))
	if err != nil {
		o.Data["json"] = dt.NewErrorResponse(500, err.Error())
		o.ServeJSON()
	}
	user, err := o.Handler.Orms.User.GetOneById(userId)
	if err != nil {
		o.Data["json"] = dt.NewErrorResponse(500, err.Error())
	} else {
		o.Data["json"] = dt.NewSuccessResponse(user)
	}
	o.ServeJSON()
}

// @Title update a user by id
// @router / [put]
func (o *UserController) Put() {
	var payload dt.UserUpdate
	json.Unmarshal(o.Ctx.Input.RequestBody, &payload)
	updateResult, err := o.Handler.Orms.User.UpdateOneById(payload.Id, &payload.Update)
	if err != nil {
		o.Data["json"] = dt.NewErrorResponse(500, err.Error())
	} else {
		o.Data["json"] = dt.NewSuccessResponse(updateResult)
	}
	o.ServeJSON()
}

// @Title delete a user by id
// @router / [delete]
func (o *UserController) Delete() {
	var payload dt.UserDelete
	json.Unmarshal(o.Ctx.Input.RequestBody, &payload)
	deleteResult, err := o.Handler.Orms.User.DeleteOneById(payload.Id)
	if err != nil {
		o.Data["json"] = dt.NewErrorResponse(500, err.Error())
		o.ServeJSON()
	}
	o.Data["json"] = dt.NewSuccessResponse(deleteResult)
	o.ServeJSON()
}
