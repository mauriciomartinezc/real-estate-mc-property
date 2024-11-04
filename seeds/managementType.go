package seeds

import (
	"context"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func ManagementTypeSeeder(db *mongo.Database) {
	collection := db.Collection("management_types")

	count, err := collection.CountDocuments(context.Background(), bson.D{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	if count == 0 {

		managementTypes := []interface{}{
			domain.ManagementType{Name: "Alquiler vacacional"},
			domain.ManagementType{Name: "Arriendo"},
			domain.ManagementType{Name: "Permuto"},
			domain.ManagementType{Name: "Venta"},
		}

		_, err = collection.InsertMany(context.Background(), managementTypes)
		if err != nil {
			log.Fatalf(err.Error())
		}

		log.Println("Management Types seeding completado con éxito.")
	}
}
