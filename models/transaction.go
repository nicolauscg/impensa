package models

import (
	"context"
	"fmt"

	"github.com/nicolauscg/impensa/constants"
	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransactionOrmer interface {
	InsertOne(insert dt.TransactionInsert) (*mongo.InsertOneResult, error)
	InsertMany(inserts []interface{}) (*mongo.InsertManyResult, error)
	GetMany(query dt.TransactionQuery) ([]*dt.Transaction, error)
	GetManyNoObjectId(query dt.TransactionQuery) ([]*dt.TransactionNoObjectId, error)
	GetSomeDescriptionsByPartialDescription(partialDescription *dt.TransactionDescriptionAutocomplete) ([]*dt.TransactionDescriptionAutocompleteResponse, error)
	GetUserIdsByIds(ids []primitive.ObjectID) ([]primitive.ObjectID, error)
	GetOneById(id primitive.ObjectID) (*dt.Transaction, error)
	UpdateManyByIds(ids []primitive.ObjectID, update *dt.TransactionUpdateFields) (*mongo.UpdateResult, error)
	DeleteManyByIds(ids []primitive.ObjectID) (*mongo.DeleteResult, error)
}

type transactionOrm struct {
	transactionCollection *mongo.Collection
}

func NewTransactionOrm(db *mongo.Database) *transactionOrm {
	return &transactionOrm{transactionCollection: db.Collection(constants.CollTransactions)}
}

func (o *transactionOrm) InsertOne(insert dt.TransactionInsert) (*mongo.InsertOneResult, error) {
	return o.transactionCollection.InsertOne(context.TODO(), insert)
}

func (o *transactionOrm) InsertMany(inserts []interface{}) (*mongo.InsertManyResult, error) {
	return o.transactionCollection.InsertMany(context.TODO(), inserts)
}

func (o *transactionOrm) GetMany(query dt.TransactionQuery) (transactions []*dt.Transaction, err error) {
	dbQuery := bson.D{{"user", query.User}}
	if query.AfterCursor != nil {
		dbQuery = append(dbQuery, bson.E{"_id", bson.M{"$gt": &query.AfterCursor}})
	}
	if query.Description != nil {
		dbQuery = append(dbQuery, bson.E{"description", bson.M{"$regex": fmt.Sprintf(".*%v.*", *query.Description), "$options": "i"}})
	}
	if query.Account != nil {
		dbQuery = append(dbQuery, bson.E{"account", *query.Account})
	}
	if query.Category != nil {
		dbQuery = append(dbQuery, bson.E{"category", *query.Category})
	}
	if query.DateTimeStart != nil && query.DateTimeEnd != nil {
		dbQuery = append(dbQuery, bson.E{"dateTime", bson.M{"$gt": *query.DateTimeStart, "$lt": *query.DateTimeEnd}})
	}
	if query.AmountMoreThan != nil {
		dbQuery = append(dbQuery, bson.E{"amount", bson.M{"$gt": *query.AmountMoreThan}})
	}
	if query.AmountLessThan != nil {
		dbQuery = append(dbQuery, bson.E{"amount", bson.M{"$lt": *query.AmountLessThan}})
	}
	findOptions := options.Find()
	if query.Limit > 0 {
		findOptions.SetLimit(int64(query.Limit))
	}
	findOptions.SetSort(bson.D{{"_id", 1}})

	cur, err := o.transactionCollection.Find(context.TODO(), dbQuery, findOptions)
	if err != nil {
		return
	}
	err = cur.All(context.TODO(), &transactions)
	if err != nil {
		return
	}

	return
}

func (o *transactionOrm) GetManyNoObjectId(query dt.TransactionQuery) (transactions []*dt.TransactionNoObjectId, err error) {
	dbQuery := bson.D{{"user", query.User}}
	if query.AfterCursor != nil {
		dbQuery = append(dbQuery, bson.E{"_id", bson.M{"$gt": &query.AfterCursor}})
	}
	if query.Description != nil {
		dbQuery = append(dbQuery, bson.E{"description", bson.M{"$regex": fmt.Sprintf(".*%v.*", *query.Description), "$options": "i"}})
	}
	if query.Account != nil {
		dbQuery = append(dbQuery, bson.E{"account", *query.Account})
	}
	if query.Category != nil {
		dbQuery = append(dbQuery, bson.E{"category", *query.Category})
	}
	if query.AmountMoreThan != nil {
		dbQuery = append(dbQuery, bson.E{"amount", bson.M{"$gt": *query.AmountMoreThan}})
	}
	if query.AmountLessThan != nil {
		dbQuery = append(dbQuery, bson.E{"amount", bson.M{"$lt": *query.AmountLessThan}})
	}
	dbOuterQuery := bson.M{"$and": bson.A{
		dbQuery,
		bson.M{"$or": bson.A{
			bson.M{"dateTime": bson.M{"$gt": *query.DateTimeStart, "$lt": *query.DateTimeEnd}},
			bson.M{"isReccurent": true, "reccurenceLastDate": bson.M{"$gt": *query.DateTimeStart}},
		}},
	}}
	findOptions := options.Find()
	if query.Limit > 0 {
		findOptions.SetLimit(int64(query.Limit))
	} else if query.Limit == 0 {
		findOptions.SetLimit(1000)
	}
	findOptions.SetSort(bson.D{{"dateTime", 1}})
	cur, err := o.transactionCollection.Aggregate(context.TODO(), bson.A{
		bson.M{"$match": dbOuterQuery},
		bson.M{"$sort": findOptions.Sort},
		bson.M{"$limit": findOptions.Limit},
		bson.M{"$lookup": bson.D{
			{"from", constants.CollAccounts},
			{"let", bson.M{"accountId": "$account"}},
			{"pipeline", bson.A{
				bson.M{"$match": bson.M{"$expr": bson.M{"$eq": bson.A{"$_id", "$$accountId"}}}},
				bson.M{"$project": bson.M{"account_name": "$name", "_id": 0}},
			}},
			{"as", "fromAccount"},
		}},
		bson.M{"$replaceRoot": bson.M{"newRoot": bson.M{"$mergeObjects": bson.A{
			bson.M{"$arrayElemAt": bson.A{"$fromAccount", 0}},
			"$$ROOT",
		}}}},
		bson.M{"$lookup": bson.D{
			{"from", constants.CollCategories},
			{"let", bson.M{"categoryId": "$category"}},
			{"pipeline", bson.A{
				bson.M{"$match": bson.M{"$expr": bson.M{"$eq": bson.A{"$_id", "$$categoryId"}}}},
				bson.M{"$project": bson.M{"category_name": "$name", "_id": 0}},
			}},
			{"as", "fromCategory"},
		}},
		bson.M{"$replaceRoot": bson.M{"newRoot": bson.M{"$mergeObjects": bson.A{
			bson.M{"$arrayElemAt": bson.A{"$fromCategory", 0}},
			"$$ROOT",
		}}}},
		bson.M{"$addFields": bson.M{"account": "$account_name", "category": "$category_name"}},
		bson.M{"$unset": bson.A{"account_name", "account_category", "fromAccount", "fromCategory", "user"}},
	})

	if err != nil {
		return
	}
	err = cur.All(context.TODO(), &transactions)
	if err != nil {
		return
	}

	return
}

func (o *transactionOrm) GetSomeDescriptionsByPartialDescription(autocomplete *dt.TransactionDescriptionAutocomplete) (suggestions []*dt.TransactionDescriptionAutocompleteResponse, err error) {
	cur, err := o.transactionCollection.Aggregate(context.TODO(), bson.A{
		bson.D{{"$match", bson.D{{"description", bson.M{"$regex": fmt.Sprintf(".*%v.*", *autocomplete.Description), "$options": "i"}}}}},
		bson.D{{"$group", bson.D{
			{"_id", "$description"},
		}}},
		bson.D{{"$limit", autocomplete.Count}},
	})
	err = cur.All(context.TODO(), &suggestions)
	if err != nil {
		return
	}

	return
}

func (o *transactionOrm) GetUserIdsByIds(ids []primitive.ObjectID) (userIds []primitive.ObjectID, err error) {
	var aggregateResult []map[string]primitive.ObjectID
	userIds = make([]primitive.ObjectID, 0)
	cur, err := o.transactionCollection.Aggregate(context.TODO(), bson.A{
		bson.D{{"$match", bson.D{{"_id", bson.D{{"$in", ids}}}}}},
		bson.D{{"$group", bson.D{
			{"_id", "$user"},
		}}},
	})
	if err != nil {
		return
	}
	err = cur.All(context.TODO(), &aggregateResult)
	if err != nil {
		return
	}
	for _, elem := range aggregateResult {
		userIds = append(userIds, elem["_id"])
	}

	return
}

func (o *transactionOrm) GetOneById(id primitive.ObjectID) (transaction *dt.Transaction, err error) {
	err = o.transactionCollection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&transaction)
	if err != nil {
		return
	}

	return
}

func (o *transactionOrm) UpdateManyByIds(ids []primitive.ObjectID, update *dt.TransactionUpdateFields) (updateResult *mongo.UpdateResult, err error) {
	updateResult, err = o.transactionCollection.UpdateMany(context.Background(), bson.D{{"_id", bson.D{{"$in", ids}}}}, bson.D{{"$set", update}})
	if err != nil {
		return
	}

	return
}

func (o *transactionOrm) DeleteManyByIds(ids []primitive.ObjectID) (deleteResult *mongo.DeleteResult, err error) {
	deleteResult, err = o.transactionCollection.DeleteMany(context.TODO(), bson.D{{"_id", bson.D{{"$in", ids}}}})
	if err != nil {
		return
	}

	return
}
