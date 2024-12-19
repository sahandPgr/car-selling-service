package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sahandPgr/car-selling-service/api/dto"
	_ "github.com/sahandPgr/car-selling-service/api/helper"
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
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/countries/ [post]
// @Security AuthBearer
func (h *CountriesHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
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
	Update(ctx, h.service.Update)
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
	Delete(ctx, h.service.Delete)
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
	GetById(ctx, h.service.GetById)
}

// GetCountries godoc
// @Summary Get countries
// @Description Get countries
// @Tags Countries
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CountryResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/countries/get-by-filter [post]
// @Security AuthBearer
func (h *CountriesHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter(ctx, h.service.GetByFilter)
}
