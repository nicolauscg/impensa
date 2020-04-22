package datatransfers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PieChartSliceInfo struct {
	Id       *primitive.ObjectID `json:"id" bson:"_id"`
	Label    string              `json:"label" bson:"label"`
	Quantity float32             `json:"quantity" bson:"quantity"`
}
