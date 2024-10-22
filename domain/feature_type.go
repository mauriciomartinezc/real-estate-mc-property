package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type FeatureType struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

type FeatureTypes = []FeatureType
