package controllers

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
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
	if userUpdate.Update.Picture != nil {
		resizedBase64Img, err := resizeBase64PngImage(*userUpdate.Update.Picture, 180)
		if err != nil {
			o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

			return
		}
		userUpdate.Update.Picture = &resizedBase64Img
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

func resizeBase64PngImage(originalBase64Png string, dimension int) (resizedBase64Png string, err error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(originalBase64Png))
	originalImage, _, err := image.Decode(reader)
	if err != nil {
		return
	}
	resizedImage := imaging.Resize(originalImage, dimension, dimension, imaging.NearestNeighbor)
	var resizedImgBuffer bytes.Buffer
	png.Encode(&resizedImgBuffer, resizedImage)
	resizedBase64Png = base64.StdEncoding.EncodeToString(resizedImgBuffer.Bytes())

	return
}
