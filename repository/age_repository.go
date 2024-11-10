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

// GetAll retrieves all age records, using cache if available
func (r *AgeRepository) GetAll() (domain.Ages, error) {
	cacheKey := "ages:all"

	// Attempt to retrieve data from cache
	var ages domain.Ages
	if err := r.Cache.Get(cacheKey, &ages); err == nil {
		// Check if ages is actually populated
		if len(ages) > 0 {
			return ages, nil
		}
	}

	// If not in cache or cache is empty, query database
	findOptions := options.Find().SetSort(bson.D{{"order", 1}})
	cursor, err := r.Collection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(context.Background()); err != nil {
			log.Printf("Error closing cursor: %v", err)
		}
	}()

	if err := cursor.All(context.Background(), &ages); err != nil {
		return nil, err
	}

	// Check if ages has data before saving to cache
	if len(ages) > 0 {
		// Store result in cache for future requests
		if err := r.Cache.Set(cacheKey, ages, 5*time.Minute); err != nil {
			log.Printf("Error saving to cache: %v", err)
		}
	}

	return ages, nil
}

// Create inserts a new age record and clears related cache entries
func (r *AgeRepository) Create(age *domain.Age) error {
	age.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), age)
	if err == nil {
		// Clear cache to ensure newly created record is included in subsequent GetAll queries
		if err := r.Cache.Delete("ages:all"); err != nil {
			log.Printf("Error deleting ages:all from cache: %v", err)
		}
	}
	return err
}

// GetByID retrieves an age record by ID, using cache if available
func (r *AgeRepository) GetByID(id primitive.ObjectID) (*domain.Age, error) {
	cacheKey := fmt.Sprintf("age:%s", id.Hex())

	// Attempt to retrieve data from cache
	var age domain.Age
	if err := r.Cache.Get(cacheKey, &age); err == nil {
		return &age, nil
	}

	// If not in cache, query database
	err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&age)
	if err != nil {
		return nil, err
	}

	// Store result in cache for future requests
	if err := r.Cache.Set(cacheKey, age, 5*time.Minute); err != nil {
		log.Printf("Error saving to cache: %v", err)
	}

	return &age, nil
}

// Update modifies an existing age record and clears related cache entries
func (r *AgeRepository) Update(age *domain.Age) error {
	_, err := r.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": age.ID},
		bson.M{"$set": age},
	)
	if err == nil {
		// Clear both the individual and list cache entries to ensure consistency
		if err := r.Cache.Delete("ages:all"); err != nil {
			log.Printf("Error deleting ages:all from cache: %v", err)
		}
		if err := r.Cache.Delete(fmt.Sprintf("age:%s", age.ID.Hex())); err != nil {
			log.Printf("Error deleting age:%s from cache: %v", age.ID.Hex(), err)
		}
	}
	return err
}

// Delete removes an age record by ID and clears related cache entries
func (r *AgeRepository) Delete(id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err == nil {
		// Clear both the individual and list cache entries to ensure consistency
		if err := r.Cache.Delete("ages:all"); err != nil {
			log.Printf("Error deleting ages:all from cache: %v", err)
		}
		if err := r.Cache.Delete(fmt.Sprintf("age:%s", id.Hex())); err != nil {
			log.Printf("Error deleting age:%s from cache: %v", id.Hex(), err)
		}
	}
	return err
}
