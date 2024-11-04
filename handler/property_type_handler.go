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

type PropertyTypeHandler struct {
	Service *services.PropertyTypeService
}

func NewPropertyTypeHandler(service *services.PropertyTypeService) *PropertyTypeHandler {
	return &PropertyTypeHandler{Service: service}
}

func (h *PropertyTypeHandler) GetPropertyTypes(c echo.Context) error {
	propertyTypes, err := h.Service.GetAll()
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	if propertyTypes == nil {
		propertyTypes = []domain.PropertyType{}
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, propertyTypes)
}

func (h *PropertyTypeHandler) CreatePropertyType(c echo.Context) error {
	var propertyType domain.PropertyType
	if err := c.Bind(&propertyType); err != nil {
		return utils.SendBadRequest(c, localesCommon.ErrorPayload)
	}
	if err := validate.Struct(propertyType); err != nil {
		return utils.SendErrorValidations(c, localesCommon.ErrorPayload, err)
	}
	if err := h.Service.Create(&propertyType); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendCreated(c, localesCommon.SuccessCreated, propertyType)
}

func (h *PropertyTypeHandler) GetPropertyType(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	propertyType, err := h.Service.GetByID(id)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, propertyType)
}

func (h *PropertyTypeHandler) UpdatePropertyType(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	var propertyType domain.PropertyType
	if err = c.Bind(&propertyType); err != nil {
		return utils.SendBadRequest(c, localesCommon.ErrorPayload)
	}
	if err = validate.Struct(propertyType); err != nil {
		return utils.SendErrorValidations(c, localesCommon.ErrorPayload, err)
	}
	propertyType.ID = id
	if err = h.Service.Update(&propertyType); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, propertyType)
}

func (h *PropertyTypeHandler) DeletePropertyType(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	if err = h.Service.Delete(id); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, nil)
}
