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
	GetAll() ([]*dt.Transaction, error)
	GetOneById(id primitive.ObjectID) (*dt.Transaction, error)
	UpdateManyByIds(ids []primitive.ObjectID, update *dt.TransactionInsert) (*mongo.UpdateResult, error)
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

func (o *transactionOrm) GetAll() (transactions []*dt.Transaction, err error) {
	cur, err := o.transactionCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return
	}
	err = cur.All(context.TODO(), &transactions)
	if err != nil {
		return
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

func (o *transactionOrm) UpdateManyByIds(ids []primitive.ObjectID, update *dt.TransactionInsert) (updateResult *mongo.UpdateResult, err error) {
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
