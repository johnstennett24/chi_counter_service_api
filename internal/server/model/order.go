package model

type Order struct {
	Id    string      `bson:"_id,omitempty"`
	Items []OrderItem `bson:"items"`
}
