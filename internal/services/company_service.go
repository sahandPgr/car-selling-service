package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/models"
)

type CompanyService struct {
	base *BaseService[models.Company, dto.CreateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse]
}

func NewCompanyService(cfg *config.Config) *CompanyService {
	return &CompanyService{
		base: NewBaseService[models.Company, dto.CreateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse](cfg,
			[]Preload{{string: "Country"}}),
	}
}

// Create
func (s *CompanyService) Create(ctx context.Context, req *dto.CreateCompanyRequest) (res *dto.CompanyResponse, err error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CompanyService) Update(ctx context.Context, id int, req *dto.UpdateCompanyRequest) (res *dto.CompanyResponse, err error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CompanyService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

// Get by Id
func (s *CompanyService) GetById(ctx context.Context, id int) (res *dto.CompanyResponse, err error) {
	return s.base.GetById(ctx, id)
}

func (s *CompanyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CompanyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
