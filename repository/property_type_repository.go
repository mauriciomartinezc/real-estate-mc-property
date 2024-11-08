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

type PropertyTypeRepository struct {
	Collection *mongo.Collection
	Cache      cache.Cache
}

func NewPropertyTypeRepository(db *mongo.Database, cache cache.Cache) *PropertyTypeRepository {
	return &PropertyTypeRepository{
		Collection: db.Collection("property_types"),
		Cache:      cache,
	}
}

func (r *PropertyTypeRepository) GetAll() (domain.PropertyTypes, error) {
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

	var propertyTypes []domain.PropertyType
	if err = cursor.All(context.Background(), &propertyTypes); err != nil {
		return nil, err
	}

	return propertyTypes, nil
}

func (r *PropertyTypeRepository) Create(propertyType *domain.PropertyType) error {
	propertyType.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), propertyType)
	return err
}

func (r *PropertyTypeRepository) GetByID(id primitive.ObjectID) (*domain.PropertyType, error) {
	var propertyType domain.PropertyType
	err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&propertyType)
	return &propertyType, err
}

func (r *PropertyTypeRepository) Update(propertyType *domain.PropertyType) error {
	_, err := r.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": propertyType.ID},
		bson.M{"$set": propertyType},
	)
	return err
}

func (r *PropertyTypeRepository) Delete(id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
