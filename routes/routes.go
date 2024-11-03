package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-property/handler"
	"github.com/mauriciomartinezc/real-estate-mc-property/repository"
	"github.com/mauriciomartinezc/real-estate-mc-property/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(e *echo.Echo, db *mongo.Database) {
	group := e.Group("api")
	// Rutas para ManagementType
	managementRepo := repository.NewManagementTypeRepository(db)
	managementService := services.NewManagementTypeService(managementRepo)
	managementHandler := handler.NewManagementTypeHandler(managementService)

	group.GET("/managementTypes", managementHandler.GetManagementTypes)
	group.POST("/managementTypes", managementHandler.CreateManagementType)
	group.GET("/managementTypes/:id", managementHandler.GetManagementType)
	group.PUT("/managementTypes/:id", managementHandler.UpdateManagementType)
	group.DELETE("/managementTypes/:id", managementHandler.DeleteManagementType)
}
