package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	StoreId   string             `bson:"storeId" json:"storeId"`
	IsManager bool               `bson:"isManager" json:"isManager"`
	Wage      float64            `bson:"wage" json:"wage"`
}
