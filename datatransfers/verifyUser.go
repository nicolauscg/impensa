package datatransfers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VerifyUser struct {
	UserId    *primitive.ObjectID `json:"userId" bson:"userId"`
	VerifyKey *string             `json:"verifyKey" bson:"verifyKey"`
	CreatedAt *time.Time          `json:"createdAt" bson:"createdAt"`
}
