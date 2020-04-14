package controllers

import (
	"net/http"

	"github.com/nicolauscg/impensa/constants"
	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountController struct {
	BaseController
}

// @Title create new account
// @Param newAccount  body  dt.AccountInsert true  "newAccount"
// @router / [post]
func (o *AccountController) CreateAccount(newAccount dt.AccountInsert) {
	newAccount.User = o.UserId
	insertResult, err := o.Handler.Orms.Account.InsertOne(newAccount)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(insertResult).ServeJSON()
}

// @Title get an account by id
// @Param  id  path  string true "id"
// @router /:id [get]
func (o *AccountController) GetAccount(id string) {
	accountId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	account, err := o.Handler.Orms.Account.GetOneById(accountId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()
	} else if account.User != o.UserId {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()
	} else {
		o.ResponseBuilder.SetData(account).ServeJSON()
	}
}

// @Title get all accounts
// @router / [get]
func (o *AccountController) GetAllAccounts() {
	accounts, err := o.Handler.Orms.Account.GetManyByUserId(o.UserId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(accounts).ServeJSON()
}

// @Title update accounts by ids
// @Param accountUpdate  body  dt.AccountUpdate true  "accountUpdate"
// @router / [put]
func (o *AccountController) UpdateAccounts(accountUpdate dt.AccountUpdate) {
	userIds, err := o.Handler.Orms.Account.GetUserIdsByIds(accountUpdate.Ids)
	if len(userIds) != 1 || userIds[0] != o.UserId {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()

		return
	}
	updateResult, err := o.Handler.Orms.Account.UpdateManyByIds(accountUpdate.Ids, &accountUpdate.Update)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(updateResult).ServeJSON()
}

// @Title delete accounts by ids
// @Param accountDelete  body  dt.AccountDelete true  "accountDelete"
// @router / [delete]
func (o *AccountController) DeleteAccounts(accountDelete dt.AccountDelete) {
	userIds, err := o.Handler.Orms.Account.GetUserIdsByIds(accountDelete.Ids)
	if len(userIds) != 1 || userIds[0] != o.UserId {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()

		return
	}
	deleteResult, err := o.Handler.Orms.Account.DeleteManyByIds(accountDelete.Ids)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(deleteResult).ServeJSON()
}
