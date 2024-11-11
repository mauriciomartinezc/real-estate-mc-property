package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	configCommon "github.com/mauriciomartinezc/real-estate-mc-common/config"
	"github.com/mauriciomartinezc/real-estate-mc-common/middleware"
	"github.com/mauriciomartinezc/real-estate-mc-property/config"
	"github.com/mauriciomartinezc/real-estate-mc-property/handlers"
	"github.com/mauriciomartinezc/real-estate-mc-property/routes"
	"github.com/mauriciomartinezc/real-estate-mc-property/seeds"
	"github.com/mauriciomartinezc/real-estate-mc-property/utils"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("application failed: %v", err)
	}
}

func run() error {
	if err := configCommon.LoadEnv(); err != nil {
		return fmt.Errorf("failed to load environment: %w", err)
	}

	if err := config.ValidateEnvironments(); err != nil {
		return fmt.Errorf("invalid environment configuration: %w", err)
	}

	uriMongoDB, databaseName := config.GetUriMongoDB()

	db, err := utils.ConnectMongoDB(uriMongoDB, databaseName)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	cacheClient := config.NewCacheClient()

	seeds.Run(db)

	e := echo.New()
	e.Use(middleware.LanguageHandler())
	handlers.InitValidate()
	routes.SetupRoutes(e, db, cacheClient)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e.Start(":" + os.Getenv("SERVER_PORT"))
}
