package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"github.com/mauriciomartinezc/real-estate-mc-property/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type ManagementTypeController struct {
	Service *services.ManagementTypeService
}

func NewManagementTypeController(service *services.ManagementTypeService) *ManagementTypeController {
	return &ManagementTypeController{Service: service}
}

func (c *ManagementTypeController) CreateManagementType(ctx echo.Context) error {
	var managementType domain.ManagementType
	if err := ctx.Bind(&managementType); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := c.Service.Create(&managementType); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not create management type"})
	}
	return ctx.JSON(http.StatusCreated, managementType)
}

func (c *ManagementTypeController) GetManagementType(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}
	managementType, err := c.Service.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Management type not found"})
	}
	return ctx.JSON(http.StatusOK, managementType)
}

func (c *ManagementTypeController) UpdateManagementType(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}
	var managementType domain.ManagementType
	if err := ctx.Bind(&managementType); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	managementType.ID = id
	if err := c.Service.Update(&managementType); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not update management type"})
	}
	return ctx.JSON(http.StatusOK, managementType)
}

func (c *ManagementTypeController) DeleteManagementType(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}
	if err := c.Service.Delete(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not delete management type"})
	}
	return ctx.JSON(http.StatusNoContent, nil)
}
