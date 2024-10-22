package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Feature struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	FeatureTypeID primitive.ObjectID `bson:"feature_type_id,omitempty"` // Relaciona con el tipo de característica
}

type Features = []Feature
