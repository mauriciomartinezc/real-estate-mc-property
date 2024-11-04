package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	URL        string             `bson:"url" json:"url"`
	PropertyID primitive.ObjectID `bson:"property_id,omitempty" json:"property_id"`
}

type Documents = []Document
