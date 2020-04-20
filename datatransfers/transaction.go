package datatransfers

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id          *primitive.ObjectID `json:"id" bson:"_id"`
	User        *primitive.ObjectID `json:"user" bson:"user"`
	Account     *primitive.ObjectID `json:"account" bson:"account"`
	Category    *primitive.ObjectID `json:"category" bson:"category"`
	Amount      *float32            `json:"amount" bson:"amount"`
	Description *string             `json:"description" bson:"description"`
	DateTime    *time.Time          `json:"dateTime" bson:"dateTime"`
	Picture     *string             `json:"picture" bson:"picture"`
}

func (t *Transaction) String() string {
	return fmt.Sprintf("<Transaction %v %v %v %v>", t.Id, t.Amount, t.Description, t.DateTime)
}

type TransactionInsert struct {
	User        *primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Account     *primitive.ObjectID `json:"account,omitempty" bson:"account,omitempty"`
	Category    *primitive.ObjectID `json:"category,omitempty" bson:"category,omitempty"`
	Amount      *float32            `json:"amount,omitempty" bson:"amount,omitempty"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
	DateTime    *time.Time          `json:"dateTime,omitempty" bson:"dateTime,omitempty"`
	Picture     *string             `json:"picture,omitempty" bson:"picture,omitempty"`
}

type TransactionQuery struct {
	Account        *primitive.ObjectID `json:"account,omitempty" bson:"account,omitempty"`
	Category       *primitive.ObjectID `json:"category,omitempty" bson:"category,omitempty"`
	Description    *string             `json:"description,omitempty" bson:"description,omitempty"`
	DateTimeStart  *time.Time          `json:"dateTimeStart,omitempty" bson:"dateTimeStart,omitempty"`
	DateTimeEnd    *time.Time          `json:"dateTimeEnd,omitempty" bson:"dateTimeEnd,omitempty"`
	AmountMoreThan *float32            `json:"amountMoreThan,omitempty" bson:"amountMoreThan,omitempty"`
	AmountLessThan *float32            `json:"amountLessThan,omitempty" bson:"amountLessThan,omitempty"`
}

type TransactionUpdate struct {
	Ids    []primitive.ObjectID    `json:"ids" bson:"ids"`
	Update TransactionUpdateFields `json:"update" bson:"update"`
}

type TransactionUpdateFields struct {
	Account     *primitive.ObjectID `json:"account,omitempty" bson:"account,omitempty"`
	Category    *primitive.ObjectID `json:"category,omitempty" bson:"category,omitempty"`
	Amount      *float32            `json:"amount,omitempty" bson:"amount,omitempty"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
	DateTime    *time.Time          `json:"dateTime,omitempty" bson:"dateTime,omitempty"`
	Picture     *string             `json:"picture,omitempty" bson:"picture,omitempty"`
}

type TransactionDelete struct {
	Ids []primitive.ObjectID `json:"ids" bson:"ids"`
}
