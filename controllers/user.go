package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nicolauscg/impensa/constants"
	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	BaseController
}

// @Title get a user by id
// @router /:userId [get]
func (o *UserController) Get() {
	userId, err := primitive.ObjectIDFromHex(o.Ctx.Input.Param(":userId"))
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	user, err := o.Handler.Orms.User.GetOneById(userId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(user).ServeJSON()
}

// @Title update a user by id
// @router / [put]
func (o *UserController) Put() {
	var payload dt.UserUpdate
	json.Unmarshal(o.Ctx.Input.RequestBody, &payload)
	if payload.Id != o.UserId {
		o.ResponseBuilder.SetError(403, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()

		return
	}
	updateResult, err := o.Handler.Orms.User.UpdateOneById(payload.Id, &payload.Update)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(updateResult).ServeJSON()
}

// @Title delete a user by id
// @router / [delete]
func (o *UserController) Delete() {
	var payload dt.UserDelete
	json.Unmarshal(o.Ctx.Input.RequestBody, &payload)
	if payload.Id != o.UserId {
		o.ResponseBuilder.SetError(403, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()

		return
	}
	deleteResult, err := o.Handler.Orms.User.DeleteOneById(payload.Id)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(deleteResult).ServeJSON()
}
