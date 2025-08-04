package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderID      string             `bson:"orderId" json:"orderId"`
	UserID       primitive.ObjectID `bson:"userId" json:"userId"`
	ProductID    primitive.ObjectID `bson:"productId" json:"productId"`
	ProductPrice float64            `bson:"productPrice" json:"productPrice"`
	PaidPrice    float64            `bson:"paidPrice" json:"paidPrice"`
	CreatedAt    primitive.DateTime `bson:"createdAt" json:"createdAt"`
}
