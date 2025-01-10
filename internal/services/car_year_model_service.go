package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/models"
)

type CarModelYearService struct {
	base *BaseService[models.CarModelYear, dto.CreateCarModelYearRequest, dto.UpdateCarModelYearRequest, dto.CarModelYearResponse]
}

func NewCarModelYearService(cfg *config.Config) *CarModelYearService {
	return &CarModelYearService{
		base: NewBaseService[models.CarModelYear, dto.CreateCarModelYearRequest, dto.UpdateCarModelYearRequest, dto.CarModelYearResponse](cfg,
			[]Preload{}),
	}
}

func (s *CarModelYearService) Create(ctx context.Context, req *dto.CreateCarModelYearRequest) (res *dto.CarModelYearResponse, err error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelYearService) Update(ctx context.Context, id int, req *dto.UpdateCarModelYearRequest) (res *dto.CarModelYearResponse, err error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelYearService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

func (s *CarModelYearService) GetById(ctx context.Context, id int) (res *dto.CarModelYearResponse, err error) {
	return s.base.GetById(ctx, id)
}

func (s *CarModelYearService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (res *dto.PagedList[dto.CarModelYearResponse], err error) {
	return s.base.GetByFilter(ctx, req)
}
