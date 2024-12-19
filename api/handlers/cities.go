package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sahandPgr/car-selling-service/api/dto"
	_ "github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{
		service: services.NewCityService(cfg),
	}
}

// CreateCity godoc
// @Summary Create a City
// @Description Create a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param request body dto.CreateCityRequest true "Create a City"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CityResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/cities/ [post]
// @Security AuthBearer
func (h *CityHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdateCity godoc
// @Summary Update a City
// @Description Update a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "Update a City"
// @Param request body dto.UpdateCityRequest true "Update a City"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CityResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/cities/{id} [put]
// @Security AuthBearer
func (h *CityHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCity godoc
// @Summary Delete a City
// @Description Delete a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "Delete a City"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/cities/{id} [delete]
// @Security AuthBearer
func (h *CityHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)
}

// GetByIdCity godoc
// @Summary GetById a City
// @Description GetById a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "GetById a City"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CityResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/cities/{id} [get]
// @Security AuthBearer
func (h *CityHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}

// GetByFilterCity godoc
// @Summary GetByFilter a City
// @Description GetByFilter a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param request body dto.PaginationInputWithFilter true "GetByFilter a City"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CityResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/cities/get-by-filter [post]
// @Security AuthBearer
func (h *CityHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter(ctx, h.service.GetByFilter)
}
