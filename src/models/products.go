package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Products struct {
	Id          primitive.ObjectID `bson:"_id, omitempty"  json:"_id"`
	Name        string             `bson:"name, omitempty" json:"name"`
	Description string             `bson:"description, omitempty" json:"description"`
	Image       string             `bson:"image" json:"image"`
	Price       float64            `bson:"price,omitempty" json:"price"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
	CreatedBy   string             `bson:"createdBy,omitempty" json:"createdBy"`
}
