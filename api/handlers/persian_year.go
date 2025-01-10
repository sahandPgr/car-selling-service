package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sahandPgr/car-selling-service/api/dto"
	_ "github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

type PersianYearsHandler struct {
	service *services.PersianYearService
}

func NewPersianYearsHandler(cfg *config.Config) *PersianYearsHandler {
	return &PersianYearsHandler{
		service: services.NewPersianYearService(cfg),
	}
}

// CreateCountry godoc
// @Summary Create a PersianYear
// @Description Create a PersianYear
// @Tags PersianYears
// @Accept json
// @Produce json
// @Param request body dto.CreatePersianYearRequest true "Create a PersianYear"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PersianYearResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/persian-years/ [post]
// @Security AuthBearer
func (h *PersianYearsHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdateCountry godoc
// @Summary Update a PersianYear
// @Description Update a PersianYear
// @Tags PersianYears
// @Accept json
// @Produce json
// @Param id path int true "Update a PersianYear"
// @Param request body dto.UpdatePersianYearRequest true "Update a PersianYear"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PersianYearResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/persian-years/{id} [put]
// @Security AuthBearer
func (h *PersianYearsHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCountry godoc
// @Summary Delete a PersianYear
// @Description Delete a PersianYear
// @Tags PersianYears
// @Accept json
// @Produce json
// @Param id path int true "Delete a PersianYear"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/persian-years/{id} [delete]
// @Security AuthBearer
func (h *PersianYearsHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)
}

// GetCountry godoc
// @Summary Get a PersianYear
// @Description Get a PersianYear
// @Tags PersianYears
// @Accept json
// @Produce json
// @Param id path int true "Get a PersianYear"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PersianYearResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/persian-years/{id} [get]
// @Security AuthBearer
func (h *PersianYearsHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}

// GetPersianYears godoc
// @Summary Get PersianYear
// @Description Get PersianYear
// @Tags PersianYears
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.PersianYearResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/persian-years/get-by-filter [post]
// @Security AuthBearer
func (h *PersianYearsHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter(ctx, h.service.GetByFilter)
}
