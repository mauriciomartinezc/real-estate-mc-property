package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Photo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	URL         string             `bson:"url"` // URL de la foto
	Description string             `bson:"description,omitempty"`
	PropertyID  primitive.ObjectID `bson:"property_id,omitempty"` // ID de la propiedad asociada
}

type Photos = []Photo
