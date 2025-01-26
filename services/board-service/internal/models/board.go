package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Board struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title"`
	OwnerID   string             `bson:"owner_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
