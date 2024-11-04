package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Property struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Price           float64              `bson:"price" json:"price"`
	PriceAdmin      float64              `bson:"price_admin" json:"price_admin"`
	AdminIncluded   bool                 `bson:"admin_included" json:"admin_included"`
	Address         string               `bson:"address" json:"address"`
	FullAddress     string               `bson:"full_address" json:"full_address"`
	Longitude       float64              `bson:"longitude" json:"longitude"`
	Latitude        float64              `bson:"latitude" json:"latitude"`
	Stratum         int                  `bson:"stratum,omitempty" json:"stratum"`
	Area            int                  `bson:"area" json:"area"`
	BuiltArea       int                  `bson:"built_area,omitempty" json:"built_area"`
	Rooms           int                  `bson:"rooms" json:"rooms"`
	Bathrooms       int                  `bson:"bathrooms" json:"bathrooms"`
	Garages         int                  `bson:"garages" json:"garages"`
	Featured        bool                 `bson:"featured" json:"featured"`
	Slug            string               `bson:"slug,omitempty" json:"slug,omitempty"`
	StateEstateID   string               `bson:"state_estate_id" json:"state_estate_id"`
	AgeID           primitive.ObjectID   `bson:"age_id,omitempty" json:"age_id,omitempty"`
	PropertyTypeID  primitive.ObjectID   `bson:"property_type_id,omitempty" json:"property_type_id,omitempty"`
	ManagementTypes []primitive.ObjectID `bson:"management_types,omitempty" json:"management_types,omitempty"`
	FeatureIDs      []primitive.ObjectID `bson:"feature_ids,omitempty" json:"feature_ids,omitempty"`
	OwnerIDs        []primitive.ObjectID `bson:"owner_ids,omitempty" json:"owner_ids,omitempty"`
	TenantIDs       []primitive.ObjectID `bson:"tenant_ids,omitempty" json:"tenant_ids,omitempty"`
	Documents       Documents            `bson:"documents,omitempty" json:"documents,omitempty"`
	Photos          Photos               `bson:"photos,omitempty" json:"photos,omitempty"`
	CreatedAt       int64                `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt       int64                `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Properties []Property

func (m *Property) BeforeCreate(ctx *mongo.SessionContext) (err error) {
	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now().Unix()
	m.UpdatedAt = time.Now().Unix()
	return
}

func (m *Property) BeforeUpdate(ctx *mongo.SessionContext) (err error) {
	m.UpdatedAt = time.Now().Unix()
	return
}
