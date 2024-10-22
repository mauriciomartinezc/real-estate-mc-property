package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-property/handler"
	"github.com/mauriciomartinezc/real-estate-mc-property/repository"
	"github.com/mauriciomartinezc/real-estate-mc-property/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(e *echo.Echo, db *mongo.Database) {
	// Rutas para ManagementType
	managementRepo := repository.NewManagementTypeRepository(db)
	managementService := services.NewManagementTypeService(managementRepo)
	managementController := handler.NewManagementTypeController(managementService)

	e.POST("/managementTypes", managementController.CreateManagementType)
	e.GET("/managementTypes/:id", managementController.GetManagementType)
	e.PUT("/managementTypes/:id", managementController.UpdateManagementType)
	e.DELETE("/managementTypes/:id", managementController.DeleteManagementType)
}
