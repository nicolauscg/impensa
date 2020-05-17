package models

import (
	"context"

	"github.com/nicolauscg/impensa/constants"
	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserOrmer interface {
	InsertOne(insert dt.AuthRegister) (*mongo.InsertOneResult, error)
	GetOneById(id primitive.ObjectID) (*dt.UserItem, error)
	GetOneWithPasswordById(id primitive.ObjectID) (*dt.User, error)
	GetOneByEmail(email string) (*dt.User, error)
	GetOneByUsername(username string) (*dt.User, error)
	UpdateOneById(id primitive.ObjectID, update *dt.UserUpdateFieldsInModel) (*mongo.UpdateResult, error)
	DeleteOneById(id primitive.ObjectID) (*mongo.DeleteResult, error)
}

type userOrm struct {
	userCollection *mongo.Collection
}

func NewUserOrm(db *mongo.Database) *userOrm {
	return &userOrm{userCollection: db.Collection(constants.CollUsers)}
}

func (o *userOrm) InsertOne(insert dt.AuthRegister) (*mongo.InsertOneResult, error) {
	return o.userCollection.InsertOne(context.TODO(), insert)
}

func (o *userOrm) GetOneById(id primitive.ObjectID) (user *dt.UserItem, err error) {
	err = o.userCollection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&user)
	if err != nil {
		return
	}

	return
}

func (o *userOrm) GetOneWithPasswordById(id primitive.ObjectID) (user *dt.User, err error) {
	err = o.userCollection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&user)
	if err != nil {
		return
	}

	return
}

func (o *userOrm) GetOneByEmail(email string) (user *dt.User, err error) {
	err = o.userCollection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		return
	}

	return
}

func (o *userOrm) GetOneByUsername(username string) (user *dt.User, err error) {
	err = o.userCollection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		return
	}

	return
}

func (o *userOrm) UpdateOneById(id primitive.ObjectID, update *dt.UserUpdateFieldsInModel) (updateResult *mongo.UpdateResult, err error) {
	updateResult, err = o.userCollection.UpdateMany(context.Background(), bson.D{{"_id", id}}, bson.D{{"$set", update}})
	if err != nil {
		return
	}

	return
}

func (o *userOrm) DeleteOneById(id primitive.ObjectID) (deleteResult *mongo.DeleteResult, err error) {
	deleteResult, err = o.userCollection.DeleteMany(context.TODO(), bson.D{{"_id", id}})
	if err != nil {
		return
	}

	return
}
