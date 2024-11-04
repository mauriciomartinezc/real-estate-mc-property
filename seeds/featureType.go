package seeds

import (
	"context"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func FeatureTypeSeeder(db *mongo.Database) {
	collection := db.Collection("feature_types")

	count, err := collection.CountDocuments(context.Background(), bson.D{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	if count == 0 {

		featureTypes := []interface{}{
			domain.FeatureType{Name: "Exterior"},
			domain.FeatureType{Name: "Interior"},
			domain.FeatureType{Name: "Sector"},
		}

		_, err = collection.InsertMany(context.Background(), featureTypes)
		if err != nil {
			log.Fatalf(err.Error())
		}

		log.Println("Feature Types seeding completado con éxito.")
	}
}
