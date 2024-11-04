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

type FeatureTypeHandler struct {
	Service *services.FeatureTypeService
}

func NewFeatureTypeHandler(service *services.FeatureTypeService) *FeatureTypeHandler {
	return &FeatureTypeHandler{Service: service}
}

func (h *FeatureTypeHandler) GetFeatureTypes(c echo.Context) error {
	featureTypes, err := h.Service.GetAll()
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	if featureTypes == nil {
		featureTypes = []domain.FeatureType{}
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, featureTypes)
}

func (h *FeatureTypeHandler) CreateFeatureType(c echo.Context) error {
	var featureType domain.FeatureType
	if err := c.Bind(&featureType); err != nil {
		return utils.SendBadRequest(c, localesCommon.ErrorPayload)
	}
	if err := validate.Struct(featureType); err != nil {
		return utils.SendErrorValidations(c, localesCommon.ErrorPayload, err)
	}
	if err := h.Service.Create(&featureType); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendCreated(c, localesCommon.SuccessCreated, featureType)
}

func (h *FeatureTypeHandler) GetFeatureType(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	featureType, err := h.Service.GetByID(id)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, featureType)
}

func (h *FeatureTypeHandler) UpdateFeatureType(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	var featureType domain.FeatureType
	if err = c.Bind(&featureType); err != nil {
		return utils.SendBadRequest(c, localesCommon.ErrorPayload)
	}
	if err = validate.Struct(featureType); err != nil {
		return utils.SendErrorValidations(c, localesCommon.ErrorPayload, err)
	}
	featureType.ID = id
	if err = h.Service.Update(&featureType); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, featureType)
}

func (h *FeatureTypeHandler) DeleteFeatureType(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	if err = h.Service.Delete(id); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, nil)
}
