package models

import (
	"context"

	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionOrmer interface {
	InsertOne(insert dt.TransactionInsert) (*mongo.InsertOneResult, error)
	GetManyByOwnerId(ownerId primitive.ObjectID) ([]*dt.Transaction, error)
	GetOwnerIdsByIds(ids []primitive.ObjectID) ([]primitive.ObjectID, error)
	GetOneById(id primitive.ObjectID) (*dt.Transaction, error)
	UpdateManyByIds(ids []primitive.ObjectID, update *dt.TransactionUpdateFields) (*mongo.UpdateResult, error)
	DeleteManyByIds(ids []primitive.ObjectID) (*mongo.DeleteResult, error)
}

type transactionOrm struct {
	transactionCollection *mongo.Collection
}

func NewTransactionOrm(db *mongo.Database) *transactionOrm {
	return &transactionOrm{transactionCollection: db.Collection("transactions")}
}

func (o *transactionOrm) InsertOne(insert dt.TransactionInsert) (*mongo.InsertOneResult, error) {
	return o.transactionCollection.InsertOne(context.TODO(), insert)
}

func (o *transactionOrm) GetManyByOwnerId(ownerId primitive.ObjectID) (transactions []*dt.Transaction, err error) {
	cur, err := o.transactionCollection.Find(context.TODO(), bson.D{{"owner", ownerId}})
	if err != nil {
		return
	}
	err = cur.All(context.TODO(), &transactions)
	if err != nil {
		return
	}

	return
}

func (o *transactionOrm) GetOwnerIdsByIds(ids []primitive.ObjectID) (ownerIds []primitive.ObjectID, err error) {
	var aggregateResult []map[string]primitive.ObjectID
	ownerIds = make([]primitive.ObjectID, 0)
	cur, err := o.transactionCollection.Aggregate(context.TODO(), bson.A{
		bson.D{{"$match", bson.D{{"_id", bson.D{{"$in", ids}}}}}},
		bson.D{{"$group", bson.D{
			{"_id", "$owner"},
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
		ownerIds = append(ownerIds, elem["_id"])
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
