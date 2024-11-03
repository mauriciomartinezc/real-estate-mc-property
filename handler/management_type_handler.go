package handler

import (
	"github.com/labstack/echo/v4"
	localesCommon "github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"github.com/mauriciomartinezc/real-estate-mc-property/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-property/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ManagementTypeHandler struct {
	Service *services.ManagementTypeService
}

func NewManagementTypeHandler(service *services.ManagementTypeService) *ManagementTypeHandler {
	return &ManagementTypeHandler{Service: service}
}

func (h *ManagementTypeHandler) GetManagementTypes(c echo.Context) error {
	managementTypes, err := h.Service.GetAll()
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	if managementTypes == nil {
		managementTypes = []domain.ManagementType{}
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, managementTypes)
}

func (h *ManagementTypeHandler) CreateManagementType(c echo.Context) error {
	var managementType domain.ManagementType
	if err := c.Bind(&managementType); err != nil {
		return utils.SendBadRequest(c, localesCommon.ErrorPayload)
	}
	if err := validate.Struct(managementType); err != nil {
		return utils.SendErrorValidations(c, localesCommon.ErrorPayload, err)
	}
	if err := h.Service.Create(&managementType); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendCreated(c, localesCommon.SuccessCreated, managementType)
}

func (h *ManagementTypeHandler) GetManagementType(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	managementType, err := h.Service.GetByID(id)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, managementType)
}

func (h *ManagementTypeHandler) UpdateManagementType(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	var managementType domain.ManagementType
	if err = c.Bind(&managementType); err != nil {
		return utils.SendBadRequest(c, localesCommon.ErrorPayload)
	}
	if err = validate.Struct(managementType); err != nil {
		return utils.SendErrorValidations(c, localesCommon.ErrorPayload, err)
	}
	managementType.ID = id
	if err = h.Service.Update(&managementType); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, managementType)
}

func (h *ManagementTypeHandler) DeleteManagementType(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	if err = h.Service.Delete(id); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, nil)
}
