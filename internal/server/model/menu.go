package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Menu struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitemtpy"`
	StoreId string             `bson:"storeId" json:"storeId"`
	Items   []MenuItem         `bson:"items" json:"items"`
}
