package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/api/helper"
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
	dto := new(dto.CreateCityRequest)
	err := ctx.ShouldBindJSON(dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}
	res, err := h.service.Create(ctx, dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Created, helper.GetBaseHttpResponse(res, true, helper.Created))
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
	id, _ := strconv.Atoi(ctx.Param("id"))
	dto := new(dto.UpdateCityRequest)
	err := ctx.ShouldBindJSON(dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}
	res, err := h.service.Update(ctx, id, dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(res, true, helper.Success))
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
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(helper.NotFound, helper.GetBaseHttpResponse(nil, false, helper.NotFound))
		return
	}
	err := h.service.Delete(ctx, id)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(nil, true, helper.Success))
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
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(helper.NotFound, helper.GetBaseHttpResponse(nil, false, helper.NotFound))
		return
	}
	res, err := h.service.GetById(ctx, id)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(res, true, helper.Success))
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
	var dto = new(dto.PaginationInputWithFilter)
	err := ctx.ShouldBindJSON(dto)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.BadRequest, helper.GetBaseHttpResponseWithValidation(nil, false, helper.BadRequest, err))
		return
	}
	res, err := h.service.GetByFilter(ctx, dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(res, true, helper.Success))
}
