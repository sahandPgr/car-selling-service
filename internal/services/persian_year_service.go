package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/models"
)

type PersianYearService struct {
	base *BaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]
}

func NewPersianYearService(cfg *config.Config) *PersianYearService {
	return &PersianYearService{
		base: NewBaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse](cfg,
			[]Preload{}),
	}
}

func (s *PersianYearService) Create(ctx context.Context, req *dto.CreatePersianYearRequest) (res *dto.PersianYearResponse, err error) {
	return s.base.Create(ctx, req)
}

func (s *PersianYearService) Update(ctx context.Context, id int, req *dto.UpdatePersianYearRequest) (res *dto.PersianYearResponse, err error) {
	return s.base.Update(ctx, id, req)
}

func (s *PersianYearService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

func (s *PersianYearService) GetById(ctx context.Context, id int) (res *dto.PersianYearResponse, err error) {
	return s.base.GetById(ctx, id)
}

func (s *PersianYearService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (res *dto.PagedList[dto.PersianYearResponse], err error) {
	return s.base.GetByFilter(ctx, req)
}
