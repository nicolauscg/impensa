package datatransfers

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Picture  string             `json:"picture" bson:"picture"`
	Verified bool               `json:"verified" bson:"verified"`
}

type UserItem struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Picture  string             `json:"picture" bson:"picture"`
	Verified bool               `json:"verified" bson:"verified"`
}

func (u *User) String() string {
	return fmt.Sprintf("<User %v %v>", u.Id, u.Email)
}

type UserUpdate struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	Update UserUpdateFields   `json:"update" bson:"update"`
}

type UserUpdateFields struct {
	Username *string `json:"username,omitempty" bson:"username,omitempty"`
	Email    *string `json:"email,omitempty" bson:"email,omitempty"`
	Picture *string `json:"picture,omitempty" bson:"picture,omitempty"`
}

type UserUpdateFieldsInModel struct {
	Username *string `json:"username,omitempty" bson:"username,omitempty"`
	Email    *string `json:"email,omitempty" bson:"email,omitempty"`
	Password *string `json:"password,omitempty" bson:"password,omitempty"`
	Picture  *string `json:"picture,omitempty" bson:"picture,omitempty"`
}

type UserDelete struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
}
