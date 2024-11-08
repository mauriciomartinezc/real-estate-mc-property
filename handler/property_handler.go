package handler

import (
	"github.com/labstack/echo/v4"
	localesCommon "github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"github.com/mauriciomartinezc/real-estate-mc-property/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-property/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type PropertyHandler struct {
	Service *services.PropertyService
}

func NewPropertyHandler(service *services.PropertyService) *PropertyHandler {
	return &PropertyHandler{Service: service}
}

func (h *PropertyHandler) GetAllPropertiesPaginated(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	properties, err := h.Service.GetAllPropertiesPaginated(page, limit)

	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}

	if properties == nil {
		properties = domain.SimpleProperties{}
	}

	return utils.SendSuccess(c, localesCommon.SuccessResponse, properties)
}

func (h *PropertyHandler) GetPropertiesByCompanyID(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	companyID := c.QueryParam("companyId")
	properties, err := h.Service.GetPropertiesByCompanyID(companyID, page, limit)

	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}

	if properties == nil {
		properties = domain.SimpleProperties{}
	}

	return utils.SendSuccess(c, localesCommon.SuccessResponse, properties)
}

func (h *PropertyHandler) CreateProperty(c echo.Context) error {
	var property domain.SimpleProperty
	if err := c.Bind(&property); err != nil {
		return utils.SendBadRequest(c, localesCommon.ErrorPayload)
	}
	if err := validate.Struct(property); err != nil {
		return utils.SendErrorValidations(c, localesCommon.ErrorPayload, err)
	}
	if err := h.Service.Create(&property); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendCreated(c, localesCommon.SuccessCreated, property)
}

func (h *PropertyHandler) GetProperty(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	property, err := h.Service.GetByID(id)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, property)
}

func (h *PropertyHandler) GetDetailProperty(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	property, err := h.Service.GetDetailByID(id)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, property)
}

func (h *PropertyHandler) UpdateProperty(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	var property domain.SimpleProperty
	if err = c.Bind(&property); err != nil {
		return utils.SendBadRequest(c, localesCommon.ErrorPayload)
	}
	if err = validate.Struct(property); err != nil {
		return utils.SendErrorValidations(c, localesCommon.ErrorPayload, err)
	}
	property.ID = id
	if err = h.Service.Update(&property); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, property)
}

func (h *PropertyHandler) ChangeStatusProperty(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	property, err := h.Service.GetByID(id)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	property.ID = id
	if err = h.Service.ChangeStatus(property); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, property)
}
