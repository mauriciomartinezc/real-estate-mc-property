package handlers

import (
	"github.com/labstack/echo/v4"
	localesCommon "github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"github.com/mauriciomartinezc/real-estate-mc-property/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-property/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureHandler struct {
	Service *services.FeatureService
}

func NewFeatureHandler(service *services.FeatureService) *FeatureHandler {
	return &FeatureHandler{Service: service}
}

func (h *FeatureHandler) GetFeatures(c echo.Context) error {
	features, err := h.Service.GetAll()
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	if features == nil {
		features = []domain.Feature{}
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, features)
}

func (h *FeatureHandler) GetFeaturesGroupedByType(c echo.Context) error {
	featuresGroupedByType, err := h.Service.GetFeaturesGroupedByType()
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, featuresGroupedByType)
}

func (h *FeatureHandler) CreateFeature(c echo.Context) error {
	var feature domain.Feature
	if err := c.Bind(&feature); err != nil {
		return utils.SendBadRequest(c, localesCommon.ErrorPayload)
	}
	if err := validate.Struct(feature); err != nil {
		return utils.SendErrorValidations(c, localesCommon.ErrorPayload, err)
	}
	if err := h.Service.Create(&feature); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendCreated(c, localesCommon.SuccessCreated, feature)
}

func (h *FeatureHandler) GetFeature(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	feature, err := h.Service.GetByID(id)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, feature)
}

func (h *FeatureHandler) UpdateFeature(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	var feature domain.Feature
	if err = c.Bind(&feature); err != nil {
		return utils.SendBadRequest(c, localesCommon.ErrorPayload)
	}
	if err = validate.Struct(feature); err != nil {
		return utils.SendErrorValidations(c, localesCommon.ErrorPayload, err)
	}
	feature.ID = id
	if err = h.Service.Update(&feature); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, feature)
}

func (h *FeatureHandler) DeleteFeature(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	if err = h.Service.Delete(id); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, nil)
}
