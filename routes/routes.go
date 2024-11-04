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
}

func managementType(g *echo.Group, db *mongo.Database) {
	repo := repository.NewManagementTypeRepository(db)
	service := services.NewManagementTypeService(repo)
	managementTypeHandler := handler.NewManagementTypeHandler(service)

	g.GET("/managementTypes", managementTypeHandler.GetManagementTypes)
	g.POST("/managementTypes", managementTypeHandler.CreateManagementType)
	g.GET("/managementTypes/:id", managementTypeHandler.GetManagementType)
	g.PUT("/managementTypes/:id", managementTypeHandler.UpdateManagementType)
	g.DELETE("/managementTypes/:id", managementTypeHandler.DeleteManagementType)
}

func age(g *echo.Group, db *mongo.Database) {
	repo := repository.NewAgeRepository(db)
	service := services.NewAgeService(repo)
	ageHandler := handler.NewAgeHandler(service)

	g.GET("/ages", ageHandler.GetAges)
	g.POST("/ages", ageHandler.CreateAge)
	g.GET("/ages/:id", ageHandler.GetAge)
	g.PUT("/ages/:id", ageHandler.UpdateAge)
	g.DELETE("/ages/:id", ageHandler.DeleteAge)
}
