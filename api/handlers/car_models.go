package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sahandPgr/car-selling-service/api/dto"
	_ "github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

type CarModelsHandler struct {
	service *services.CarModelService
}

func NewCarModelsHandler(cfg *config.Config) *CarModelsHandler {
	return &CarModelsHandler{
		service: services.NewCarModelService(cfg),
	}
}

// CreateCarModel godoc
// @Summary Create a car model
// @Description Create a car model
// @Tags CarModels
// @Accept json
// @Produce json
// @Param request body dto.CreateCarModelRequest true "Create a car model"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-models/ [post]
// @Security AuthBearer
func (h *CarModelsHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdateCarModel godoc
// @Summary Update a car model
// @Description Update a car model
// @Tags CarModels
// @Accept json
// @Produce json
// @Param id path int true "Update a car model"
// @Param request body dto.UpdateCarModelRequest true "Update a car model"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-models/{id} [put]
// @Security AuthBearer
func (h *CarModelsHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCarModel godoc
// @Summary Delete a car model
// @Description Delete a car model
// @Tags CarModels
// @Accept json
// @Produce json
// @Param id path int true "Delete a car model"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-models/{id} [delete]
// @Security AuthBearer
func (h *CarModelsHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)
}

// GetCarModel godoc
// @Summary Get a car model
// @Description Get a car model
// @Tags CarModels
// @Accept json
// @Produce json
// @Param id path int true "Get a car model"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-models/{id} [get]
// @Security AuthBearer
func (h *CarModelsHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}

// GetCarModels godoc
// @Summary Get CarModels
// @Description Get CarModels
// @Tags CarModels
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-models/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelsHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter(ctx, h.service.GetByFilter)
}
