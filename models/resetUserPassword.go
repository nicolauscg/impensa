package models

import (
	"context"
	"time"

	"github.com/nicolauscg/impensa/constants"
	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResetUserPasswordOrmer interface {
	InsertOne(id primitive.ObjectID) (*mongo.InsertOneResult, string, error)
	Verify(id primitive.ObjectID, verifyKey string) (exist bool, err error)
}

type resetUserPasswordOrm struct {
	resetUserPasswordCollection *mongo.Collection
	userCollection              *mongo.Collection
}

func NewResetUserPassword(db *mongo.Database) *resetUserPasswordOrm {
	return &resetUserPasswordOrm{
		resetUserPasswordCollection: db.Collection(constants.CollResetUserPasswords),
		userCollection:              db.Collection(constants.CollUsers),
	}
}

func (o *resetUserPasswordOrm) InsertOne(id primitive.ObjectID) (*mongo.InsertOneResult, string, error) {
	verifyKey := generateRandomString(10)
	now := time.Now()

	insertResult, err := o.resetUserPasswordCollection.InsertOne(context.TODO(), dt.VerifyUser{
		UserId:    &id,
		VerifyKey: &verifyKey,
		CreatedAt: &now,
	})

	return insertResult, verifyKey, err
}

func (o *resetUserPasswordOrm) Verify(id primitive.ObjectID, verifyKey string) (exist bool, err error) {
	var verifyUser dt.VerifyUser

	err = o.resetUserPasswordCollection.FindOne(context.TODO(), bson.D{
		{"userId", id},
		{"verifyKey", verifyKey},
	}).Decode(&verifyUser)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return
	}

	_, err = o.resetUserPasswordCollection.DeleteOne(context.TODO(), bson.D{
		{"userId", id},
	})
	if err != nil {
		return
	}

	singleResult := o.userCollection.FindOneAndUpdate(
		context.TODO(),
		bson.D{{"_id", id}},
		bson.D{{"$set", bson.D{{"verified", true}}}},
	)

	if err := singleResult.Err(); err != nil {
		return false, err
	}
	return true, nil
}
