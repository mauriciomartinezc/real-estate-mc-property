package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-property/handler"
	"github.com/mauriciomartinezc/real-estate-mc-property/repository"
	"github.com/mauriciomartinezc/real-estate-mc-property/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(e *echo.Echo, db *mongo.Database) {
	g := e.Group("api")
	managementType(g, db)
	age(g, db)
	featureType(g, db)
	feature(g, db)
	propertyType(g, db)
}

func managementType(g *echo.Group, db *mongo.Database) {
	repo := repository.NewManagementTypeRepository(db)
	service := services.NewManagementTypeService(repo)
	managementTypeHandler := handler.NewManagementTypeHandler(service)

	g.GET("/managementTypes", managementTypeHandler.GetManagementTypes)
	//g.POST("/managementTypes", managementTypeHandler.CreateManagementType)
	//g.GET("/managementTypes/:id", managementTypeHandler.GetManagementType)
	//g.PUT("/managementTypes/:id", managementTypeHandler.UpdateManagementType)
	//g.DELETE("/managementTypes/:id", managementTypeHandler.DeleteManagementType)
}

func age(g *echo.Group, db *mongo.Database) {
	repo := repository.NewAgeRepository(db)
	service := services.NewAgeService(repo)
	ageHandler := handler.NewAgeHandler(service)

	g.GET("/ages", ageHandler.GetAges)
	//g.POST("/ages", ageHandler.CreateAge)
	//g.GET("/ages/:id", ageHandler.GetAge)
	//g.PUT("/ages/:id", ageHandler.UpdateAge)
	//g.DELETE("/ages/:id", ageHandler.DeleteAge)
}

func featureType(g *echo.Group, db *mongo.Database) {
	repo := repository.NewFeatureTypeRepository(db)
	service := services.NewFeatureTypeService(repo)
	featureTypeHandler := handler.NewFeatureTypeHandler(service)

	g.GET("/featureTypes", featureTypeHandler.GetFeatureTypes)
	//g.POST("/featureTypes", featureTypeHandler.CreateFeatureType)
	//g.GET("/featureTypes/:id", featureTypeHandler.GetAFeatureType)
	//g.PUT("/featureTypes/:id", featureTypeHandler.UpdateFeatureType)
	//g.DELETE("/featureTypes/:id", featureTypeHandler.DeleteFeatureType)
}

func feature(g *echo.Group, db *mongo.Database) {
	repo := repository.NewFeatureRepository(db)
	service := services.NewFeatureService(repo)
	featureHandler := handler.NewFeatureHandler(service)

	g.GET("/features", featureHandler.GetFeatures)
	g.GET("/features/grouped", featureHandler.GetFeaturesGroupedByType)
	//g.POST("/features", featureHandler.CreateFeature)
	//g.GET("/features/:id", featureHandler.GetAFeature)
	//g.PUT("/features/:id", featureHandler.UpdateFeature)
	//g.DELETE("/features/:id", featureHandler.DeleteFeature)
}

func propertyType(g *echo.Group, db *mongo.Database) {
	repo := repository.NewPropertyTypeRepository(db)
	service := services.NewPropertyTypeService(repo)
	propertyTypeHandler := handler.NewPropertyTypeHandler(service)

	g.GET("/propertyTypes", propertyTypeHandler.GetPropertyTypes)
	//g.POST("/propertyTypes", propertyTypeHandler.CreatePropertyType)
	//g.GET("/propertyTypes/:id", propertyTypeHandler.GetPropertyType)
	//g.PUT("/propertyTypes/:id", propertyTypeHandler.UpdatePropertyType)
	//g.DELETE("/propertyTypes/:id", propertyTypeHandler.DeletePropertyType)
}
