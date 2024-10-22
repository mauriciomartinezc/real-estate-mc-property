package repository

import (
	"context"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ManagementTypeRepository struct {
	Collection *mongo.Collection
}

func NewManagementTypeRepository(db *mongo.Database) *ManagementTypeRepository {
	return &ManagementTypeRepository{
		Collection: db.Collection("management_types"),
	}
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
