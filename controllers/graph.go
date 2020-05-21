package controllers

import (
	"net/http"
	"time"

	dt "github.com/nicolauscg/impensa/datatransfers"
)

type GraphController struct {
	BaseController
}

// @Title get data for transaction graph by category
// @Param dateTimeStart  query  time.Time false  "dateTimeStart"
// @Param dateTimeEnd  query  time.Time false  "dateTimeEnd"
// @router /transaction/category [get]
func (o *GraphController) GetTransactionCategorySummary(
	dateTimeStart *time.Time,
	dateTimeEnd *time.Time,
) {
	transactions, err := o.Handler.Orms.Transaction.GetManyNoObjectId(
		dt.TransactionQuery{User: &o.UserId, DateTimeStart: dateTimeStart, DateTimeEnd: dateTimeEnd, Limit: 0},
	)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}

	transactionWithInferredRecurrences := make([]*dt.TransactionNoObjectId, 0)
	for _, transaction := range transactions {
		if (transaction.IsReccurent == nil || !*transaction.IsReccurent) &&
			transaction.DateTime.After(*dateTimeStart) && transaction.DateTime.Before(*dateTimeEnd) {
			transactionWithInferredRecurrences = append(transactionWithInferredRecurrences, transaction)
		} else if transaction.IsReccurent != nil && *transaction.IsReccurent &&
			transaction.ReccurenceLastDate.After(*dateTimeStart) {
			incrementingDateTime := *transaction.DateTime
			for incrementingDateTime.Before(*dateTimeEnd) && incrementingDateTime.Before(*transaction.ReccurenceLastDate) {
				transactionCopy := transaction.CloneWithDifferentDateTime()
				temp := incrementingDateTime.Add(0)
				transactionCopy.DateTime = &temp
				transactionWithInferredRecurrences = append(transactionWithInferredRecurrences, transactionCopy)
				incrementingDateTime = transaction.RepeatInterval.GetTimeFrom(incrementingDateTime, 1)
			}
		}
	}

	categoryNameToAmountMap := make(map[string]float32)
	pieChartSliceInfos := make([]dt.PieChartSliceInfoWithoutId, 0)

	for _, transaction := range transactionWithInferredRecurrences {
		if transaction.Category != nil {
			categoryNameToAmountMap[*transaction.Category] += *transaction.Amount
		}
	}

	for categoryName, amount := range categoryNameToAmountMap {
		pieChartSliceInfos = append(pieChartSliceInfos, dt.PieChartSliceInfoWithoutId{categoryName, amount})
	}

	o.ResponseBuilder.SetData(pieChartSliceInfos).ServeJSON()
}

// @Title get data for transaction graph by account
// @router /transaction/account [get]
func (o *GraphController) GetTransactionAccountSummary(
	dateTimeStart *time.Time,
	dateTimeEnd *time.Time,
) {
	transactions, err := o.Handler.Orms.Transaction.GetManyNoObjectId(
		dt.TransactionQuery{User: &o.UserId, DateTimeStart: dateTimeStart, DateTimeEnd: dateTimeEnd, Limit: 0},
	)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}

	transactionWithInferredRecurrences := make([]*dt.TransactionNoObjectId, 0)
	for _, transaction := range transactions {
		if (transaction.IsReccurent == nil || !*transaction.IsReccurent) &&
			transaction.DateTime.After(*dateTimeStart) && transaction.DateTime.Before(*dateTimeEnd) {
			transactionWithInferredRecurrences = append(transactionWithInferredRecurrences, transaction)
		} else if transaction.IsReccurent != nil && *transaction.IsReccurent &&
			transaction.ReccurenceLastDate.After(*dateTimeStart) {
			incrementingDateTime := *transaction.DateTime
			for incrementingDateTime.Before(*dateTimeEnd) && incrementingDateTime.Before(*transaction.ReccurenceLastDate) {
				transactionCopy := transaction.CloneWithDifferentDateTime()
				temp := incrementingDateTime.Add(0)
				transactionCopy.DateTime = &temp
				transactionWithInferredRecurrences = append(transactionWithInferredRecurrences, transactionCopy)
				incrementingDateTime = transaction.RepeatInterval.GetTimeFrom(incrementingDateTime, 1)
			}
		}
	}

	accountNameToAmountMap := make(map[string]float32)
	pieChartSliceInfos := make([]dt.PieChartSliceInfoWithoutId, 0)

	for _, transaction := range transactionWithInferredRecurrences {
		if transaction.Category != nil {
			accountNameToAmountMap[*transaction.Account] += *transaction.Amount
		}
	}

	for accountName, amount := range accountNameToAmountMap {
		pieChartSliceInfos = append(pieChartSliceInfos, dt.PieChartSliceInfoWithoutId{accountName, amount})
	}

	o.ResponseBuilder.SetData(pieChartSliceInfos).ServeJSON()
}
