package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"github.com/mauriciomartinezc/real-estate-mc-common/middlewares"
	"github.com/mauriciomartinezc/real-estate-mc-property/handlers"
	"github.com/mauriciomartinezc/real-estate-mc-property/repositories"
	"github.com/mauriciomartinezc/real-estate-mc-property/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(e *echo.Echo, db *mongo.Database, cache cache.Cache) {
	g := e.Group("api")
	managementType(g, db, cache)
	age(g, db, cache)
	featureType(g, db, cache)
	feature(g, db, cache)
	propertyType(g, db, cache)
	property(g, db, cache)
}

func managementType(g *echo.Group, db *mongo.Database, cache cache.Cache) {
	repo := repositories.NewManagementTypeRepository(db, cache)
	service := services.NewManagementTypeService(repo)
	managementTypeHandler := handlers.NewManagementTypeHandler(service)

	g.GET("/managementTypes", managementTypeHandler.GetManagementTypes)
	//g.POST("/managementTypes", managementTypeHandler.CreateManagementType)
	//g.GET("/managementTypes/:id", managementTypeHandler.GetManagementType)
	//g.PUT("/managementTypes/:id", managementTypeHandler.UpdateManagementType)
	//g.DELETE("/managementTypes/:id", managementTypeHandler.DeleteManagementType)
}

func age(g *echo.Group, db *mongo.Database, cache cache.Cache) {
	repo := repositories.NewAgeRepository(db, cache)
	service := services.NewAgeService(repo)
	ageHandler := handlers.NewAgeHandler(service)

	g.GET("/ages", ageHandler.GetAges)
	//g.POST("/ages", ageHandler.CreateAge)
	//g.GET("/ages/:id", ageHandler.GetAge)
	//g.PUT("/ages/:id", ageHandler.UpdateAge)
	//g.DELETE("/ages/:id", ageHandler.DeleteAge)
}

func featureType(g *echo.Group, db *mongo.Database, cache cache.Cache) {
	repo := repositories.NewFeatureTypeRepository(db, cache)
	service := services.NewFeatureTypeService(repo)
	featureTypeHandler := handlers.NewFeatureTypeHandler(service)

	g.GET("/featureTypes", featureTypeHandler.GetFeatureTypes)
	//g.POST("/featureTypes", featureTypeHandler.CreateFeatureType)
	//g.GET("/featureTypes/:id", featureTypeHandler.GetAFeatureType)
	//g.PUT("/featureTypes/:id", featureTypeHandler.UpdateFeatureType)
	//g.DELETE("/featureTypes/:id", featureTypeHandler.DeleteFeatureType)
}

func feature(g *echo.Group, db *mongo.Database, cache cache.Cache) {
	repo := repositories.NewFeatureRepository(db, cache)
	service := services.NewFeatureService(repo)
	featureHandler := handlers.NewFeatureHandler(service)

	g.GET("/features", featureHandler.GetFeatures)
	g.GET("/features/grouped", featureHandler.GetFeaturesGroupedByType)
	//g.POST("/features", featureHandler.CreateFeature)
	//g.GET("/features/:id", featureHandler.GetAFeature)
	//g.PUT("/features/:id", featureHandler.UpdateFeature)
	//g.DELETE("/features/:id", featureHandler.DeleteFeature)
}

func propertyType(g *echo.Group, db *mongo.Database, cache cache.Cache) {
	repo := repositories.NewPropertyTypeRepository(db, cache)
	service := services.NewPropertyTypeService(repo)
	propertyTypeHandler := handlers.NewPropertyTypeHandler(service)

	g.GET("/propertyTypes", propertyTypeHandler.GetPropertyTypes)
	//g.POST("/propertyTypes", propertyTypeHandler.CreatePropertyType)
	//g.GET("/propertyTypes/:id", propertyTypeHandler.GetPropertyType)
	//g.PUT("/propertyTypes/:id", propertyTypeHandler.UpdatePropertyType)
	//g.DELETE("/propertyTypes/:id", propertyTypeHandler.DeletePropertyType)
}

func property(g *echo.Group, db *mongo.Database, cache cache.Cache) {
	repo := repositories.NewPropertyRepository(db, cache)
	service := services.NewPropertyService(repo)
	propertyHandler := handlers.NewPropertyHandler(service)

	g.GET("/properties/company", propertyHandler.GetPropertiesByCompanyID, middlewares.CompanyHandler())
	g.GET("/properties", propertyHandler.GetAllPropertiesPaginated)
	g.GET("/properties/:id", propertyHandler.GetDetailProperty)
	g.POST("/properties", propertyHandler.CreateProperty)
	g.PUT("/properties/:id", propertyHandler.UpdateProperty)
	g.PATCH("/properties/:id", propertyHandler.ChangeStatusProperty)
	g.GET("/properties/:slug", propertyHandler.GetBySlug)
}
