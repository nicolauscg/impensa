package controllers

import (
	"net/http"

	"github.com/nicolauscg/impensa/constants"
	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionController struct {
	BaseController
}

// @Title create new transaction
// @Param newTransaction  body  dt.TransactionInsert true  "newTransaction"
// @router / [post]
func (o *TransactionController) CreateTransaction(newTransaction dt.TransactionInsert) {
	newTransaction.Owner = o.UserId
	insertResult, err := o.Handler.Orms.Transaction.InsertOne(newTransaction)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(insertResult).ServeJSON()
}

// @Title get a transaction by id
// @Param  id  path  string true "id"
// @router /:id [get]
func (o *TransactionController) GetTransaction(id string) {
	transactionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	transaction, err := o.Handler.Orms.Transaction.GetOneById(transactionId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()
	} else if transaction.Owner != o.UserId {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()
	} else {
		o.ResponseBuilder.SetData(transaction).ServeJSON()
	}
}

// @Title get all transactions
// @router / [get]
func (o *TransactionController) GetAllTransactions() {
	transactions, err := o.Handler.Orms.Transaction.GetManyByOwnerId(o.UserId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(transactions).ServeJSON()
}

// @Title update transactions by ids
// @Param transactionUpdate  body  dt.TransactionUpdate true  "transactionUpdate"
// @router / [put]
func (o *TransactionController) UpdateTransactions(transactionUpdate dt.TransactionUpdate) {
	ownerIds, err := o.Handler.Orms.Transaction.GetOwnerIdsByIds(transactionUpdate.Ids)
	if len(ownerIds) != 1 || ownerIds[0] != o.UserId {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()

		return
	}
	updateResult, err := o.Handler.Orms.Transaction.UpdateManyByIds(transactionUpdate.Ids, &transactionUpdate.Update)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(updateResult).ServeJSON()
}

// @Title delete transactions by ids
// @Param transactionDelete  body  dt.TransactionDelete true  "transactionDelete"
// @router / [delete]
func (o *TransactionController) DeleteTransactions(transactionDelete dt.TransactionDelete) {
	ownerIds, err := o.Handler.Orms.Transaction.GetOwnerIdsByIds(transactionDelete.Ids)
	if len(ownerIds) != 1 || ownerIds[0] != o.UserId {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()

		return
	}
	deleteResult, err := o.Handler.Orms.Transaction.DeleteManyByIds(transactionDelete.Ids)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(deleteResult).ServeJSON()
}
