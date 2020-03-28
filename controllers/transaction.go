package controllers

import (
	"encoding/json"

	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionController struct {
	BaseController
}

// @Title create new transaction
// @router / [post]
func (o *TransactionController) Post() {
	var transaction dt.TransactionInsert
	json.Unmarshal(o.Ctx.Input.RequestBody, &transaction)
	transaction.Owner = o.UserId
	insertResult, err := o.Handler.Orms.Transaction.InsertOne(transaction)
	if err != nil {
		o.ResponseBuilder.SetError(500, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(insertResult).ServeJSON()
}

// @Title get a transaction by id
// @router /:transactionId [get]
func (o *TransactionController) Get() {
	transactionId, err := primitive.ObjectIDFromHex(o.Ctx.Input.Param(":transactionId"))
	if err != nil {
		o.ResponseBuilder.SetError(500, err.Error()).ServeJSON()

		return
	}
	transaction, err := o.Handler.Orms.Transaction.GetOneById(transactionId)
	if err != nil {
		o.ResponseBuilder.SetError(500, err.Error()).ServeJSON()
	} else if transaction.Owner != o.UserId {
		o.ResponseBuilder.SetError(403, "not authorized to access").ServeJSON()
	} else {
		o.ResponseBuilder.SetData(transaction).ServeJSON()
	}
}

// @Title get all transactions
// @router / [get]
func (o *TransactionController) GetAll() {
	transactions, err := o.Handler.Orms.Transaction.GetManyByOwnerId(o.UserId)
	if err != nil {
		o.ResponseBuilder.SetError(500, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(transactions).ServeJSON()
}

// @Title update transactions by ids
// @router / [put]
func (o *TransactionController) Put() {
	var payload dt.TransactionUpdate
	json.Unmarshal(o.Ctx.Input.RequestBody, &payload)
	ownerIds, err := o.Handler.Orms.Transaction.GetOwnerIdsByIds(payload.Ids)
	if len(ownerIds) != 1 || ownerIds[0] != o.UserId {
		o.ResponseBuilder.SetError(403, "not authorized to access or missing resource").ServeJSON()

		return
	}
	updateResult, err := o.Handler.Orms.Transaction.UpdateManyByIds(payload.Ids, &payload.Update)
	if err != nil {
		o.ResponseBuilder.SetError(500, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(updateResult).ServeJSON()
}

// @Title delete transactions by ids
// @router / [delete]
func (o *TransactionController) Delete() {
	var payload dt.TransactionDelete
	json.Unmarshal(o.Ctx.Input.RequestBody, &payload)
	ownerIds, err := o.Handler.Orms.Transaction.GetOwnerIdsByIds(payload.Ids)
	if len(ownerIds) != 1 || ownerIds[0] != o.UserId {
		o.ResponseBuilder.SetError(403, "not authorized to access or missing resource").ServeJSON()

		return
	}
	deleteResult, err := o.Handler.Orms.Transaction.DeleteManyByIds(payload.Ids)
	if err != nil {
		o.ResponseBuilder.SetError(500, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(deleteResult).ServeJSON()
}
