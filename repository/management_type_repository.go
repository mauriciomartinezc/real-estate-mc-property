package repository

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

// GetAll retrieves all management types, using cache if available
func (r *ManagementTypeRepository) GetAll() (domain.ManagementTypes, error) {
	cacheKey := "management_types:all"

	// Attempt to retrieve data from cache
	var managementTypes domain.ManagementTypes
	if err := r.Cache.Get(cacheKey, &managementTypes); err == nil {
		return managementTypes, nil
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

	if err := cursor.All(context.Background(), &managementTypes); err != nil {
		return nil, err
	}

	// Store result in cache
	if err := r.Cache.Set(cacheKey, managementTypes, 5*time.Minute); err != nil {
		log.Printf("Error caching management types: %v", err)
	}

	return managementTypes, nil
}

// Create inserts a new management type and clears related cache entries
func (r *ManagementTypeRepository) Create(managementType *domain.ManagementType) error {
	managementType.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), managementType)
	if err == nil {
		// Clear cache to ensure new data is reflected
		if err := r.Cache.Delete("management_types:all"); err != nil {
			log.Printf("Error deleting 'management_types:all' from cache: %v", err)
		}
	}
	return err
}

// GetByID retrieves a management type by ID, using cache if available
func (r *ManagementTypeRepository) GetByID(id primitive.ObjectID) (*domain.ManagementType, error) {
	cacheKey := fmt.Sprintf("management_type:%s", id.Hex())

	// Attempt to retrieve data from cache
	var managementType domain.ManagementType
	if err := r.Cache.Get(cacheKey, &managementType); err == nil {
		return &managementType, nil
	}

	// If not in cache, query the database
	err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&managementType)
	if err != nil {
		return nil, err
	}

	// Store result in cache
	if err := r.Cache.Set(cacheKey, managementType, 5*time.Minute); err != nil {
		log.Printf("Error caching management type by ID: %v", err)
	}

	return &managementType, nil
}

// Update modifies an existing management type and clears related cache entries
func (r *ManagementTypeRepository) Update(managementType *domain.ManagementType) error {
	_, err := r.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": managementType.ID},
		bson.M{"$set": managementType},
	)
	if err == nil {
		// Clear both the individual and list cache entries to ensure consistency
		if err := r.Cache.Delete("management_types:all"); err != nil {
			log.Printf("Error deleting 'management_types:all' from cache: %v", err)
		}
		if err := r.Cache.Delete(fmt.Sprintf("management_type:%s", managementType.ID.Hex())); err != nil {
			log.Printf("Error deleting management_type:%s from cache: %v", managementType.ID.Hex(), err)
		}
	}
	return err
}

// Delete removes a management type by ID and clears related cache entries
func (r *ManagementTypeRepository) Delete(id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err == nil {
		// Clear both the individual and list cache entries to ensure consistency
		if err := r.Cache.Delete("management_types:all"); err != nil {
			log.Printf("Error deleting 'management_types:all' from cache: %v", err)
		}
		if err := r.Cache.Delete(fmt.Sprintf("management_type:%s", id.Hex())); err != nil {
			log.Printf("Error deleting management_type:%s from cache: %v", id.Hex(), err)
		}
	}
	return err
}
