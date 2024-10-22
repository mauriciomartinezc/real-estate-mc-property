package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Property struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty"`
	Price           float64              `bson:"price"`
	PriceAdmin      float64              `bson:"price_admin"`
	AdminIncluded   bool                 `bson:"admin_included"`
	Address         string               `bson:"address"`
	FullAddress     string               `bson:"full_address"`
	Longitude       float64              `bson:"longitude"`
	Latitude        float64              `bson:"latitude"`
	Stratum         int                  `bson:"stratum,omitempty"`
	Area            int                  `bson:"area"`
	BuiltArea       int                  `bson:"built_area,omitempty"`
	Rooms           int                  `bson:"rooms"`
	Bathrooms       int                  `bson:"bathrooms"`
	Garages         int                  `bson:"garages"`
	Featured        bool                 `bson:"featured"`
	Slug            string               `bson:"slug,omitempty"`
	StateEstateID   string               `bson:"state_estate_id"`
	AgeID           primitive.ObjectID   `bson:"age_id,omitempty"`
	PropertyTypeID  primitive.ObjectID   `bson:"property_type_id,omitempty"`
	ManagementTypes []primitive.ObjectID `bson:"management_types,omitempty"`
	FeatureIDs      []primitive.ObjectID `bson:"feature_ids,omitempty"`
	OwnerIDs        []primitive.ObjectID `bson:"owner_ids,omitempty"`
	TenantIDs       []primitive.ObjectID `bson:"tenant_ids,omitempty"`
	Documents       Documents            `bson:"documents,omitempty"`
	Photos          Photos               `bson:"photos,omitempty"`
	CreatedAt       int64                `bson:"created_at,omitempty"`
	UpdatedAt       int64                `bson:"updated_at,omitempty"`
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
