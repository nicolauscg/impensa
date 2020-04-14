package datatransfers

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Owner       primitive.ObjectID `json:"owner" bson:"owner"`
	Amount      float32            `json:"amount" bson:"amount"`
	Description string             `json:"description" bson:"description"`
	DateTime    time.Time          `json:"dateTime" bson:"dateTime"`
}

func (t *Transaction) String() string {
	return fmt.Sprintf("<Transaction %v %v %v %v>", t.Id, t.Amount, t.Description, t.DateTime)
}

type TransactionInsert struct {
	Owner       primitive.ObjectID `json:"owner,omitempty" bson:"owner,omitempty"`
	Amount      float32            `json:"amount,omitempty" bson:"amount,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	DateTime    time.Time          `json:"dateTime,omitempty" bson:"dateTime,omitempty"`
}

type TransactionUpdate struct {
	Ids    []primitive.ObjectID    `json:"ids" bson:"ids"`
	Update TransactionUpdateFields `json:"update" bson:"update"`
}

type TransactionUpdateFields struct {
	Amount      float32   `json:"amount,omitempty" bson:"amount,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	DateTime    time.Time `json:"dateTime,omitempty" bson:"dateTime,omitempty"`
}

type TransactionDelete struct {
	Ids []primitive.ObjectID `json:"ids" bson:"ids"`
}
