package repository

import (
	"context"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type FeatureRepository struct {
	FeatureCollection     *mongo.Collection
	FeatureTypeCollection *mongo.Collection
}

func NewFeatureRepository(db *mongo.Database) *FeatureRepository {
	return &FeatureRepository{
		FeatureCollection:     db.Collection("features"),
		FeatureTypeCollection: db.Collection("feature_types"),
	}
}

func (r *FeatureRepository) GetAll() (domain.Features, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"name", 1}})

	cursor, err := r.FeatureCollection.Find(context.Background(), bson.M{}, findOptions)

	if err != nil {
		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

	var features []domain.Feature
	if err = cursor.All(context.Background(), &features); err != nil {
		return nil, err
	}

	return features, nil
}

func (r *FeatureRepository) GetFeaturesGroupedByType() (map[string][]domain.Feature, error) {
	var featureTypes []domain.FeatureType
	cursor, err := r.FeatureTypeCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &featureTypes); err != nil {
		return nil, err
	}

	featuresGroupedByType := make(map[string][]domain.Feature)

	for _, featureType := range featureTypes {
		var features []domain.Feature
		cursor, err := r.FeatureCollection.Find(context.Background(), bson.M{"feature_type_id": featureType.ID})
		if err != nil {
			log.Printf("Error al buscar features para el tipo %s: %v", featureType.Name, err)
			continue
		}
		defer cursor.Close(context.Background())

		if err := cursor.All(context.Background(), &features); err != nil {
			log.Printf("Error al decodificar features para el tipo %s: %v", featureType.Name, err)
			continue
		}

		featuresGroupedByType[featureType.Name] = features
	}

	return featuresGroupedByType, nil
}

func (r *FeatureRepository) Create(feature *domain.Feature) error {
	feature.ID = primitive.NewObjectID()
	_, err := r.FeatureCollection.InsertOne(context.Background(), feature)
	return err
}

func (r *FeatureRepository) GetByID(id primitive.ObjectID) (*domain.Feature, error) {
	var feature domain.Feature
	err := r.FeatureCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&feature)
	return &feature, err
}

func (r *FeatureRepository) Update(feature *domain.Feature) error {
	_, err := r.FeatureCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": feature.ID},
		bson.M{"$set": feature},
	)
	return err
}

func (r *FeatureRepository) Delete(id primitive.ObjectID) error {
	_, err := r.FeatureCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
