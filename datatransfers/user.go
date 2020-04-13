package datatransfers

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type UserItem struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Email string             `json:"email" bson:"email"`
}

func (u *User) String() string {
	return fmt.Sprintf("<User %v %v>", u.Id, u.Email)
}

type UserUpdate struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	Update UserUpdateFields   `json:"update" bson:"update"`
}

type UserUpdateFields struct {
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type UserDelete struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
}
