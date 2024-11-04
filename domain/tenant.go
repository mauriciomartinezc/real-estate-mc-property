package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tenant struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Email     string             `bson:"email,omitempty" json:"email"`
	Phone     string             `bson:"phone,omitempty" json:"phone"`
}

type Tenants = []Tenant
