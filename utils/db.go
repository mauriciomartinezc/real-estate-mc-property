package utils

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(uri string, databaseName string) (*mongo.Database, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	if err = client.Database(databaseName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		err = client.Disconnect(context.TODO())
		if err != nil {
			return nil, fmt.Errorf("error in Disconect to MongoDB: %v", err)
		}
		return nil, fmt.Errorf("error in RunCommand to MongoDB: %v", err)
	}

	return client.Database(databaseName), nil
}
