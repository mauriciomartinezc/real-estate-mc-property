package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ManagementType struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name" validate:"required"`
}

type ManagementTypes = []ManagementType
