package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Photo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	URL         string             `bson:"url" json:"url"`
	Description string             `bson:"description,omitempty" json:"description"`
	PropertyID  primitive.ObjectID `bson:"property_id,omitempty" json:"property_id"`
}

type Photos = []Photo
