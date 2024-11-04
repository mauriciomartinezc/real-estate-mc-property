package seeds

import (
	"context"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func PropertyTypeSeeder(db *mongo.Database) {
	collection := db.Collection("property_types")

	count, err := collection.CountDocuments(context.Background(), bson.D{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	if count == 0 {

		propertyTypes := []interface{}{
			domain.PropertyType{Name: "Apartaestudio"},
			domain.PropertyType{Name: "Apartamento"},
			domain.PropertyType{Name: "Bodega"},
			domain.PropertyType{Name: "Cabaña"},
			domain.PropertyType{Name: "Casa"},
			domain.PropertyType{Name: "Casa campestre"},
			domain.PropertyType{Name: "Casa lote"},
			domain.PropertyType{Name: "Finca"},
			domain.PropertyType{Name: "Habitación"},
			domain.PropertyType{Name: "Lote"},
			domain.PropertyType{Name: "Oficina"},
			domain.PropertyType{Name: "Parqueadero"},
		}

		_, err = collection.InsertMany(context.Background(), propertyTypes)
		if err != nil {
			log.Fatalf(err.Error())
		}

		log.Println("Property Types seeding completado con éxito.")
	}
}
