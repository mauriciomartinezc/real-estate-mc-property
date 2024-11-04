package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Feature struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `bson:"name" json:"name" validate:"required"`
	FeatureTypeID primitive.ObjectID `bson:"feature_type_id,omitempty" json:"feature_type_id" validate:"required"`
}

type Features = []Feature
