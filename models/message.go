package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Sender    primitive.ObjectID `bson:"sender" json:"sender"`
	Recipient primitive.ObjectID `bson:"recipient" json:"recipient"`
	Text      string             `bson:"text" json:"text"`
	File      string             `bson:"file,omitempty" json:"file,omitempty"`
}