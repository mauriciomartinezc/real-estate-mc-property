package repositories

import (
	"context"
	"fmt"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
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

// GetAll retrieves all feature types, using cache if available
func (r *FeatureTypeRepository) GetAll() (domain.FeatureTypes, error) {
	cacheKey := "feature_types:all"

	// Attempt to retrieve data from cache
	var featureTypes domain.FeatureTypes
	if err := r.Cache.Get(cacheKey, &featureTypes); err == nil {
		return featureTypes, nil
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

	if err := cursor.All(context.Background(), &featureTypes); err != nil {
		return nil, err
	}

	// Store result in cache
	if err := r.Cache.Set(cacheKey, featureTypes, 5*time.Minute); err != nil {
		log.Printf("Error caching feature types: %v", err)
	}

	return featureTypes, nil
}

// Create inserts a new feature type and clears related cache entries
func (r *FeatureTypeRepository) Create(featureType *domain.FeatureType) error {
	featureType.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), featureType)
	if err == nil {
		// Clear cache to ensure new data is reflected
		if err := r.Cache.Delete("feature_types:all"); err != nil {
			log.Printf("Error deleting 'feature_types:all' from cache: %v", err)
		}
	}
	return err
}

// GetByID retrieves a feature type by ID, using cache if available
func (r *FeatureTypeRepository) GetByID(id primitive.ObjectID) (*domain.FeatureType, error) {
	cacheKey := fmt.Sprintf("feature_type:%s", id.Hex())

	// Attempt to retrieve data from cache
	var featureType domain.FeatureType
	if err := r.Cache.Get(cacheKey, &featureType); err == nil {
		return &featureType, nil
	}

	// If not in cache, query the database
	err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&featureType)
	if err != nil {
		return nil, err
	}

	// Store result in cache
	if err := r.Cache.Set(cacheKey, featureType, 5*time.Minute); err != nil {
		log.Printf("Error caching feature type by ID: %v", err)
	}

	return &featureType, nil
}

// Update modifies an existing feature type and clears related cache entries
func (r *FeatureTypeRepository) Update(featureType *domain.FeatureType) error {
	_, err := r.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": featureType.ID},
		bson.M{"$set": featureType},
	)
	if err == nil {
		// Clear both the individual and list cache entries to ensure consistency
		if err := r.Cache.Delete("feature_types:all"); err != nil {
			log.Printf("Error deleting 'feature_types:all' from cache: %v", err)
		}
		if err := r.Cache.Delete(fmt.Sprintf("feature_type:%s", featureType.ID.Hex())); err != nil {
			log.Printf("Error deleting feature_type:%s from cache: %v", featureType.ID.Hex(), err)
		}
	}
	return err
}

// Delete removes a feature type by ID and clears related cache entries
func (r *FeatureTypeRepository) Delete(id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err == nil {
		// Clear both the individual and list cache entries to ensure consistency
		if err := r.Cache.Delete("feature_types:all"); err != nil {
			log.Printf("Error deleting 'feature_types:all' from cache: %v", err)
		}
		if err := r.Cache.Delete(fmt.Sprintf("feature_type:%s", id.Hex())); err != nil {
			log.Printf("Error deleting feature_type:%s from cache: %v", id.Hex(), err)
		}
	}
	return err
}
