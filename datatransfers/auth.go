package datatransfers

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthLogin struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type AuthRegister struct {
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type AuthPayload struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Email string             `json:"email" bson:"email"`
	Token string             `json:"token" bson:"token"`
}
