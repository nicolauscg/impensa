package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Transaction struct {
	Id          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Amount      float32
	Description string
	DateTime    time.Time
}

func (t *Transaction) String() string {
	return fmt.Sprintf("<Transaction %v %v %v %v>", t.Id, t.Amount, t.Description, t.DateTime)
}

type TransactionOrmer interface {
	InsertOne(amount float32, description string) (*mongo.InsertOneResult, error)
	GetAll() ([]*Transaction, error)
	GetOneById(id string) (*Transaction, error)
	UpdateOneById(id string, update *Transaction) (*mongo.UpdateResult, error)
	DeleteManyById(ids []string) (*mongo.DeleteResult, error)
}

type transactionOrm struct {
	collection *mongo.Collection
}

func NewTransactionOrm(db *mongo.Database) *transactionOrm {
	return &transactionOrm{db.Collection("transactions")}
}

func (o *transactionOrm) InsertOne(amount float32, description string) (*mongo.InsertOneResult, error) {
	return o.collection.InsertOne(context.TODO(), Transaction{Amount: amount, Description: description, DateTime: time.Now()})
}

func (o *transactionOrm) GetAll() (transactions []*Transaction, err error) {
	cur, err := o.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return
	}
	err = cur.All(context.TODO(), &transactions)
	if err != nil {
		return
	}

	return
}

func (o *transactionOrm) GetOneById(id string) (transaction *Transaction, err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	err = o.collection.FindOne(context.TODO(), bson.D{{"_id", objectId}}).Decode(&transaction)
	if err != nil {
		return
	}

	return
}

func (o *transactionOrm) UpdateOneById(id string, update *Transaction) (updateResult *mongo.UpdateResult, err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	update.Id = ""
	if err != nil {
		return
	}
	updateResult, err = o.collection.UpdateOne(context.TODO(), bson.D{{"_id", objectId}}, bson.D{{"$set", update}})
	if err != nil {
		return
	}

	return
}

func (o *transactionOrm) DeleteManyById(ids []string) (deleteResult *mongo.DeleteResult, err error) {
	objectIds := make([]primitive.ObjectID, len(ids))
	for i := 0; i < len(objectIds); i++ {
		objectIds[i], err = primitive.ObjectIDFromHex(ids[i])
	}
	if err != nil {
		return
	}
	deleteResult, err = o.collection.DeleteMany(context.TODO(), bson.D{{"_id", bson.D{{"$in", objectIds}}}})
	if err != nil {
		return
	}

	return
}
