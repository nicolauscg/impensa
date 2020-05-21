package controllers

import (
	"fmt"
	"net/http"
	"os"
	"sort"
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
	if *newTransaction.IsReccurent && (newTransaction.RepeatCount == nil || newTransaction.RepeatInterval == nil) {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, constants.ErrorTransactionRecurrenceOptionInvalid).ServeJSON()

		return
	}
	if *newTransaction.IsReccurent {
		temp := newTransaction.RepeatInterval.GetTimeFrom(*newTransaction.DateTime, *newTransaction.RepeatCount)
		newTransaction.ReccurenceLastDate = &temp
	}
	insertResult, err := o.Handler.Orms.Transaction.InsertOne(newTransaction)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(insertResult).ServeJSON()
}

// @Title get a transaction by id with accounts and categories
// @Param  id  path  string true "id"
// @router /edit/:id [get]
func (o *TransactionController) GetTransactionWithAccountsCategoriesRecurrence(id string) {
	transactionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	transaction, err := o.Handler.Orms.Transaction.GetOneById(transactionId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	} else if *transaction.User != o.UserId {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorResourceForbiddenOrNotFound).ServeJSON()

		return
	}
	accounts, err := o.Handler.Orms.Account.GetManyByUserId(o.UserId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	categories, err := o.Handler.Orms.Category.GetManyByUserId(o.UserId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}

	o.ResponseBuilder.SetData(map[string]interface{}{
		"transaction": transaction,
		"accounts":    accounts,
		"categories":  categories,
	}).ServeJSON()
}

// @Title get accounts and categories
// @router /create [get]
func (o *TransactionController) GetAccountsAndCategories() {
	accounts, err := o.Handler.Orms.Account.GetManyByUserId(o.UserId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	categories, err := o.Handler.Orms.Category.GetManyByUserId(o.UserId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}

	o.ResponseBuilder.SetData(map[string]interface{}{
		"accounts":   accounts,
		"categories": categories,
	}).ServeJSON()
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
// @router /table [get]
func (o *TransactionController) GetAllTransactionsForTable(
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
	transactions, err := o.Handler.Orms.Transaction.GetManyNoObjectId(
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

	var hasNext bool
	var reccurencePagingDateTimeStart, reccurencePagingDateTimeEnd time.Time
	var afterCursorInPaging *primitive.ObjectID
	var nextUrlString *string

	if afterCursor != nil {
		for _, transaction := range transactions {
			if transaction.Id.Hex() == *afterCursor {
				reccurencePagingDateTimeStart = *transaction.DateTime
				break
			}
		}
	} else {
		reccurencePagingDateTimeStart = *dateTimeStart
	}

	if transactions == nil || len(transactions) < 1 {
		hasNext = false
		reccurencePagingDateTimeEnd = *dateTimeEnd
	} else if len(transactions) < limitPlusOne {
		hasNext = false
		reccurencePagingDateTimeEnd = *dateTimeEnd
	} else {
		hasNext = true
		reccurencePagingDateTimeEnd = *transactions[len(transactions)-1].DateTime
		// remove extra document as hasNext already determined
		transactions = transactions[:len(transactions)-1]
	}

	if hasNext && (transactions != nil || len(transactions) > 0) {
		newQueryParams := o.Ctx.Request.URL.Query()
		lastIndex := len(transactions) - 1
		afterCursorInPaging = transactions[lastIndex].Id
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
		temp := tempRequest.URL.String()
		nextUrlString = &temp
		newQueryParams.Set("next", *nextUrlString)
	} else {
		afterCursorInPaging = nil
		nextUrlString = nil
	}

	inferredTransactionsFromReccurences := make([]*dt.TransactionNoObjectId, 0)
	var incrementingDateTime time.Time
	var transactionCopy *dt.TransactionNoObjectId
	for _, transaction := range transactions {
		if transaction.IsReccurent != nil && *transaction.IsReccurent {
			if err != nil {
				o.ResponseBuilder.SetError(http.StatusInternalServerError, constants.ErrorCloningReccurentTransaction).ServeJSON()

				return
			}
			incrementingDateTime = *transaction.DateTime
			for incrementingDateTime.Before(*transaction.ReccurenceLastDate) &&
				incrementingDateTime.Before(*dateTimeEnd) &&
				incrementingDateTime.Before(reccurencePagingDateTimeEnd) {
				transactionCopy = transaction.CloneWithDifferentDateTime()
				if incrementingDateTime.After(*transaction.DateTime) &&
					incrementingDateTime.After(*dateTimeStart) &&
					incrementingDateTime.After(reccurencePagingDateTimeStart) {
					temp := incrementingDateTime.Add(0)
					transactionCopy.DateTime = &temp
					inferredTransactionsFromReccurences = append(inferredTransactionsFromReccurences, transactionCopy)
				}
				incrementingDateTime = transaction.RepeatInterval.GetTimeFrom(incrementingDateTime, 1)
			}
		}
	}
	result := append(transactions, inferredTransactionsFromReccurences...)
	sort.SliceStable(result, func(i, j int) bool {
		return (*result[i].DateTime).Before(*result[j].DateTime)
	})
	filteredResult := make([]*dt.TransactionNoObjectId, 0)
	for _, transaction := range result {
		if transaction.DateTime.After(*dateTimeStart) && transaction.DateTime.Before(*dateTimeEnd) {
			filteredResult = append(filteredResult, transaction)
		}
	}
	o.ResponseBuilder.SetData(filteredResult).SetPaging(hasNext, afterCursorInPaging, nextUrlString).ServeJSON()
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
	if transactionUpdate.Update.IsReccurent != nil &&
		transactionUpdate.Update.RepeatCount != nil &&
		transactionUpdate.Update.RepeatInterval != nil {
		temp := transactionUpdate.Update.RepeatInterval.GetTimeFrom(
			*transactionUpdate.Update.DateTime, *transactionUpdate.Update.RepeatCount,
		)
		transactionUpdate.Update.ReccurenceLastDate = &temp
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
