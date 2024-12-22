package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sahandPgr/car-selling-service/api/dto"
	_ "github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

type CompaniesHandler struct {
	service *services.CompanyService
}

func NewCompaniesHandler(cfg *config.Config) *CompaniesHandler {
	return &CompaniesHandler{
		service: services.NewCompanyService(cfg),
	}
}

// CreateCompany godoc
// @Summary Create a Company
// @Description Create a Company
// @Tags Companies
// @Accept json
// @Produce json
// @Param request body dto.CreateCompanyRequest true "Create a Company"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CompanyResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/companies/ [post]
// @Security AuthBearer
func (h *CompaniesHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdateCompany godoc
// @Summary Update a Company
// @Description Update a Company
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path int true "Update a Company"
// @Param request body dto.UpdateCompanyRequest true "Update a Company"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CompanyResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/companies/{id} [put]
// @Security AuthBearer
func (h *CompaniesHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCompany godoc
// @Summary Delete a Company
// @Description Delete a Company
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path int true "Delete a Company"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/companies/{id} [delete]
// @Security AuthBearer
func (h *CompaniesHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)
}

// GetCompany godoc
// @Summary Get a Company
// @Description Get a Company
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path int true "Get a Company"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CompanyResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/companies/{id} [get]
// @Security AuthBearer
func (h *CompaniesHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}

// GetCompany godoc
// @Summary Get Company
// @Description Get Company
// @Tags Companies
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CompanyResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/companies/get-by-filter [post]
// @Security AuthBearer
func (h *CompaniesHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter(ctx, h.service.GetByFilter)
}
