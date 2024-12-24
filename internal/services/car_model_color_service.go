package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/models"
)

type CarModelColorService struct {
	base *BaseService[models.CarModelColor, dto.CreateCarModelColorRequest, dto.UpdateCarModelColorRequest, dto.CarModelColorResponse]
}

func NewCarModelColorService(cfg *config.Config) *CarModelColorService {
	return &CarModelColorService{
		base: NewBaseService[models.CarModelColor, dto.CreateCarModelColorRequest, dto.UpdateCarModelColorRequest, dto.CarModelColorResponse](cfg,
			[]Preload{{string: "Color"}}),
	}
}

func (s *CarModelColorService) Create(ctx context.Context, req *dto.CreateCarModelColorRequest) (res *dto.CarModelColorResponse, err error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelColorService) Update(ctx context.Context, id int, req *dto.UpdateCarModelColorRequest) (res *dto.CarModelColorResponse, err error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelColorService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

func (s *CarModelColorService) GetById(ctx context.Context, id int) (res *dto.CarModelColorResponse, err error) {
	return s.base.GetById(ctx, id)
}

func (s *CarModelColorService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (res *dto.PagedList[dto.CarModelColorResponse], err error) {
	return s.base.GetByFilter(ctx, req)
}
