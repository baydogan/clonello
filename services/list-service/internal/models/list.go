package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type List struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title"`
	BoardID string             `bson:"board_id"`
}
