package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ingredient struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	name          string             `bson:"name"`
	costPerUnit   float64            `bson:"cost"`
	unit          string             `bson:"unit"`
	amountInStock int                `bson:"stock"`
	storeId       string             `bson:"storeId"`
}
