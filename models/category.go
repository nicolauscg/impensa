package models

import (
	"context"

	"github.com/nicolauscg/impensa/constants"
	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryOrmer interface {
	InsertOne(insert dt.CategoryInsert) (*mongo.InsertOneResult, error)
	GetManyByUserId(id primitive.ObjectID) ([]*dt.Category, error)
	GetUserIdsByIds(ids []primitive.ObjectID) ([]primitive.ObjectID, error)
	GetOneById(id primitive.ObjectID) (*dt.Category, error)
	UpdateManyByIds(ids []primitive.ObjectID, update *dt.CategoryUpdateFields) (*mongo.UpdateResult, error)
	DeleteManyByIds(ids []primitive.ObjectID) (*mongo.DeleteResult, error)
}

type categoryOrm struct {
	categoryCollection *mongo.Collection
}

func NewCategoryOrm(db *mongo.Database) *categoryOrm {
	return &categoryOrm{categoryCollection: db.Collection(constants.CollCategories)}
}

func (o *categoryOrm) InsertOne(insert dt.CategoryInsert) (*mongo.InsertOneResult, error) {
	return o.categoryCollection.InsertOne(context.TODO(), insert)
}

func (o *categoryOrm) GetManyByUserId(userId primitive.ObjectID) (categories []*dt.Category, err error) {
	cur, err := o.categoryCollection.Find(context.TODO(), bson.D{{"user", userId}})
	if err != nil {
		return
	}
	err = cur.All(context.TODO(), &categories)
	if err != nil {
		return
	}

	return
}

func (o *categoryOrm) GetUserIdsByIds(ids []primitive.ObjectID) (userIds []primitive.ObjectID, err error) {
	var aggregateResult []map[string]primitive.ObjectID
	userIds = make([]primitive.ObjectID, 0)
	cur, err := o.categoryCollection.Aggregate(context.TODO(), bson.A{
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

func (o *categoryOrm) GetOneById(id primitive.ObjectID) (category *dt.Category, err error) {
	err = o.categoryCollection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&category)
	if err != nil {
		return
	}

	return
}

func (o *categoryOrm) UpdateManyByIds(ids []primitive.ObjectID, update *dt.CategoryUpdateFields) (updateResult *mongo.UpdateResult, err error) {
	updateResult, err = o.categoryCollection.UpdateMany(context.Background(), bson.D{{"_id", bson.D{{"$in", ids}}}}, bson.D{{"$set", update}})
	if err != nil {
		return
	}

	return
}

func (o *categoryOrm) DeleteManyByIds(ids []primitive.ObjectID) (deleteResult *mongo.DeleteResult, err error) {
	deleteResult, err = o.categoryCollection.DeleteMany(context.TODO(), bson.D{{"_id", bson.D{{"$in", ids}}}})
	if err != nil {
		return
	}

	return
}
