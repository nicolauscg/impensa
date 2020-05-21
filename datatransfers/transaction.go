package datatransfers

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReccurenceTransactionInterval int

const (
	RepeatDay ReccurenceTransactionInterval = iota
	RepeatWeek
	RepeatMonth
	RepeatYear
)

func (r ReccurenceTransactionInterval) GetTimeFrom(initial time.Time, multiplicant int) time.Time {
	years, months, days := 0, 0, 0
	switch r {
	case RepeatDay:
		days = 1
	case RepeatWeek:
		days = 7
	case RepeatMonth:
		months = 1
	case RepeatYear:
		years = 1
	}
	years *= multiplicant
	months *= multiplicant
	days *= multiplicant
	result := initial.AddDate(years, months, days)
	return result
}

func (r ReccurenceTransactionInterval) String() string {
	return [...]string{"day", "week", "month", "year"}[r]
}

type Transaction struct {
	Id                 *primitive.ObjectID `json:"id" bson:"_id"`
	User               *primitive.ObjectID `json:"user" bson:"user"`
	Account            *primitive.ObjectID `json:"account" bson:"account"`
	Category           *primitive.ObjectID `json:"category" bson:"category"`
	Amount             *float32            `json:"amount" bson:"amount"`
	Description        *string             `json:"description" bson:"description"`
	DateTime           *time.Time          `json:"dateTime" bson:"dateTime"`
	Picture            *string             `json:"picture" bson:"picture"`
	Location           *string             `json:"location" bson:"location"`
	IsReccurent        *bool               `json:"isReccurent" bson:"isReccurent"`
	RepeatCount        *int                `json:"repeatCount" bson:"repeatCount"`
	RepeatInterval     *int                `json:"repeatInterval" bson:"repeatInterval"`
	ReccurenceLastDate *time.Time          `json:"reccurenceLastDate" bson:"reccurenceLastDate"`
}

func (t *Transaction) String() string {
	return fmt.Sprintf("<Transaction %v %v %v %v>", t.Id, t.Amount, t.Description, t.DateTime)
}

type TransactionNoObjectId struct {
	Id                 *primitive.ObjectID            `json:"id" bson:"_id"`
	Account            *string                        `json:"account" bson:"account"`
	Category           *string                        `json:"category" bson:"category"`
	Amount             *float32                       `json:"amount" bson:"amount"`
	Description        *string                        `json:"description" bson:"description"`
	DateTime           *time.Time                     `json:"dateTime" bson:"dateTime"`
	Picture            *string                        `json:"picture" bson:"picture"`
	Location           *string                        `json:"location" bson:"location"`
	IsReccurent        *bool                          `json:"isReccurent" bson:"isReccurent"`
	RepeatCount        *int                           `json:"repeatCount" bson:"repeatCount"`
	RepeatInterval     *ReccurenceTransactionInterval `json:"repeatInterval" bson:"repeatInterval"`
	ReccurenceLastDate *time.Time                     `json:"reccurenceLastDate" bson:"reccurenceLastDate"`
}

func (t *TransactionNoObjectId) CloneWithDifferentDateTime() *TransactionNoObjectId {
	newDateTime := (*t.DateTime).Add(0)
	return &TransactionNoObjectId{
		t.Id, t.Account, t.Category, t.Amount, t.Description, &newDateTime,
		t.Picture, t.Location, t.IsReccurent, t.RepeatCount, t.RepeatInterval, t.ReccurenceLastDate,
	}
}

type TransactionInsert struct {
	User               *primitive.ObjectID            `jso:"user,omitempty" bson:"user,omitempty"`
	Account            *primitive.ObjectID            `json:"account,omitempty" bson:"account,omitempty"`
	Category           *primitive.ObjectID            `json:"category,omitempty" bson:"category,omitempty"`
	Amount             *float32                       `json:"amount,omitempty" bson:"amount,omitempty"`
	Description        *string                        `json:"description,omitempty" bson:"description,omitempty"`
	DateTime           *time.Time                     `json:"dateTime,omitempty" bson:"dateTime,omitempty"`
	Picture            *string                        `json:"picture,omitempty" bson:"picture,omitempty"`
	Location           *string                        `json:"location" bson:"location"`
	IsReccurent        *bool                          `json:"isReccurent" bson:"isReccurent"`
	RepeatCount        *int                           `json:"repeatCount" bson:"repeatCount"`
	RepeatInterval     *ReccurenceTransactionInterval `json:"repeatInterval" bson:"repeatInterval"`
	ReccurenceLastDate *time.Time                     `json:"reccurenceLastDate" bson:"reccurenceLastDate"`
}

type TransactionQuery struct {
	User           *primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Account        *primitive.ObjectID `json:"account,omitempty" bson:"account,omitempty"`
	Category       *primitive.ObjectID `json:"category,omitempty" bson:"category,omitempty"`
	Description    *string             `json:"description,omitempty" bson:"description,omitempty"`
	DateTimeStart  *time.Time          `json:"dateTimeStart,omitempty" bson:"dateTimeStart,omitempty"`
	DateTimeEnd    *time.Time          `json:"dateTimeEnd,omitempty" bson:"dateTimeEnd,omitempty"`
	AmountMoreThan *float32            `json:"amountMoreThan,omitempty" bson:"amountMoreThan,omitempty"`
	AmountLessThan *float32            `json:"amountLessThan,omitempty" bson:"amountLessThan,omitempty"`
	Limit          int                 `json:"limit,omitempty" bson:"limit,omitempty"`
	AfterCursor    *primitive.ObjectID `json:"afterCursor,omitempty" bson:"afterCursor,omitempty"`
}

type TransactionDescriptionAutocomplete struct {
	User        *primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
	Count       int                 `json:"count,omitempty" bson:"count,omitempty"`
}

type TransactionDescriptionAutocompleteResponse struct {
	Id *string `json:"_id,omitempty" bson:"_id,omitempty"`
}

type TransactionUpdate struct {
	Ids    []primitive.ObjectID    `json:"ids" bson:"ids"`
	Update TransactionUpdateFields `json:"update" bson:"update"`
}

type TransactionUpdateFields struct {
	Account            *primitive.ObjectID            `json:"account,omitempty" bson:"account,omitempty"`
	Category           *primitive.ObjectID            `json:"category,omitempty" bson:"category,omitempty"`
	Amount             *float32                       `json:"amount,omitempty" bson:"amount,omitempty"`
	Description        *string                        `json:"description,omitempty" bson:"description,omitempty"`
	DateTime           *time.Time                     `json:"dateTime,omitempty" bson:"dateTime,omitempty"`
	Picture            *string                        `json:"picture,omitempty" bson:"picture,omitempty"`
	Location           *string                        `json:"location" bson:"location"`
	IsReccurent        *bool                          `json:"isReccurent" bson:"isReccurent"`
	RepeatCount        *int                           `json:"repeatCount" bson:"repeatCount"`
	RepeatInterval     *ReccurenceTransactionInterval `json:"repeatInterval" bson:"repeatInterval"`
	ReccurenceLastDate *time.Time                     `json:"reccurenceLastDate" bson:"reccurenceLastDate"`
}

type TransactionDelete struct {
	Ids []primitive.ObjectID `json:"ids" bson:"ids"`
}
