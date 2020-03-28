package controllers

import (
	"net/http"

	"github.com/nicolauscg/impensa/constants"
	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	BaseController
}

// @Title get a user by id
// @Param   id    path    string     true  "id"
// @router /:id [get]
func (o *UserController) GetUser(id string) {
	userId, err := primitive.ObjectIDFromHex(id)
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
// @Param userUpdate  body  dt.UserDelete true  "userUpdate"
// @router / [put]
func (o *UserController) UpdateUser(userUpdate dt.UserUpdate) {
	if userUpdate.Id != o.UserId {
		o.ResponseBuilder.SetError(403, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()

		return
	}
	updateResult, err := o.Handler.Orms.User.UpdateOneById(userUpdate.Id, &userUpdate.Update)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(updateResult).ServeJSON()
}

// @Title delete a user by id
// @Param userDelete  body  dt.UserDelete true  "userDelete"
// @router / [delete]
func (o *UserController) DeleteUser(userDelete dt.UserDelete) {
	if userDelete.Id != o.UserId {
		o.ResponseBuilder.SetError(403, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()

		return
	}
	deleteResult, err := o.Handler.Orms.User.DeleteOneById(userDelete.Id)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(deleteResult).ServeJSON()
}
