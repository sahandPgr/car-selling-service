package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sahandPgr/car-selling-service/api/dto"
	_ "github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

type CarTypesHandler struct {
	service *services.CarTypeService
}

func NewCarTypesHandler(cfg *config.Config) *CarTypesHandler {
	return &CarTypesHandler{
		service: services.NewCarTypeService(cfg),
	}
}

// CreateCarType godoc
// @Summary Create a car type
// @Description Create a car type
// @Tags CarTypes
// @Accept json
// @Produce json
// @Param request body dto.CreateCarTypeRequest true "Create a car type"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarTypeResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-types/ [post]
// @Security AuthBearer
func (h *CarTypesHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdateCarType godoc
// @Summary Update a car type
// @Description Update a car type
// @Tags CarTypes
// @Accept json
// @Produce json
// @Param id path int true "Update a car type"
// @Param request body dto.UpdateCarTypeRequest true "Update a car type"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarTypeResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-types/{id} [put]
// @Security AuthBearer
func (h *CarTypesHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCarType godoc
// @Summary Delete a car type
// @Description Delete a car type
// @Tags CarTypes
// @Accept json
// @Produce json
// @Param id path int true "Delete a car type"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-types/{id} [delete]
// @Security AuthBearer
func (h *CarTypesHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)
}

// GetCarType godoc
// @Summary Get a car type
// @Description Get a car type
// @Tags CarTypes
// @Accept json
// @Produce json
// @Param id path int true "Get a car type"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarTypeResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-types/{id} [get]
// @Security AuthBearer
func (h *CarTypesHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}

// GetCarTypes godoc
// @Summary Get CarTypes
// @Description Get CarTypes
// @Tags CarTypes
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarTypeResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-types/get-by-filter [post]
// @Security AuthBearer
func (h *CarTypesHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter(ctx, h.service.GetByFilter)
}
