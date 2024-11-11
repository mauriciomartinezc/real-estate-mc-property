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

type FeatureRepository struct {
	FeatureCollection     *mongo.Collection
	FeatureTypeCollection *mongo.Collection
	Cache                 cache.Cache
}

func NewFeatureRepository(db *mongo.Database, cache cache.Cache) *FeatureRepository {
	return &FeatureRepository{
		FeatureCollection:     db.Collection("features"),
		FeatureTypeCollection: db.Collection("feature_types"),
		Cache:                 cache,
	}
}

// GetAll retrieves all features, using cache if available
func (r *FeatureRepository) GetAll() (domain.Features, error) {
	cacheKey := "features:all"

	// Attempt to retrieve data from cache
	var features domain.Features
	if err := r.Cache.Get(cacheKey, &features); err == nil {
		return features, nil
	}

	// If not in cache, query the database
	findOptions := options.Find().SetSort(bson.D{{"name", 1}})
	cursor, err := r.FeatureCollection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(context.Background()); err != nil {
			log.Printf("Error closing cursor: %v", err)
		}
	}()

	if err := cursor.All(context.Background(), &features); err != nil {
		return nil, err
	}

	// Store result in cache
	if err := r.Cache.Set(cacheKey, features, 5*time.Minute); err != nil {
		log.Printf("Error caching features: %v", err)
	}

	return features, nil
}

// GetFeaturesGroupedByType retrieves features grouped by feature type, using cache if available
func (r *FeatureRepository) GetFeaturesGroupedByType() (map[string][]domain.Feature, error) {
	cacheKey := "features:grouped_by_type"

	// Attempt to retrieve data from cache
	var featuresGroupedByType map[string][]domain.Feature
	if err := r.Cache.Get(cacheKey, &featuresGroupedByType); err == nil {
		return featuresGroupedByType, nil
	}

	// If not in cache, query the database
	var featureTypes []domain.FeatureType
	cursor, err := r.FeatureTypeCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(context.Background()); err != nil {
			log.Printf("Error closing cursor for feature types: %v", err)
		}
	}()

	if err := cursor.All(context.Background(), &featureTypes); err != nil {
		return nil, err
	}

	featuresGroupedByType = make(map[string][]domain.Feature)
	for _, featureType := range featureTypes {
		var features []domain.Feature
		cursor, err := r.FeatureCollection.Find(context.Background(), bson.M{"feature_type_id": featureType.ID})
		if err != nil {
			log.Printf("Error retrieving features for type %s: %v", featureType.Name, err)
			continue
		}
		defer func() {
			if err := cursor.Close(context.Background()); err != nil {
				log.Printf("Error closing cursor for features: %v", err)
			}
		}()

		if err := cursor.All(context.Background(), &features); err != nil {
			log.Printf("Error decoding features for type %s: %v", featureType.Name, err)
			continue
		}

		featuresGroupedByType[featureType.Name] = features
	}

	// Store result in cache
	if err := r.Cache.Set(cacheKey, featuresGroupedByType, 5*time.Minute); err != nil {
		log.Printf("Error caching grouped features: %v", err)
	}

	return featuresGroupedByType, nil
}

// Create inserts a new feature and clears relevant cache entries
func (r *FeatureRepository) Create(feature *domain.Feature) error {
	feature.ID = primitive.NewObjectID()
	_, err := r.FeatureCollection.InsertOne(context.Background(), feature)
	if err == nil {
		// Clear cache entries to ensure new data is reflected
		if err := r.Cache.Delete("features:all"); err != nil {
			log.Printf("Error deleting 'features:all' from cache: %v", err)
		}
		if err := r.Cache.Delete("features:grouped_by_type"); err != nil {
			log.Printf("Error deleting 'features:grouped_by_type' from cache: %v", err)
		}
	}
	return err
}

// GetByID retrieves a feature by ID, using cache if available
func (r *FeatureRepository) GetByID(id primitive.ObjectID) (*domain.Feature, error) {
	cacheKey := fmt.Sprintf("feature:%s", id.Hex())

	// Attempt to retrieve data from cache
	var feature domain.Feature
	if err := r.Cache.Get(cacheKey, &feature); err == nil {
		return &feature, nil
	}

	// If not in cache, query the database
	err := r.FeatureCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&feature)
	if err != nil {
		return nil, err
	}

	// Store result in cache
	if err := r.Cache.Set(cacheKey, feature, 5*time.Minute); err != nil {
		log.Printf("Error caching feature by ID: %v", err)
	}

	return &feature, nil
}

// Update modifies an existing feature and clears relevant cache entries
func (r *FeatureRepository) Update(feature *domain.Feature) error {
	_, err := r.FeatureCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": feature.ID},
		bson.M{"$set": feature},
	)
	if err == nil {
		// Clear cache entries to ensure consistency
		if err := r.Cache.Delete("features:all"); err != nil {
			log.Printf("Error deleting 'features:all' from cache: %v", err)
		}
		if err := r.Cache.Delete("features:grouped_by_type"); err != nil {
			log.Printf("Error deleting 'features:grouped_by_type' from cache: %v", err)
		}
		if err := r.Cache.Delete(fmt.Sprintf("feature:%s", feature.ID.Hex())); err != nil {
			log.Printf("Error deleting feature:%s from cache: %v", feature.ID.Hex(), err)
		}
	}
	return err
}

// Delete removes a feature by ID and clears relevant cache entries
func (r *FeatureRepository) Delete(id primitive.ObjectID) error {
	_, err := r.FeatureCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err == nil {
		// Clear cache entries to ensure consistency
		if err := r.Cache.Delete("features:all"); err != nil {
			log.Printf("Error deleting 'features:all' from cache: %v", err)
		}
		if err := r.Cache.Delete("features:grouped_by_type"); err != nil {
			log.Printf("Error deleting 'features:grouped_by_type' from cache: %v", err)
		}
		if err := r.Cache.Delete(fmt.Sprintf("feature:%s", id.Hex())); err != nil {
			log.Printf("Error deleting feature:%s from cache: %v", id.Hex(), err)
		}
	}
	return err
}
