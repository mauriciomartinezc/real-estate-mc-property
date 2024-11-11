package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

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

// GetAll retrieves all property types, using cache if available
func (r *PropertyTypeRepository) GetAll() (domain.PropertyTypes, error) {
	cacheKey := "property_types:all"

	// Attempt to retrieve data from cache
	var propertyTypes domain.PropertyTypes
	if err := r.Cache.Get(cacheKey, &propertyTypes); err == nil {
		return propertyTypes, nil
	}

	// If not in cache, query the database
	findOptions := options.Find().SetSort(bson.D{{"name", 1}})
	cursor, err := r.Collection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(context.Background()); err != nil {
			log.Printf("Error closing cursor: %v", err)
		}
	}()

	if err := cursor.All(context.Background(), &propertyTypes); err != nil {
		return nil, err
	}

	// Store result in cache
	if err := r.Cache.Set(cacheKey, propertyTypes, 5*time.Minute); err != nil {
		log.Printf("Error caching property types: %v", err)
	}

	return propertyTypes, nil
}

// Create inserts a new property type and clears related cache entries
func (r *PropertyTypeRepository) Create(propertyType *domain.PropertyType) error {
	propertyType.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), propertyType)
	if err == nil {
		// Clear cache to ensure new data is reflected
		if err := r.Cache.Delete("property_types:all"); err != nil {
			log.Printf("Error deleting 'property_types:all' from cache: %v", err)
		}
	}
	return err
}

// GetByID retrieves a property type by ID, using cache if available
func (r *PropertyTypeRepository) GetByID(id primitive.ObjectID) (*domain.PropertyType, error) {
	cacheKey := fmt.Sprintf("property_type:%s", id.Hex())

	// Attempt to retrieve data from cache
	var propertyType domain.PropertyType
	if err := r.Cache.Get(cacheKey, &propertyType); err == nil {
		return &propertyType, nil
	}

	// If not in cache, query the database
	err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&propertyType)
	if err != nil {
		return nil, err
	}

	// Store result in cache
	if err := r.Cache.Set(cacheKey, propertyType, 5*time.Minute); err != nil {
		log.Printf("Error caching property type by ID: %v", err)
	}

	return &propertyType, nil
}

// Update modifies an existing property type and clears related cache entries
func (r *PropertyTypeRepository) Update(propertyType *domain.PropertyType) error {
	_, err := r.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": propertyType.ID},
		bson.M{"$set": propertyType},
	)
	if err == nil {
		// Clear both the individual and list cache entries to ensure consistency
		if err := r.Cache.Delete("property_types:all"); err != nil {
			log.Printf("Error deleting 'property_types:all' from cache: %v", err)
		}
		if err := r.Cache.Delete(fmt.Sprintf("property_type:%s", propertyType.ID.Hex())); err != nil {
			log.Printf("Error deleting property_type:%s from cache: %v", propertyType.ID.Hex(), err)
		}
	}
	return err
}

// Delete removes a property type by ID and clears related cache entries
func (r *PropertyTypeRepository) Delete(id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err == nil {
		// Clear both the individual and list cache entries to ensure consistency
		if err := r.Cache.Delete("property_types:all"); err != nil {
			log.Printf("Error deleting 'property_types:all' from cache: %v", err)
		}
		if err := r.Cache.Delete(fmt.Sprintf("property_type:%s", id.Hex())); err != nil {
			log.Printf("Error deleting property_type:%s from cache: %v", id.Hex(), err)
		}
	}
	return err
}
