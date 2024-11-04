package seeds

import "go.mongodb.org/mongo-driver/mongo"

func Run(db *mongo.Database) {
	ManagementTypeSeeder(db)
	AgeSeeder(db)
}
