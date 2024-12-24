package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sahandPgr/car-selling-service/api/dto"
	_ "github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

type CarModelColorHandler struct {
	service *services.CarModelColorService
}

func NewCarModelColorHandler(cfg *config.Config) *CarModelColorHandler {
	return &CarModelColorHandler{
		service: services.NewCarModelColorService(cfg),
	}
}

// CreateCarModelColor godoc
// @Summary Create a CarModelColor
// @Description Create a CarModelColor
// @Tags CarModelColors
// @Accept json
// @Produce json
// @Param request body dto.CreateCarModelColorRequest true "Create a CarModelColor"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelColorResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-model-colors/ [post]
// @Security AuthBearer
func (h *CarModelColorHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdateCarModelColor godoc
// @Summary Update a CarModelColor
// @Description Update a CarModelColor
// @Tags CarModelColors
// @Accept json
// @Produce json
// @Param id path int true "Update a CarModelColor"
// @Param request body dto.UpdateCarModelColorRequest true "Update a CarModelColor"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelColorResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-model-colors/{id} [put]
// @Security AuthBearer
func (h *CarModelColorHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCarModelColor godoc
// @Summary Delete a CarModelColor
// @Description Delete a CarModelColor
// @Tags CarModelColors
// @Accept json
// @Produce json
// @Param id path int true "Delete a CarModelColor"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-model-colors/{id} [delete]
// @Security AuthBearer
func (h *CarModelColorHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)
}

// GetByIdCarModelColor godoc
// @Summary GetById a CarModelColor
// @Description GetById a CarModelColor
// @Tags CarModelColors
// @Accept json
// @Produce json
// @Param id path int true "GetById a CarModelColor"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelColorResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-model-colors/{id} [get]
// @Security AuthBearer
func (h *CarModelColorHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}

// GetByFilterCarModelColor godoc
// @Summary GetByFilter a CarModelColor
// @Description GetByFilter a CarModelColor
// @Tags CarModelColors
// @Accept json
// @Produce json
// @Param request body dto.PaginationInputWithFilter true "GetByFilter a CarModelColor"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelColorResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-model-colors/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelColorHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter(ctx, h.service.GetByFilter)
}
