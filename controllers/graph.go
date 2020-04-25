package controllers

import (
	"net/http"

	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GraphController struct {
	BaseController
}

// @Title get data for transaction graph by category
// @router /transaction/category [get]
func (o *GraphController) GetTransactionCategorySummary() {
	transactions, err := o.Handler.Orms.Transaction.GetMany(dt.TransactionQuery{User: &o.UserId, Limit: 0})
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	categories, err := o.Handler.Orms.Category.GetManyByUserId(o.UserId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	categoryIdToAmountMap := make(map[primitive.ObjectID]float32)
	categoryIdToCategoryName := make(map[primitive.ObjectID]string)
	pieChartSliceInfos := make([]dt.PieChartSliceInfo, 0)

	for _, transaction := range transactions {
		if transaction.Category != nil {
			categoryIdToAmountMap[*transaction.Category] += *transaction.Amount //here
		}
	}
	for _, category := range categories {
		categoryIdToCategoryName[*category.Id] = *category.Name
	}

	for categoryId, amount := range categoryIdToAmountMap {
		pieChartSliceInfos = append(pieChartSliceInfos, dt.PieChartSliceInfo{&categoryId, categoryIdToCategoryName[categoryId], amount})
	}

	o.ResponseBuilder.SetData(pieChartSliceInfos).ServeJSON()
}

// @Title get data for transaction graph by account
// @router /transaction/account [get]
func (o *GraphController) GetTransactionAccountSummary() {
	transactions, err := o.Handler.Orms.Transaction.GetMany(dt.TransactionQuery{User: &o.UserId, Limit: 0})
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	accounts, err := o.Handler.Orms.Account.GetManyByUserId(o.UserId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	accountIdToAmountMap := make(map[primitive.ObjectID]float32)
	accountIdToAccountName := make(map[primitive.ObjectID]string)
	pieChartSliceInfos := make([]dt.PieChartSliceInfo, 0)

	for _, transaction := range transactions {
		if transaction.Account != nil {
			accountIdToAmountMap[*transaction.Account] += *transaction.Amount //here
		}
	}
	for _, account := range accounts {
		accountIdToAccountName[*account.Id] = *account.Name
	}

	for accountId, amount := range accountIdToAmountMap {
		pieChartSliceInfos = append(pieChartSliceInfos, dt.PieChartSliceInfo{&accountId, accountIdToAccountName[accountId], amount})
	}

	o.ResponseBuilder.SetData(pieChartSliceInfos).ServeJSON()
}
