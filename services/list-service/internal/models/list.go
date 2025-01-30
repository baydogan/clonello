package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type List struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BoardID   string             `bson:"board_id" json:"board_id"`
	Title     string             `bson:"title" json:"title"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
