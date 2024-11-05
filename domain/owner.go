package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Owner struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ClientID   primitive.ObjectID `bson:"client_id" json:"client_id"`
	PropertyID primitive.ObjectID `bson:"property_id" json:"property_id"`
}

type Owners = []Owner
