package datatransfers

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	Id   *primitive.ObjectID `json:"id" bson:"_id"`
	Name *string             `json:"name" bson:"name"`
	Icon *int                `json:"icon" bson:"icon"`
	User *primitive.ObjectID `json:"user" bson:"user"`
}

func (a *Category) String() string {
	return fmt.Sprintf("<Category %v %v %v %v>", a.Id, a.Name, a.Icon, a.User)
}

type CategoryInsert struct {
	Id   *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name *string             `json:"name,omitempty" bson:"name,omitempty"`
	Icon *int                `json:"icon,omitempty" bson:"icon,omitempty"`
	User *primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
}

type CategoryUpdate struct {
	Ids    []primitive.ObjectID `json:"ids" bson:"ids"`
	Update CategoryUpdateFields `json:"update" bson:"update"`
}

type CategoryUpdateFields struct {
	Name *string `json:"name,omitempty" bson:"name,omitempty"`
	Icon *int    `json:"icon,omitempty" bson:"icon,omitempty"`
}

type CategoryDelete struct {
	Ids []primitive.ObjectID `json:"ids" bson:"ids"`
}
