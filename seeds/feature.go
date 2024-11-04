package seeds

import (
	"context"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func FeatureSeeder(db *mongo.Database) {
	collection := db.Collection("features")
	featureTypeCollection := db.Collection("feature_types")

	count, _ := collection.CountDocuments(context.Background(), bson.D{})

	if count == 0 {
		var exteriorType domain.FeatureType
		var interiorType domain.FeatureType
		var sectorType domain.FeatureType

		err := featureTypeCollection.FindOne(context.Background(), bson.M{"name": "Exterior"}).Decode(&exteriorType)
		if err != nil {
			log.Fatalf("Error al buscar tipo de característica 'Exterior': %v", err)
		}

		err = featureTypeCollection.FindOne(context.Background(), bson.M{"name": "Interior"}).Decode(&interiorType)
		if err != nil {
			log.Fatalf("Error al buscar tipo de característica 'Interior': %v", err)
		}

		err = featureTypeCollection.FindOne(context.Background(), bson.M{"name": "Sector"}).Decode(&sectorType)
		if err != nil {
			log.Fatalf("Error al buscar tipo de característica 'Sector': %v", err)
		}

		features := []interface{}{
			domain.Feature{Name: "Aire Acondicionado", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Alarma", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Amoblado", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Balcón", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Baño Auxiliar", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Calentador", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Chimenea", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Citófono", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Cocina Integral", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Cocina tipo Americano", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Comedor auxiliar", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Cuarto de conductores", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Cuarto de servicio", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Depósito/Bodega", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Despensa", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Estudio", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Hall de Alcobas", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Instalación de gas", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Patio", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Piso en Baldosa/Mármol", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Zona de lavandería", FeatureTypeID: interiorType.ID},
			domain.Feature{Name: "Ascensor", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Cancha de Tennis", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Canchas Deportivas", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Circuito cerrado de TV", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "En conjunto cerrado", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Garaje(s)", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Gimnasio", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Jardín", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Jaula de Golf", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Oficina de negocios", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Parqueadero Visitantes", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Piscina", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Sala de internet", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Salón Comunal", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Terraza", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Vigilancia", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Vista panorámica", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Vivienda Bifamiliar", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Vivienda Multifamiliar", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Zonas Verdes", FeatureTypeID: exteriorType.ID},
			domain.Feature{Name: "Colegios / Universidades", FeatureTypeID: sectorType.ID},
			domain.Feature{Name: "Parques cercanos", FeatureTypeID: sectorType.ID},
			domain.Feature{Name: "Sobre vía principal", FeatureTypeID: sectorType.ID},
			domain.Feature{Name: "Supermercados", FeatureTypeID: sectorType.ID},
			domain.Feature{Name: "C.Comerciales", FeatureTypeID: sectorType.ID},
			domain.Feature{Name: "Trans. Público cercano", FeatureTypeID: sectorType.ID},
			domain.Feature{Name: "Zona Campestre", FeatureTypeID: sectorType.ID},
			domain.Feature{Name: "Zona Comercial", FeatureTypeID: sectorType.ID},
			domain.Feature{Name: "Zona Residencial", FeatureTypeID: sectorType.ID},
		}

		_, err = collection.InsertMany(context.Background(), features)
		if err != nil {
			log.Fatalf("Error al insertar características: %v", err)
		}

		log.Println("Features seeding completado con éxito.")
	}
}
