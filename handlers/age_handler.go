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

type AgeHandler struct {
	Service *services.AgeService
}

func NewAgeHandler(service *services.AgeService) *AgeHandler {
	return &AgeHandler{Service: service}
}

func (h *AgeHandler) GetAges(c echo.Context) error {
	ages, err := h.Service.GetAll()
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	if ages == nil {
		ages = []domain.Age{}
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, ages)
}

func (h *AgeHandler) CreateAge(c echo.Context) error {
	var age domain.Age
	if err := c.Bind(&age); err != nil {
		return utils.SendBadRequest(c, localesCommon.ErrorPayload)
	}
	if err := validate.Struct(age); err != nil {
		return utils.SendErrorValidations(c, localesCommon.ErrorPayload, err)
	}
	if err := h.Service.Create(&age); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendCreated(c, localesCommon.SuccessCreated, age)
}

func (h *AgeHandler) GetAge(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	age, err := h.Service.GetByID(id)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, age)
}

func (h *AgeHandler) UpdateAge(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	var age domain.Age
	if err = c.Bind(&age); err != nil {
		return utils.SendBadRequest(c, localesCommon.ErrorPayload)
	}
	if err = validate.Struct(age); err != nil {
		return utils.SendErrorValidations(c, localesCommon.ErrorPayload, err)
	}
	age.ID = id
	if err = h.Service.Update(&age); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, age)
}

func (h *AgeHandler) DeleteAge(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidId)
	}
	if err = h.Service.Delete(id); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesCommon.SuccessResponse, nil)
}
