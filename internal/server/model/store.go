package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Store struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Address string             `bson:"address" json:"address"`
	Name    string             `bson:"name" json:"name"`
	Type    string             `bson:"type" json:"type"`
}
