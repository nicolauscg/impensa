package controllers

import (
	"fmt"
	"net/http"
	"os"
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
// @Param limit  query  int false  "limit"
// @Param afterCursor  query  string false  "afterCursor"
// @router / [get]
func (o *TransactionController) GetAllTransactions(
	description *string,
	account *string,
	category *string,
	dateTimeStart *time.Time,
	dateTimeEnd *time.Time,
	amountMoreThan *float32,
	amountLessThan *float32,
	limit *int,
	afterCursor *string,
) {
	var accountObjectId, categoryObjectId, afterCursorObjectId *primitive.ObjectID = nil, nil, nil
	limitPlusOne := 20 + 1 // to check hasNext with cursor based pagination
	if account != nil {
		tmp, _ := primitive.ObjectIDFromHex(*account)
		accountObjectId = &tmp
	}
	if category != nil {
		tmp, _ := primitive.ObjectIDFromHex(*category)
		categoryObjectId = &tmp
	}
	if afterCursor != nil {
		tmp, _ := primitive.ObjectIDFromHex(*afterCursor)
		afterCursorObjectId = &tmp
	}
	if limit != nil {
		limitPlusOne = *limit + 1
	}
	transactions, err := o.Handler.Orms.Transaction.GetMany(
		dt.TransactionQuery{
			&o.UserId, accountObjectId, categoryObjectId,
			description, dateTimeStart, dateTimeEnd,
			amountMoreThan, amountLessThan,
			limitPlusOne,
			afterCursorObjectId,
		},
	)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	if transactions == nil || len(transactions) < 1 {
		o.ResponseBuilder.SetData([]*dt.Transaction{}).SetPaging(false, nil, nil).ServeJSON()

		return
	}

	var hasNext bool
	if len(transactions) < limitPlusOne {
		hasNext = false
	} else {
		hasNext = true
		// remove extra document as hasNext already determined
		transactions = transactions[:len(transactions)-1]
	}

	newQueryParams := o.Ctx.Request.URL.Query()
	lastIndex := len(transactions) - 1
	afterCursorInPaging := transactions[lastIndex].Id
	newQueryParams.Set("afterCursor", afterCursorInPaging.Hex())
	tempRequest, err := http.NewRequest(
		o.Ctx.Input.Method(),
		fmt.Sprintf("%v%v", os.Getenv(constants.EnvBackendUrl), o.Ctx.Input.URI()),
		nil,
	)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	tempRequest.URL.RawQuery = newQueryParams.Encode()
	nextUrlString := tempRequest.URL.String()
	newQueryParams.Set("next", nextUrlString)
	o.ResponseBuilder.SetData(transactions).SetPaging(hasNext, afterCursorInPaging, &nextUrlString).ServeJSON()
}

// @Title get description autocomplete
// @Param description  query sring false  "description"
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
