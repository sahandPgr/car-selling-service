package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/models"
)

type CityService struct {
	base *BaseService[models.City, dto.CreateCityRequest, dto.UpdateCityRequest, dto.CityResponse]
}

func NewCityService(cfg *config.Config) *CityService {
	return &CityService{
		base: NewBaseService[models.City, dto.CreateCityRequest, dto.UpdateCityRequest, dto.CityResponse](cfg,
			[]Preload{{string: "Country"}}),
	}
}

func (s *CityService) Create(ctx context.Context, req *dto.CreateCityRequest) (res *dto.CityResponse, err error) {
	return s.base.Create(ctx, req)
}

func (s *CityService) Update(ctx context.Context, id int, req *dto.UpdateCityRequest) (res *dto.CityResponse, err error) {
	return s.base.Update(ctx, id, req)
}

func (s *CityService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

func (s *CityService) GetById(ctx context.Context, id int) (res *dto.CityResponse, err error) {
	return s.base.GetById(ctx, id)
}

func (s *CityService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (res *dto.PagedList[dto.CityResponse], err error) {
	return s.base.GetByFilter(ctx, req)
}
