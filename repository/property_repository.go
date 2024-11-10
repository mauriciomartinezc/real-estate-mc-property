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

// GetAllPropertiesPaginated retrieves properties with pagination, using cache if available
func (r *PropertyRepository) GetAllPropertiesPaginated(page int, limit int) (domain.SimpleProperties, error) {
	cacheKey := fmt.Sprintf("properties:page:%d:limit:%d", page, limit)

	// Attempt to retrieve data from cache
	var properties domain.SimpleProperties
	if err := r.Cache.Get(cacheKey, &properties); err == nil {
		return properties, nil
	}

	// If not in cache, query the database
	skip := (page - 1) * limit
	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))

	cursor, err := r.PropertyCollection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(context.Background()); err != nil {
			log.Printf("Error closing cursor for paginated properties: %v", err)
		}
	}()

	if err := cursor.All(context.Background(), &properties); err != nil {
		return nil, err
	}

	// Store result in cache
	if err := r.Cache.Set(cacheKey, properties, 5*time.Minute); err != nil {
		log.Printf("Error caching paginated properties: %v", err)
	}

	return properties, nil
}

// GetPropertiesByCompanyID retrieves properties by company ID with pagination, using cache if available
func (r *PropertyRepository) GetPropertiesByCompanyID(companyID string, page, limit int) (domain.SimpleProperties, error) {
	cacheKey := fmt.Sprintf("properties:company:%s:page:%d:limit:%d", companyID, page, limit)

	// Attempt to retrieve data from cache
	var properties domain.SimpleProperties
	if err := r.Cache.Get(cacheKey, &properties); err == nil {
		return properties, nil
	}

	// If not in cache, query the database
	skip := (page - 1) * limit
	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))
	filter := bson.M{"company_id": companyID}

	cursor, err := r.PropertyCollection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(context.Background()); err != nil {
			log.Printf("Error closing cursor for company properties: %v", err)
		}
	}()

	if err := cursor.All(context.Background(), &properties); err != nil {
		return nil, err
	}

	// Store result in cache
	if err := r.Cache.Set(cacheKey, properties, 5*time.Minute); err != nil {
		log.Printf("Error caching company properties: %v", err)
	}

	return properties, nil
}

// Create inserts a new property and updates the cache for relevant entries
func (r *PropertyRepository) Create(property *domain.SimpleProperty) error {
	property.ID = primitive.NewObjectID()
	property.CreatedAt = time.Now().Unix()
	property.UpdatedAt = time.Now().Unix()

	_, err := r.PropertyCollection.InsertOne(context.Background(), property)
	if err == nil {
		// Update cache with the new property entry
		r.updateCacheAfterChange(fmt.Sprintf("property:%s", property.ID.Hex()), property)
	}
	return err
}

// Update modifies an existing property and updates the cache for relevant entries
func (r *PropertyRepository) Update(property *domain.SimpleProperty) error {
	property.UpdatedAt = time.Now().Unix()

	_, err := r.PropertyCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": property.ID},
		bson.M{"$set": property},
	)
	if err == nil {
		// Update cache with the updated property entry
		r.updateCacheAfterChange(fmt.Sprintf("property:%s", property.ID.Hex()), property)
	}
	return err
}

// ChangeStatus toggles the active status of a property and updates the cache
func (r *PropertyRepository) ChangeStatus(property *domain.SimpleProperty) error {
	property.UpdatedAt = time.Now().Unix()
	property.Active = !property.Active

	_, err := r.PropertyCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": property.ID},
		bson.M{"$set": property},
	)
	if err == nil {
		// Update cache with the updated property status
		r.updateCacheAfterChange(fmt.Sprintf("property:%s", property.ID.Hex()), property)
	}
	return err
}

// GetDetailByID retrieves detailed property information by ID, using cache if available
func (r *PropertyRepository) GetDetailByID(id primitive.ObjectID) (*domain.DetailProperty, error) {
	cacheKey := fmt.Sprintf("property_detail:%s", id.Hex())

	var property domain.DetailProperty
	if err := r.Cache.Get(cacheKey, &property); err == nil {
		return &property, nil
	}

	err := r.PropertyCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&property)
	if err != nil {
		return nil, err
	}

	if err := r.Cache.Set(cacheKey, property, 5*time.Minute); err != nil {
		log.Printf("Error caching property detail by ID: %v", err)
	}

	return &property, nil
}

// GetByID retrieves a property by ID, using cache if available
func (r *PropertyRepository) GetByID(id primitive.ObjectID) (*domain.SimpleProperty, error) {
	cacheKey := fmt.Sprintf("property:%s", id.Hex())

	var property domain.SimpleProperty
	if err := r.Cache.Get(cacheKey, &property); err == nil {
		return &property, nil
	}

	err := r.PropertyCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&property)
	if err != nil {
		return nil, err
	}

	if err := r.Cache.Set(cacheKey, property, 5*time.Minute); err != nil {
		log.Printf("Error caching property by ID: %v", err)
	}

	return &property, nil
}

// updateCacheAfterChange updates the cache for specific keys after create, update, or status change
func (r *PropertyRepository) updateCacheAfterChange(cacheKey string, property *domain.SimpleProperty) {
	if err := r.Cache.Set(cacheKey, property, 5*time.Minute); err != nil {
		log.Printf("Error updating cache for key %s: %v", cacheKey, err)
	}
}
