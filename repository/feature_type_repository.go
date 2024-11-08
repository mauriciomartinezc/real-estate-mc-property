package repository

import (
	"context"
	"github.com/mauriciomartinezc/real-estate-mc-property/cache"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FeatureTypeRepository struct {
	Collection *mongo.Collection
	Cache      cache.Cache
}

func NewFeatureTypeRepository(db *mongo.Database, cache cache.Cache) *FeatureTypeRepository {
	return &FeatureTypeRepository{
		Collection: db.Collection("feature_types"),
		Cache:      cache,
	}
}

func (r *FeatureTypeRepository) GetAll() (domain.FeatureTypes, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"name", 1}})

	cursor, err := r.Collection.Find(context.Background(), bson.M{}, findOptions)

	if err != nil {
		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

	var featureTypes []domain.FeatureType
	if err = cursor.All(context.Background(), &featureTypes); err != nil {
		return nil, err
	}

	return featureTypes, nil
}

func (r *FeatureTypeRepository) Create(featureType *domain.FeatureType) error {
	featureType.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), featureType)
	return err
}

func (r *FeatureTypeRepository) GetByID(id primitive.ObjectID) (*domain.FeatureType, error) {
	var featureType domain.FeatureType
	err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&featureType)
	return &featureType, err
}

func (r *FeatureTypeRepository) Update(featureType *domain.FeatureType) error {
	_, err := r.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": featureType.ID},
		bson.M{"$set": featureType},
	)
	return err
}

func (r *FeatureTypeRepository) Delete(id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
