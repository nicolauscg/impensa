package controllers

import (
	"net/http"
	"time"

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
	newTransaction.User = &o.UserId
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
	} else if *transaction.User != o.UserId {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()
	} else {
		o.ResponseBuilder.SetData(transaction).ServeJSON()
	}
}

// @Title get all transactions
// @Param description  query  string false  "description"
// @Param account  query  string false  "account"
// @Param category  query string false  "category"
// @Param dateTimeStart  query  time.Time false  "dateTimeStart"
// @Param dateTimeEnd  query  time.Time false  "dateTimeEnd"
// @router / [get]
func (o *TransactionController) GetAllTransactions(
	description *string,
	account *string,
	category *string,
	dateTimeStart *time.Time,
	dateTimeEnd *time.Time,
	amountMoreThan *float32,
	amountLessThan *float32,
) {
	var accountObjectId, categoryObjectId *primitive.ObjectID = nil, nil
	if account != nil {
		tmp, _ := primitive.ObjectIDFromHex(*account)
		accountObjectId = &tmp
	}
	if category != nil {
		tmp, _ := primitive.ObjectIDFromHex(*category)
		categoryObjectId = &tmp
	}
	transactions, err := o.Handler.Orms.Transaction.GetMany(
		dt.TransactionQuery{&o.UserId, accountObjectId, categoryObjectId, description, dateTimeStart, dateTimeEnd, amountMoreThan, amountLessThan},
	)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	if transactions == nil {
		transactions = []*dt.Transaction{}
	}
	o.ResponseBuilder.SetData(transactions).ServeJSON()
}

// @Title get description autocomplete
// @Param description  query  string false  "description"
// @router /description/complete [get]
func (o *TransactionController) GetSomeDescriptionAutocomplete(description *string) {
	if description == nil || len(*description) < 3 {
		o.ResponseBuilder.SetData([]string{}).ServeJSON()

		return
	}
	suggestions := []string{}
	descriptionObjectList, err := o.Handler.Orms.Transaction.GetSomeDescriptionsByPartialDescription(
		&dt.TransactionDescriptionAutocomplete{&o.UserId, description, 5},
	)
	for _, descriptionObject := range descriptionObjectList {
		suggestions = append(suggestions, *descriptionObject.Id)
	}
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(suggestions).ServeJSON()
}

// @Title update transactions by ids
// @Param transactionUpdate  body  dt.TransactionUpdate true  "transactionUpdate"
// @router / [put]
func (o *TransactionController) UpdateTransactions(transactionUpdate dt.TransactionUpdate) {
	userIds, err := o.Handler.Orms.Transaction.GetUserIdsByIds(transactionUpdate.Ids)
	if len(userIds) != 1 || userIds[0] != o.UserId {
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
	userIds, err := o.Handler.Orms.Transaction.GetUserIdsByIds(transactionDelete.Ids)
	if len(userIds) != 1 || userIds[0] != o.UserId {
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
