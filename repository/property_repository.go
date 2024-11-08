package repository

import (
	"context"
	"github.com/mauriciomartinezc/real-estate-mc-property/cache"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type PropertyRepository struct {
	PropertyCollection *mongo.Collection
	Cache              cache.Cache
}

func NewPropertyRepository(db *mongo.Database, cache cache.Cache) *PropertyRepository {
	return &PropertyRepository{
		PropertyCollection: db.Collection("properties"),
		Cache:              cache,
	}
}

func (r *PropertyRepository) GetAllPropertiesPaginated(page int, limit int) (domain.SimpleProperties, error) {
	skip := (page - 1) * limit
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	cursor, err := r.PropertyCollection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

	var properties domain.SimpleProperties
	if err := cursor.All(context.Background(), &properties); err != nil {
		return nil, err
	}

	return properties, nil
}

func (r *PropertyRepository) GetPropertiesByCompanyID(companyID string, page, limit int) (domain.SimpleProperties, error) {
	skip := (page - 1) * limit
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	filter := bson.M{"company_id": companyID}
	cursor, err := r.PropertyCollection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

	var properties domain.SimpleProperties
	if err := cursor.All(context.Background(), &properties); err != nil {
		return nil, err
	}

	return properties, nil
}

func (r *PropertyRepository) Create(property *domain.SimpleProperty) error {
	property.ID = primitive.NewObjectID()
	property.CreatedAt = time.Now().Unix()
	property.UpdatedAt = time.Now().Unix()
	_, err := r.PropertyCollection.InsertOne(context.Background(), property)
	return err
}

func (r *PropertyRepository) Update(property *domain.SimpleProperty) error {
	property.UpdatedAt = time.Now().Unix()
	_, err := r.PropertyCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": property.ID},
		bson.M{"$set": property},
	)
	return err
}

func (r *PropertyRepository) GetDetailByID(id primitive.ObjectID) (*domain.DetailProperty, error) {
	var property domain.DetailProperty
	err := r.PropertyCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&property)
	return &property, err
}

func (r *PropertyRepository) GetByID(id primitive.ObjectID) (*domain.SimpleProperty, error) {
	var property domain.SimpleProperty
	err := r.PropertyCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&property)
	return &property, err
}

func (r *PropertyRepository) ChangeStatus(property *domain.SimpleProperty) error {
	property.UpdatedAt = time.Now().Unix()
	property.Active = !property.Active
	_, err := r.PropertyCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": property.ID},
		bson.M{"$set": property},
	)

	return err
}
