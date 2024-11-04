package seeds

import (
	"context"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func AgeSeeder(db *mongo.Database) {
	collection := db.Collection("ages")

	count, err := collection.CountDocuments(context.Background(), bson.D{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	if count == 0 {

		ages := []interface{}{
			domain.Age{Name: "Menos de un año", Order: 1},
			domain.Age{Name: "De 1 a 8 años", Order: 2},
			domain.Age{Name: "De 9 a 15 años", Order: 3},
			domain.Age{Name: "De 16 a 30 años", Order: 4},
			domain.Age{Name: "Mas de 30 años", Order: 5},
		}

		_, err = collection.InsertMany(context.Background(), ages)
		if err != nil {
			log.Fatalf(err.Error())
		}

		log.Println("Ages seeding completado con éxito.")
	}
}
