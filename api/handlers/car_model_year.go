package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sahandPgr/car-selling-service/api/dto"
	_ "github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

type CarModelYearHandler struct {
	service *services.CarModelYearService
}

func NewCarModelYearHandler(cfg *config.Config) *CarModelYearHandler {
	return &CarModelYearHandler{
		service: services.NewCarModelYearService(cfg),
	}
}

// CreateCarModelYear godoc
// @Summary Create a CarModelYear
// @Description Create a CarModelYear
// @Tags CarModelYears
// @Accept json
// @Produce json
// @Param request body dto.CreateCarModelYearRequest true "Create a CarModelYear"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-model-years/ [post]
// @Security AuthBearer
func (h *CarModelYearHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdateCarModelYear godoc
// @Summary Update a CarModelYear
// @Description Update a CarModelYear
// @Tags CarModelYears
// @Accept json
// @Produce json
// @Param id path int true "Update a CarModelYear"
// @Param request body dto.UpdateCarModelYearRequest true "Update a CarModelYear"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-model-years/{id} [put]
// @Security AuthBearer
func (h *CarModelYearHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCarModelYear godoc
// @Summary Delete a CarModelYear
// @Description Delete a CarModelYear
// @Tags CarModelYears
// @Accept json
// @Produce json
// @Param id path int true "Delete a CarModelYear"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-model-years/{id} [delete]
// @Security AuthBearer
func (h *CarModelYearHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)
}

// GetByIdCarModelYear godoc
// @Summary GetById a CarModelYear
// @Description GetById a CarModelYear
// @Tags CarModelYears
// @Accept json
// @Produce json
// @Param id path int true "GetById a CarModelYear"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-model-years/{id} [get]
// @Security AuthBearer
func (h *CarModelYearHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}

// GetByFilterCarModelYear godoc
// @Summary GetByFilter a CarModelYear
// @Description GetByFilter a CarModelYear
// @Tags CarModelYears
// @Accept json
// @Produce json
// @Param request body dto.PaginationInputWithFilter true "GetByFilter a CarModelYear"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelYearResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/car-model-years/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelYearHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter(ctx, h.service.GetByFilter)
}
