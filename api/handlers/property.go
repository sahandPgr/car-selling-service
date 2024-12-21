package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sahandPgr/car-selling-service/api/dto"
	_ "github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

type PropertiesHandler struct {
	service *services.PropertyService
}

func NewPropertiesHandler(cfg *config.Config) *PropertiesHandler {
	return &PropertiesHandler{
		service: services.NewPropertyService(cfg),
	}
}

// CreateCountry godoc
// @Summary Create a Property
// @Description Create a Property
// @Tags Properties
// @Accept json
// @Produce json
// @Param request body dto.CreatePropertyRequest true "Create a Property"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PropertyResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/properties/ [post]
// @Security AuthBearer
func (h *PropertiesHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdateCountry godoc
// @Summary Update a Property
// @Description Update a Property
// @Tags Properties
// @Accept json
// @Produce json
// @Param id path int true "Update a Property"
// @Param request body dto.UpdatePropertyRequest true "Update a Property"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/properties/{id} [put]
// @Security AuthBearer
func (h *PropertiesHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCountry godoc
// @Summary Delete a Property
// @Description Delete a Property
// @Tags Properties
// @Accept json
// @Produce json
// @Param id path int true "Delete a Property"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/properties/{id} [delete]
// @Security AuthBearer
func (h *PropertiesHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)
}

// GetCountry godoc
// @Summary Get a Property
// @Description Get a Property
// @Tags Properties
// @Accept json
// @Produce json
// @Param id path int true "Get a Property"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/properties/{id} [get]
// @Security AuthBearer
func (h *PropertiesHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}

// GetProperties godoc
// @Summary Get Property
// @Description Get Property
// @Tags Properties
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.PropertyResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/properties/get-by-filter [post]
// @Security AuthBearer
func (h *PropertiesHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter(ctx, h.service.GetByFilter)
}
