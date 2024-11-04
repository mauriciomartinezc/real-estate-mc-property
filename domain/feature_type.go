package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type FeatureType struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name" validate:"required"`
}

type FeatureTypes = []FeatureType
