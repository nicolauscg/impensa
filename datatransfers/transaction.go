package datatransfers

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	Amount      float32            `json:"amount" bson:"amount"`
	Description string             `json:"description" bson:"description"`
	DateTime    time.Time          `json:"dateTime" bson:"dateTime"`
}

func (t *Transaction) String() string {
	return fmt.Sprintf("<Transaction %v %v %v %v>", t.Id, t.Amount, t.Description, t.DateTime)
}

type TransactionInsert struct {
	Amount      float32   `bson:"amount"`
	Description string    `bson:"description"`
	DateTime    time.Time `bson:"dateTime"`
}

type TransactionUpdate struct {
	Ids    []primitive.ObjectID
	Update TransactionUpdateFields
}

type TransactionUpdateFields struct {
	Amount      float32   `bson:"amount,omitempty"`
	Description string    `bson:"description,omitempty"`
	DateTime    time.Time `bson:"dateTime,omitempty"`
}

type TransactionDelete struct {
	Ids []primitive.ObjectID
}
