package model

type OrderItem struct {
	id          string       `bson:"_id,omitempty"`
	quantity    int          `bson:"quantity"`
	price       float64      `bson:"price"`
	ingredients []Ingredient `bson:"ingredients"`
}
