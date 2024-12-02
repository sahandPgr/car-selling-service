package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

type CountriesHandler struct {
	service *services.CountryService
}

func NewCountriesHandler(cfg *config.Config) *CountriesHandler {
	return &CountriesHandler{
		service: services.NewCountryService(cfg),
	}
}

// CreateCountry godoc
// @Summary Create a country
// @Description Create a country
// @Tags Countries
// @Accept json
// @Produce json
// @Param request body dto.CreateUpdateCountryRequest true "Create a country"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/countries/ [post]
// @Security AuthBearer
func (h *CountriesHandler) Create(ctx *gin.Context) {
	var dto = new(dto.CreateUpdateCountryRequest)
	err := ctx.ShouldBindJSON(dto)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.BadRequest, helper.GetBaseHttpResponseWithValidation(nil, false, helper.BadRequest, err))
		return
	}
	res, err := h.service.Create(ctx, dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Created, helper.GetBaseHttpResponse(res, true, helper.Created))
}

// UpdateCountry godoc
// @Summary Update a country
// @Description Update a country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "Update a country"
// @Param request body dto.CreateUpdateCountryRequest true "Update a country"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/countries/{id} [put]
// @Security AuthBearer
func (h *CountriesHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var dto = new(dto.CreateUpdateCountryRequest)
	err := ctx.ShouldBindJSON(dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}
	res, err := h.service.Update(ctx, id, dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(res, true, helper.Success))
}

// DeleteCountry godoc
// @Summary Delete a country
// @Description Delete a country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "Delete a country"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/countries/{id} [delete]
// @Security AuthBearer
func (h *CountriesHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(helper.NotFound, helper.GetBaseHttpResponse(nil, false, helper.NotFound))
		return
	}
	err := h.service.Delete(ctx, id)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(nil, true, helper.Success))
}

// GetCountry godoc
// @Summary Get a country
// @Description Get a country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "Get a country"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/countries/{id} [get]
// @Security AuthBearer
func (h *CountriesHandler) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(helper.NotFound, helper.GetBaseHttpResponse(nil, false, helper.NotFound))
		return
	}
	res, err := h.service.GetById(ctx, id)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(res, true, helper.Success))
}
