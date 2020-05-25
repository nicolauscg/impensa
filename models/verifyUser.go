package models

import (
	"context"
	"math/rand"
	"time"

	"github.com/nicolauscg/impensa/constants"
	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VerifyUserOrmer interface {
	InsertOne(id primitive.ObjectID) (*mongo.InsertOneResult, string, error)
	Verify(id primitive.ObjectID, verifyKey string) (exist bool, err error)
}

type verifyUserOrm struct {
	verifyUserCollection *mongo.Collection
	userCollection       *mongo.Collection
}

func NewVerifyUserOrm(db *mongo.Database) *verifyUserOrm {
	return &verifyUserOrm{
		verifyUserCollection: db.Collection(constants.CollVerifyUsers),
		userCollection:       db.Collection(constants.CollUsers),
	}
}

func (o *verifyUserOrm) InsertOne(id primitive.ObjectID) (*mongo.InsertOneResult, string, error) {
	verifyKey := generateRandomString(10)
	now := time.Now()

	insertResult, err := o.verifyUserCollection.InsertOne(context.TODO(), dt.VerifyUser{
		UserId:    &id,
		VerifyKey: &verifyKey,
		CreatedAt: &now,
	})

	return insertResult, verifyKey, err
}

func (o *verifyUserOrm) Verify(id primitive.ObjectID, verifyKey string) (exist bool, err error) {
	var verifyUser dt.VerifyUser

	err = o.verifyUserCollection.FindOne(context.TODO(), bson.D{
		{"userId", id},
		{"verifyKey", verifyKey},
	}).Decode(&verifyUser)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return
	}

	_, err = o.verifyUserCollection.DeleteOne(context.TODO(), bson.D{
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

func generateRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
