package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MenuItem struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	StoreId     primitive.ObjectID `bson:"storeId" json:"storeId"`
	Ingredietns []Ingredient       `bson:"ingredients" json:"ingredients"`
	Price       float64            `bson:"price" json:"price"`
}
