package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	ServerPort string
	MongoURI   string
	MongoUser  string
	MongoPass  string
	Database   string
}

func ValidateEnvironments() error {
	requiredEnvs := []string{
		"SERVER_PORT",
		"ALLOWED_ORIGINS",
		"ALLOWED_METHODS",
		"MONGO_HOST",
		"MONGO_APP_NAME",
		//"MONGO_PORT",
		"MONGO_USERNAME",
		"MONGO_PASSWORD",
		"MONGO_DATABASE",
	}

	for _, env := range requiredEnvs {
		if getEnvironment(env) == "" {
			return getErrorSetEnvironment(env)
		}
	}

	return nil
}

func GetUriMongoDB() (string, string) {
	host := os.Getenv("MONGO_HOST")
	//port := os.Getenv("MONGO_PORT")
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	appName := os.Getenv("MONGO_APP_NAME")
	database := os.Getenv("MONGO_DATABASE")
	return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s", username, password, host, appName), database
}

func getEnvironment(environmentName string) string {
	return strings.TrimSpace(os.Getenv(environmentName))
}

func getErrorSetEnvironment(environmentName string) error {
	return fmt.Errorf("the environment variable %s is not set", environmentName)
}
