package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	handlerPkg "github.com/nicolauscg/impensa/handler"
	"github.com/nicolauscg/impensa/models"
)

type TransactionController struct {
	beego.Controller
	Handler *handlerPkg.Handler
}

// @router / [post]
func (o *TransactionController) Post() {
	var transaction models.Transaction
	json.Unmarshal(o.Ctx.Input.RequestBody, &transaction)
	insertResult, err := o.Handler.TransactionOrm.InsertOne(transaction.Amount, transaction.Description)
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}
	o.Data["json"] = insertResult
	o.ServeJSON()
}

// @router /:transactionId [get]
func (o *TransactionController) Get() {
	transactionId := o.Ctx.Input.Param(":transactionId")
	if transactionId != "" {
		transaction, err := o.Handler.TransactionOrm.GetOneById(transactionId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = transaction
		}
	}
	o.ServeJSON()
}

// @router / [get]
func (o *TransactionController) GetAll() {
	transactions, err := o.Handler.TransactionOrm.GetAll()
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}
	o.Data["json"] = transactions
	o.ServeJSON()
}

// @router / [put]
func (o *TransactionController) Put() {
	var transaction models.Transaction
	json.Unmarshal(o.Ctx.Input.RequestBody, &transaction)
	updateResult, err := o.Handler.TransactionOrm.UpdateOneById(transaction.Id, &transaction)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = updateResult
	}
	o.ServeJSON()
}

// @router / [delete]
func (o *TransactionController) Delete() {
	var payload map[string][]string
	json.Unmarshal(o.Ctx.Input.RequestBody, &payload)
	deleteResult, err := o.Handler.TransactionOrm.DeleteManyById(payload["transactionIds"])
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}
	o.Data["json"] = deleteResult
	o.ServeJSON()
}
