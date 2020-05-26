package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/go-echarts/go-echarts/charts"
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
	categoryNameToAmountMap, err := o.getTransactionCategorySummaryData(dateTimeStart, dateTimeEnd)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	pieChartSliceInfos := make([]dt.PieChartSliceInfoWithoutId, 0)
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
	accountNameToAmountMap, err := o.getTransactionAccountSummaryData(dateTimeStart, dateTimeEnd)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	pieChartSliceInfos := make([]dt.PieChartSliceInfoWithoutId, 0)
	for accountName, amount := range accountNameToAmountMap {
		pieChartSliceInfos = append(pieChartSliceInfos, dt.PieChartSliceInfoWithoutId{accountName, amount})
	}

	o.ResponseBuilder.SetData(pieChartSliceInfos).ServeJSON()
}

// @Title get data for transaction graph by account
// @router /mail [get]
func (o *GraphController) MailTransactionSummary(
	dateTimeStart *time.Time,
	dateTimeEnd *time.Time,
	email *string,
) {
	dataCategory, err := o.getTransactionCategorySummaryData(dateTimeStart, dateTimeEnd)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	dataAccount, err := o.getTransactionAccountSummaryData(dateTimeStart, dateTimeEnd)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	mailMessage := o.Handler.Orms.MailGun.CreateMailMessage(dt.MailParam{
		Recipient: *email,
		Subject:   "Impensa transaction summary",
		Body:      "view attachments for summary of transactions",
	})

	for i, data := range []map[string]float32{dataCategory, dataAccount} {
		var dataSummaryType, fileName string
		if i == 0 {
			dataSummaryType = "Category"
			fileName = "summaryCategories"
		} else {
			dataSummaryType = "Account"
			fileName = "summaryAccounts"
		}

		pieData := make(map[string]interface{})
		for categoryName, amount := range data {
			pieData[categoryName] = amount
		}
		pie := charts.NewPie()
		pie.SetGlobalOptions(charts.TitleOpts{
			Title: fmt.Sprintf("Transaction per %v %v - %v", dataSummaryType, dateTimeStart, dateTimeEnd),
		})
		pie.Add("pie", pieData)
		buffer := bytes.NewBufferString("")
		pie.Render(buffer)
		mailMessage.AddBufferAttachment(fmt.Sprintf("%v.html", fileName), buffer.Bytes())
	}

	sendMailResponse, err := o.Handler.Orms.MailGun.SendMailMessage(mailMessage)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(sendMailResponse).ServeJSON()
}

func (o *GraphController) getTransactionCategorySummaryData(
	dateTimeStart *time.Time,
	dateTimeEnd *time.Time,
) (map[string]float32, error) {
	transactions, err := o.Handler.Orms.Transaction.GetManyNoObjectId(
		dt.TransactionQuery{User: &o.UserId, DateTimeStart: dateTimeStart, DateTimeEnd: dateTimeEnd, Limit: 0},
	)
	if err != nil {
		return nil, err
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
	for _, transaction := range transactionWithInferredRecurrences {
		if transaction.Category != nil {
			categoryNameToAmountMap[*transaction.Category] += *transaction.Amount
		}
	}

	return categoryNameToAmountMap, nil
}

func (o *GraphController) getTransactionAccountSummaryData(
	dateTimeStart *time.Time,
	dateTimeEnd *time.Time,
) (map[string]float32, error) {

	transactions, err := o.Handler.Orms.Transaction.GetManyNoObjectId(
		dt.TransactionQuery{User: &o.UserId, DateTimeStart: dateTimeStart, DateTimeEnd: dateTimeEnd, Limit: 0},
	)
	if err != nil {
		return nil, err
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
	for _, transaction := range transactionWithInferredRecurrences {
		if transaction.Category != nil {
			accountNameToAmountMap[*transaction.Account] += *transaction.Amount
		}
	}

	return accountNameToAmountMap, nil
}
