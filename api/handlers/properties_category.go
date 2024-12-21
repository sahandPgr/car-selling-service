package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sahandPgr/car-selling-service/api/dto"
	_ "github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

type PropertiesCategoryHandler struct {
	service *services.PropertyCategoryService
}

func NewPropertiesCategoryHandler(cfg *config.Config) *PropertiesCategoryHandler {
	return &PropertiesCategoryHandler{
		service: services.NewPropertyCategoryService(cfg),
	}
}

// CreateCountry godoc
// @Summary Create a property-category
// @Description Create a property-category
// @Tags PropertiesCategory
// @Accept json
// @Produce json
// @Param request body dto.CreatePropertyCategoryRequest true "Create a property-category"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PropertyCategoryResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/property-category/ [post]
// @Security AuthBearer
func (h *PropertiesCategoryHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdateCountry godoc
// @Summary Update a property-category
// @Description Update a property-category
// @Tags PropertiesCategory
// @Accept json
// @Produce json
// @Param id path int true "Update a property-category"
// @Param request body dto.UpdatePropertyCategoryRequest true "Update a property-category"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyCategoryResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/property-category/{id} [put]
// @Security AuthBearer
func (h *PropertiesCategoryHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCountry godoc
// @Summary Delete a property-category
// @Description Delete a property-category
// @Tags PropertiesCategory
// @Accept json
// @Produce json
// @Param id path int true "Delete a property-category"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/property-category/{id} [delete]
// @Security AuthBearer
func (h *PropertiesCategoryHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)
}

// GetCountry godoc
// @Summary Get a property-category
// @Description Get a property-category
// @Tags PropertiesCategory
// @Accept json
// @Produce json
// @Param id path int true "Get a property-category"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyCategoryResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/property-category/{id} [get]
// @Security AuthBearer
func (h *PropertiesCategoryHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}

// GetPropertiesCategory godoc
// @Summary Get property-category
// @Description Get property-category
// @Tags PropertiesCategory
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.PropertyCategoryResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/property-category/get-by-filter [post]
// @Security AuthBearer
func (h *PropertiesCategoryHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter(ctx, h.service.GetByFilter)
}
