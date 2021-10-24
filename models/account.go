package models

import (
	"context"

	"github.com/nicolauscg/impensa/constants"
	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountOrmer interface {
	InsertOne(insert dt.AccountInsert) (*mongo.InsertOneResult, error)
	GetManyByUserId(id primitive.ObjectID) ([]*dt.Account, error)
	GetUserIdsByIds(ids []primitive.ObjectID) ([]primitive.ObjectID, error)
	GetOneById(id primitive.ObjectID) (*dt.Account, error)
	UpdateManyByIds(ids []primitive.ObjectID, update *dt.AccountUpdateFields) (*mongo.UpdateResult, error)
	DeleteManyByIds(ids []primitive.ObjectID) (*mongo.DeleteResult, error)
}

type accountOrm struct {
	accountCollection *mongo.Collection
}

func NewAccountOrm(db *mongo.Database) *accountOrm {
	return &accountOrm{accountCollection: db.Collection(constants.CollAccounts)}
}

func (o *accountOrm) InsertOne(insert dt.AccountInsert) (*mongo.InsertOneResult, error) {
	return o.accountCollection.InsertOne(context.TODO(), insert)
}

func (o *accountOrm) GetManyByUserId(userId primitive.ObjectID) (accounts []*dt.Account, err error) {
	cur, err := o.accountCollection.Find(context.TODO(), bson.D{{"user", userId}})
	if err != nil {
		return
	}
	err = cur.All(context.TODO(), &accounts)
	if err != nil {
		return
	}

	return
}

func (o *accountOrm) GetUserIdsByIds(ids []primitive.ObjectID) (userIds []primitive.ObjectID, err error) {
	var aggregateResult []map[string]primitive.ObjectID
	userIds = make([]primitive.ObjectID, 0)
	cur, err := o.accountCollection.Aggregate(context.TODO(), bson.A{
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

func (o *accountOrm) GetOneById(id primitive.ObjectID) (account *dt.Account, err error) {
	err = o.accountCollection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&account)
	if err != nil {
		return
	}

	return
}

func (o *accountOrm) UpdateManyByIds(ids []primitive.ObjectID, update *dt.AccountUpdateFields) (updateResult *mongo.UpdateResult, err error) {
	updateResult, err = o.accountCollection.UpdateMany(context.Background(), bson.D{{"_id", bson.D{{"$in", ids}}}}, bson.D{{"$set", update}})
	if err != nil {
		return
	}

	return
}

func (o *accountOrm) DeleteManyByIds(ids []primitive.ObjectID) (deleteResult *mongo.DeleteResult, err error) {
	deleteResult, err = o.accountCollection.DeleteMany(context.TODO(), bson.D{{"_id", bson.D{{"$in", ids}}}})
	if err != nil {
		return
	}

	return
}
