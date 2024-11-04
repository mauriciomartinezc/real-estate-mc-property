package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Age struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name" validate:"required"`
	Order int64              `bson:"order" json:"order" validate:"required"`
}

type Ages = []Age
