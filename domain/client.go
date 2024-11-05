package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName string             `bson:"first_name" json:"first_name" validate:"required"`
	LastName  string             `bson:"last_name" json:"last_name" validate:"required"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	Phone     string             `bson:"phone" json:"phone" validate:"required"`
	CityID    string             `bson:"city_id" json:"city_id" validate:"required"`
	CompanyId string             `bson:"company_id" json:"company_id" validate:"required"`
}

type Clients []Client
