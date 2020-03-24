package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	dt "github.com/nicolauscg/impensa/datatransfers"
	handlerPkg "github.com/nicolauscg/impensa/handlers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionController struct {
	beego.Controller
	Handler *handlerPkg.Handler
}

// @router / [post]
func (o *TransactionController) Post() {
	var transaction dt.TransactionInsert
	json.Unmarshal(o.Ctx.Input.RequestBody, &transaction)
	insertResult, err := o.Handler.Orms.Transaction.InsertOne(transaction)
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}
	o.Data["json"] = insertResult
	o.ServeJSON()
}

// @router /:transactionId [get]
func (o *TransactionController) Get() {
	transactionId, err := primitive.ObjectIDFromHex(o.Ctx.Input.Param(":transactionId"))
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}
	transaction, err := o.Handler.Orms.Transaction.GetOneById(transactionId)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = transaction
	}
	o.ServeJSON()
}

// @router / [get]
func (o *TransactionController) GetAll() {
	transactions, err := o.Handler.Orms.Transaction.GetAll()
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}
	o.Data["json"] = transactions
	o.ServeJSON()
}

// @router / [put]
func (o *TransactionController) Put() {
	var payload dt.TransactionUpdate
	json.Unmarshal(o.Ctx.Input.RequestBody, &payload)
	updateResult, err := o.Handler.Orms.Transaction.UpdateManyByIds(payload.Ids, &payload.Update)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = updateResult
	}
	o.ServeJSON()
}

// @router / [delete]
func (o *TransactionController) Delete() {
	var payload dt.TransactionDelete
	json.Unmarshal(o.Ctx.Input.RequestBody, &payload)
	deleteResult, err := o.Handler.Orms.Transaction.DeleteManyByIds(payload.Ids)
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}
	o.Data["json"] = deleteResult
	o.ServeJSON()
}
