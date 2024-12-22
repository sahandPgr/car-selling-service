package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sahandPgr/car-selling-service/api/dto"
	_ "github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

type GearboxesHandler struct {
	service *services.GearboxService
}

func NewGearboxesHandler(cfg *config.Config) *GearboxesHandler {
	return &GearboxesHandler{
		service: services.NewGearboxService(cfg),
	}
}

// CreateGearbox godoc
// @Summary Create a gearbox
// @Description Create a gearbox
// @Tags Gearboxes
// @Accept json
// @Produce json
// @Param request body dto.CreateGearboxRequest true "Create a gearbox"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/gearboxes/ [post]
// @Security AuthBearer
func (h *GearboxesHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdateGearbox godoc
// @Summary Update a gearbox
// @Description Update a gearbox
// @Tags Gearboxes
// @Accept json
// @Produce json
// @Param id path int true "Update a gearbox"
// @Param request body dto.UpdateGearboxRequest true "Update a gearbox"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/gearboxes/{id} [put]
// @Security AuthBearer
func (h *GearboxesHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteGearbox godoc
// @Summary Delete a gearbox
// @Description Delete a gearbox
// @Tags Gearboxes
// @Accept json
// @Produce json
// @Param id path int true "Delete a gearbox"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/gearboxes/{id} [delete]
// @Security AuthBearer
func (h *GearboxesHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)
}

// GetGearbox godoc
// @Summary Get a gearbox
// @Description Get a gearbox
// @Tags Gearboxes
// @Accept json
// @Produce json
// @Param id path int true "Get a gearbox"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/gearboxes/{id} [get]
// @Security AuthBearer
func (h *GearboxesHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}

// GetGearboxes godoc
// @Summary Get Gearboxes
// @Description Get Gearboxes
// @Tags Gearboxes
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.GearboxResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/gearboxes/get-by-filter [post]
// @Security AuthBearer
func (h *GearboxesHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter(ctx, h.service.GetByFilter)
}
