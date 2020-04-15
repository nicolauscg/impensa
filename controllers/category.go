package controllers

import (
	"net/http"

	"github.com/nicolauscg/impensa/constants"
	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryController struct {
	BaseController
}

// @Title create new category
// @Param newCategory  body  dt.CategoryInsert true  "newCategory"
// @router / [post]
func (o *CategoryController) CreateCategory(newCategory dt.CategoryInsert) {
	newCategory.User = o.UserId
	insertResult, err := o.Handler.Orms.Category.InsertOne(newCategory)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(insertResult).ServeJSON()
}

// @Title get a category by id
// @Param  id  path  string true "id"
// @router /:id [get]
func (o *CategoryController) GetCategory(id string) {
	categoryId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	category, err := o.Handler.Orms.Category.GetOneById(categoryId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()
	} else if category.User != o.UserId {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()
	} else {
		o.ResponseBuilder.SetData(category).ServeJSON()
	}
}

// @Title get all categories
// @router / [get]
func (o *CategoryController) GetAllCategories() {
	categories, err := o.Handler.Orms.Category.GetManyByUserId(o.UserId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(categories).ServeJSON()
}

// @Title update categories by ids
// @Param categoryUpdate  body  dt.CategoryUpdate true  "categoryUpdate"
// @router / [put]
func (o *CategoryController) UpdateCategories(categoryUpdate dt.CategoryUpdate) {
	userIds, err := o.Handler.Orms.Category.GetUserIdsByIds(categoryUpdate.Ids)
	if len(userIds) != 1 || userIds[0] != o.UserId {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()

		return
	}
	updateResult, err := o.Handler.Orms.Category.UpdateManyByIds(categoryUpdate.Ids, &categoryUpdate.Update)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(updateResult).ServeJSON()
}

// @Title delete categories by ids
// @Param categoryDelete  body  dt.CategoryDelete true  "categoryDelete"
// @router / [delete]
func (o *CategoryController) DeleteCategories(categoryDelete dt.CategoryDelete) {
	userIds, err := o.Handler.Orms.Category.GetUserIdsByIds(categoryDelete.Ids)
	if len(userIds) != 1 || userIds[0] != o.UserId {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()

		return
	}
	deleteResult, err := o.Handler.Orms.Category.DeleteManyByIds(categoryDelete.Ids)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(deleteResult).ServeJSON()
}
