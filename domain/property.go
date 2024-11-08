package domain

import (
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Property struct {
	ID                 primitive.ObjectID   `bson:"_id" json:"id"`
	Price              float64              `bson:"price" json:"price" validate:"required"`
	PriceAdmin         float64              `bson:"price_admin" json:"price_admin" validate:"required"`
	AdminIncluded      bool                 `bson:"admin_included" json:"admin_included" validate:"required"`
	Address            string               `bson:"address" json:"address" validate:"required"`
	FullAddress        string               `bson:"full_address" json:"full_address" validate:"required"`
	Longitude          float64              `bson:"longitude" json:"longitude" validate:"required"`
	Latitude           float64              `bson:"latitude" json:"latitude" validate:"required"`
	Area               int                  `bson:"area" json:"area" validate:"required"`
	BuiltArea          int                  `bson:"built_area" json:"built_area" validate:"required"`
	Rooms              int                  `bson:"rooms" json:"rooms" validate:"required"`
	Bathrooms          int                  `bson:"bathrooms" json:"bathrooms" validate:"required"`
	Garages            int                  `bson:"garages" json:"garages" validate:"required"`
	Slug               string               `bson:"slug" json:"slug"`
	Active             bool                 `bson:"active" json:"active"`
	CompanyID          string               `bson:"company_id" json:"company_id" validate:"required"`
	CityID             string               `bson:"city_id" json:"city_id" validate:"required"`
	NeighborhoodID     string               `bson:"neighborhood_id" json:"neighborhood_id" validate:"required"`
	AgeID              primitive.ObjectID   `bson:"age_id" json:"age_id" validate:"required"`
	PropertyTypeID     primitive.ObjectID   `bson:"property_type_id" json:"property_type_id" validate:"required"`
	ManagementTypesIDs []primitive.ObjectID `bson:"management_types_ids" json:"management_types_ids" validate:"required"`
	Age                Age                  `bson:"age" json:"age"`
	PropertyType       PropertyType         `bson:"property_type" json:"property_type"`
	ManagementTypes    ManagementTypes      `bson:"management_types" json:"management_types"`
	City               domain.City          `bson:"city" json:"city"`
	Neighborhood       domain.Neighborhood  `bson:"neighborhood" json:"neighborhood"`
	Photo              Photo                `bson:"photo" json:"photo"`
	FeatureIDs         []primitive.ObjectID `bson:"feature_ids,omitempty" json:"feature_ids,omitempty"`
	OwnerIDs           []primitive.ObjectID `bson:"owner_ids,omitempty" json:"owner_ids,omitempty"`
	TenantIDs          []primitive.ObjectID `bson:"tenant_ids,omitempty" json:"tenant_ids,omitempty"`
	Features           Features             `bson:"features,omitempty" json:"features,omitempty"`
	Owners             Owners               `bson:"owners,omitempty" json:"owners,omitempty"`
	Tenants            Tenants              `bson:"tenants,omitempty" json:"tenants,omitempty"`
	Documents          Documents            `bson:"documents,omitempty" json:"documents,omitempty"`
	Photos             Photos               `bson:"photos,omitempty" json:"photos,omitempty"`
	CreatedAt          int64                `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt          int64                `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type SimpleProperty struct {
	ID              primitive.ObjectID  `bson:"_id" json:"id"`
	Price           float64             `bson:"price" json:"price"`
	PriceAdmin      float64             `bson:"price_admin" json:"price_admin"`
	AdminIncluded   bool                `bson:"admin_included" json:"admin_included"`
	Address         string              `bson:"address" json:"address"`
	FullAddress     string              `bson:"full_address" json:"full_address"`
	Longitude       float64             `bson:"longitude" json:"longitude"`
	Latitude        float64             `bson:"latitude" json:"latitude"`
	Area            int                 `bson:"area" json:"area"`
	BuiltArea       int                 `bson:"built_area" json:"built_area"`
	Rooms           int                 `bson:"rooms" json:"rooms"`
	Bathrooms       int                 `bson:"bathrooms" json:"bathrooms"`
	Garages         int                 `bson:"garages" json:"garages"`
	Slug            string              `bson:"slug" json:"slug"`
	Active          bool                `bson:"active" json:"active"`
	Age             Age                 `bson:"age" json:"age"`
	PropertyType    PropertyType        `bson:"property_type" json:"property_type"`
	ManagementTypes ManagementTypes     `bson:"management_types" json:"management_types"`
	City            domain.City         `bson:"city" json:"city"`
	Neighborhood    domain.Neighborhood `bson:"neighborhood" json:"neighborhood"`
	Photo           Photo               `bson:"photo" json:"photo"`
	CreatedAt       int64               `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt       int64               `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type DetailProperty struct {
	ID              primitive.ObjectID  `bson:"_id" json:"id"`
	Price           float64             `bson:"price" json:"price"`
	PriceAdmin      float64             `bson:"price_admin" json:"price_admin"`
	AdminIncluded   bool                `bson:"admin_included" json:"admin_included"`
	Address         string              `bson:"address" json:"address"`
	FullAddress     string              `bson:"full_address" json:"full_address"`
	Longitude       float64             `bson:"longitude" json:"longitude"`
	Latitude        float64             `bson:"latitude" json:"latitude"`
	Area            int                 `bson:"area" json:"area"`
	BuiltArea       int                 `bson:"built_area" json:"built_area"`
	Rooms           int                 `bson:"rooms" json:"rooms"`
	Bathrooms       int                 `bson:"bathrooms" json:"bathrooms"`
	Garages         int                 `bson:"garages" json:"garages"`
	Slug            string              `bson:"slug" json:"slug"`
	Age             Age                 `bson:"age" json:"age"`
	PropertyType    PropertyType        `bson:"property_type" json:"property_type"`
	ManagementTypes ManagementTypes     `bson:"management_types" json:"management_types"`
	City            domain.City         `bson:"city" json:"city"`
	Neighborhood    domain.Neighborhood `bson:"neighborhood" json:"neighborhood"`
	Features        Features            `bson:"features,omitempty" json:"features,omitempty"`
	Photos          Photos              `bson:"photos,omitempty" json:"photos,omitempty"`
}

type Properties []Property
type SimpleProperties []SimpleProperty
