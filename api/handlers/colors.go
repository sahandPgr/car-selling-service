package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sahandPgr/car-selling-service/api/dto"
	_ "github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

type ColorHandler struct {
	service *services.ColorService
}

func NewColorHandler(cfg *config.Config) *ColorHandler {
	return &ColorHandler{
		service: services.NewColorService(cfg),
	}
}

// CreateColor godoc
// @Summary Create a Color
// @Description Create a Color
// @Tags Colors
// @Accept json
// @Produce json
// @Param request body dto.CreateColorRequest true "Create a Color"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.ColorResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/colors/ [post]
// @Security AuthBearer
func (h *ColorHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdateColor godoc
// @Summary Update a Color
// @Description Update a Color
// @Tags Colors
// @Accept json
// @Produce json
// @Param id path int true "Update a Color"
// @Param request body dto.UpdateColorRequest true "Update a Color"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ColorResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/colors/{id} [put]
// @Security AuthBearer
func (h *ColorHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteColor godoc
// @Summary Delete a Color
// @Description Delete a Color
// @Tags Colors
// @Accept json
// @Produce json
// @Param id path int true "Delete a Color"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/colors/{id} [delete]
// @Security AuthBearer
func (h *ColorHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)
}

// GetByIdColor godoc
// @Summary GetById a Color
// @Description GetById a Color
// @Tags Colors
// @Accept json
// @Produce json
// @Param id path int true "GetById a Color"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ColorResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/colors/{id} [get]
// @Security AuthBearer
func (h *ColorHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}

// GetByFilterColor godoc
// @Summary GetByFilter a Color
// @Description GetByFilter a Color
// @Tags Colors
// @Accept json
// @Produce json
// @Param request body dto.PaginationInputWithFilter true "GetByFilter a Color"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.ColorResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/colors/get-by-filter [post]
// @Security AuthBearer
func (h *ColorHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter(ctx, h.service.GetByFilter)
}
