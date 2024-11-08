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

type ManagementTypeRepository struct {
	Collection *mongo.Collection
	Cache      cache.Cache
}

func NewManagementTypeRepository(db *mongo.Database, cache cache.Cache) *ManagementTypeRepository {
	return &ManagementTypeRepository{
		Collection: db.Collection("management_types"),
		Cache:      cache,
	}
}

func (r *ManagementTypeRepository) GetAll() (domain.ManagementTypes, error) {
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

	var managementTypes []domain.ManagementType
	if err = cursor.All(context.Background(), &managementTypes); err != nil {
		return nil, err
	}

	return managementTypes, nil
}

func (r *ManagementTypeRepository) Create(managementType *domain.ManagementType) error {
	managementType.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), managementType)
	return err
}

func (r *ManagementTypeRepository) GetByID(id primitive.ObjectID) (*domain.ManagementType, error) {
	var managementType domain.ManagementType
	err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&managementType)
	return &managementType, err
}

func (r *ManagementTypeRepository) Update(managementType *domain.ManagementType) error {
	_, err := r.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": managementType.ID},
		bson.M{"$set": managementType},
	)
	return err
}

func (r *ManagementTypeRepository) Delete(id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
