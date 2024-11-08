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

type AgeRepository struct {
	Collection *mongo.Collection
	Cache      cache.Cache
}

func NewAgeRepository(db *mongo.Database, cache cache.Cache) *AgeRepository {
	return &AgeRepository{
		Collection: db.Collection("ages"),
		Cache:      cache,
	}
}

func (r *AgeRepository) GetAll() (domain.Ages, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"order", 1}})

	cursor, err := r.Collection.Find(context.Background(), bson.M{}, findOptions)

	if err != nil {
		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

	var ages []domain.Age
	if err = cursor.All(context.Background(), &ages); err != nil {
		return nil, err
	}

	return ages, nil
}

func (r *AgeRepository) Create(age *domain.Age) error {
	age.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), age)
	return err
}

func (r *AgeRepository) GetByID(id primitive.ObjectID) (*domain.Age, error) {
	var age domain.Age
	err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&age)
	return &age, err
}

func (r *AgeRepository) Update(age *domain.Age) error {
	_, err := r.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": age.ID},
		bson.M{"$set": age},
	)
	return err
}

func (r *AgeRepository) Delete(id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
