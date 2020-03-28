package datatransfers

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthLogin struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type AuthRegister struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type AuthPayload struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Email string             `json:"email" bson:"email"`
	Token string             `json:"token" bson:"token"`
}
