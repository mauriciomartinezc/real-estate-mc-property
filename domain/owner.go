package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Owner struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Email     string             `bson:"email,omitempty"`
	Phone     string             `bson:"phone,omitempty"`
}

type Owners = []Owner
