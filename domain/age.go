package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Age struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name" validate:"required"`
	Order int64              `bson:"order" validate:"required"`
}

type Ages = []Age
