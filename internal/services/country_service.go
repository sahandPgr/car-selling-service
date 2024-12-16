package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/db"
	"github.com/sahandPgr/car-selling-service/internal/models"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
)

type CountryService struct {
	base *BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		base: &BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CountryResponse]{
			Database: db.GetDB(),
			Log:      logger.NewLogger(cfg),
			Preloads: []Preload{{string: "Cities"}},
		},
	}
}

// Create
func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (res *dto.CountryResponse, err error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CountryService) Update(ctx context.Context, id int, req *dto.CreateUpdateCountryRequest) (res *dto.CountryResponse, err error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CountryService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

// Get by Id
func (s *CountryService) GetById(ctx context.Context, id int) (res *dto.CountryResponse, err error) {
	return s.base.GetById(ctx, id)
}

func (s *CountryService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CountryResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
