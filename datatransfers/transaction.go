package datatransfers

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
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

func (t *TransactionInsert) UnMarshalCSV(line []string) (err error) {
	var accountId, categoryId *primitive.ObjectID

	if line[0] != "" {
		temp, err := primitive.ObjectIDFromHex(line[0])
		accountId = &temp
		if err != nil {
			return err
		}
	} else {
		accountId = nil
	}

	if line[1] != "" {
		temp, err := primitive.ObjectIDFromHex(line[1])
		categoryId = &temp
		if err != nil {
			return err
		}
	} else {
		categoryId = nil
	}

	amount, err := strconv.ParseFloat(line[2], 32)
	if err != nil {
		return
	}
	amount32 := float32(amount)
	dateTime, err := time.Parse(time.RFC3339, line[4])
	if err != nil {
		return
	}
	var isReccurent *bool
	if line[7] != "" {
		temp, err := strconv.ParseBool(line[7])
		isReccurent = &temp
		if err != nil {

			return err
		}
	} else {
		isReccurent = nil
	}
	var repeatCount *int
	if line[8] != "" {
		temp, err := strconv.ParseInt(line[8], 10, 32)
		temp2 := int(temp)
		repeatCount = &temp2
		if err != nil {

			return err
		}
	} else {
		repeatCount = nil
	}
	var repeatInterval *ReccurenceTransactionInterval
	if line[9] != "" {
		temp, err := strconv.ParseInt(line[9], 10, 32)
		temp2 := int(temp)
		temp3 := ReccurenceTransactionInterval(temp2)
		repeatInterval = &temp3

		if err != nil {

			return err
		}
	} else {
		repeatInterval = nil
	}
	t.Account = accountId
	t.Category = categoryId
	t.Amount = &amount32
	t.Description = &line[3]
	t.DateTime = &dateTime
	t.Picture = &line[5]
	t.Location = &line[6]
	t.IsReccurent = isReccurent
	if isReccurent != nil && *isReccurent {
		t.RepeatCount = repeatCount
		t.RepeatInterval = repeatInterval
		temp2 := repeatInterval.GetTimeFrom(dateTime, *repeatCount)
		t.ReccurenceLastDate = &temp2
	}

	return
}

func (t *Transaction) MarshalCSV() ([]string, error) {
	var err error
	csv := make([]string, 0)
	transactionType := reflect.TypeOf(t).Elem()
	transactionValues := reflect.ValueOf(t).Elem()
	for i := 0; i < transactionType.NumField(); i++ {
		if i == 0 || i == 1 {
			continue
		}
		value := transactionValues.Field(i)
		switch fieldType := value.Interface().(type) {
		default:
			err = errors.New(fmt.Sprintf("invalid type %v when marshal transaction to csv", fieldType))
			break
		case *primitive.ObjectID:
			if v := value.Interface().(*primitive.ObjectID); v != nil {
				csv = append(csv, v.Hex())
			} else {
				csv = append(csv, "")
			}
		case *time.Time:
			if v := value.Interface().(*time.Time); v != nil {
				csv = append(csv, v.Format(time.RFC3339))
			} else {
				csv = append(csv, "")
			}
		case *float32:
			if v := value.Interface().(*float32); v != nil {
				csv = append(csv, fmt.Sprintf("%.2f", *v))
			} else {
				csv = append(csv, "")
			}
		case *string:
			if v := value.Interface().(*string); v != nil {
				csv = append(csv, *v)
			} else {
				csv = append(csv, "")
			}
		case *bool:
			if v := value.Interface().(*bool); v != nil {
				csv = append(csv, fmt.Sprintf("%v", *v))
			} else {
				csv = append(csv, "")
			}
		case *int:
			if v := value.Interface().(*int); v != nil {
				csv = append(csv, fmt.Sprintf("%v", *v))
			} else {
				csv = append(csv, "")
			}
		}
	}

	return csv, err
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

type TransactionImportCsv struct {
	Csv *string `json:"csv,omitempty" bson:"csv,omitempty"`
}

type TransactionExportCsv struct {
	DateTimeStart *time.Time `json:"dateTimeStart,omitempty" bson:"dateTimeStart,omitempty"`
	DateTimeEnd   *time.Time `json:"dateTimeEnd,omitempty" bson:"dateTimeEnd,omitempty"`
}
