package datatransfers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResetUserPassword struct {
	UserId    *primitive.ObjectID `json:"userId" bson:"userId"`
	VerifyKey *string             `json:"verifyKey" bson:"verifyKey"`
	CreatedAt *time.Time          `json:"createdAt" bson:"createdAt"`
}

type ResetUserPasswordBody struct {
	Email       *string `json:"email,omitempty" bson:"email,omitempty"`
	VerifyKey   *string `json:"verifyKey" bson:"verifyKey"`
	OldPassword *string `json:"oldPassword,omitempty" bson:"oldPassword,omitempty"`
	NewPassword *string `json:"newPassword,omitempty" bson:"newPassword,omitempty"`
}

type RequestResetUserPasswordBody struct {
	Email *string `json:"email,omitempty" bson:"email,omitempty"`
}
