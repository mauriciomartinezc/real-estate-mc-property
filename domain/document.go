package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	URL        string             `bson:"url"`                   // URL o ruta del documento
	PropertyID primitive.ObjectID `bson:"property_id,omitempty"` // ID de la propiedad asociada
}

type Documents = []Document
