package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/nicolauscg/impensa/models"
)

type TransactionController struct {
	beego.Controller
}

// @router / [post]
func (o *TransactionController) Post() {
	var transaction models.Transaction
	json.Unmarshal(o.Ctx.Input.RequestBody, &transaction)
	orm, err := models.NewTransactionOrm()
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}
	insertResult, err := orm.InsertOne(transaction.Amount, transaction.Description)
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
	orm, err := models.NewTransactionOrm()
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}
	if transactionId != "" {
		transaction, err := orm.GetOneById(transactionId)
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
	orm, err := models.NewTransactionOrm()
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}
	transactions, err := orm.GetAll()
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

	orm, err := models.NewTransactionOrm()
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}

	updateResult, err := orm.UpdateOneById(transaction.Id, &transaction)
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
	orm, err := models.NewTransactionOrm()
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}
	deleteResult, err := orm.DeleteManyById(payload["transactionIds"])
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
	}
	o.Data["json"] = deleteResult
	o.ServeJSON()
}
