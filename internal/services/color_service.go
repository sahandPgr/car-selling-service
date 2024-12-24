package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/models"
)

type ColorService struct {
	base *BaseService[models.Color, dto.CreateColorRequest, dto.UpdateColorRequest, dto.ColorResponse]
}

func NewColorService(cfg *config.Config) *ColorService {
	return &ColorService{
		base: NewBaseService[models.Color, dto.CreateColorRequest, dto.UpdateColorRequest, dto.ColorResponse](cfg,
			[]Preload{}),
	}
}

func (s *ColorService) Create(ctx context.Context, req *dto.CreateColorRequest) (res *dto.ColorResponse, err error) {
	return s.base.Create(ctx, req)
}

func (s *ColorService) Update(ctx context.Context, id int, req *dto.UpdateColorRequest) (res *dto.ColorResponse, err error) {
	return s.base.Update(ctx, id, req)
}

func (s *ColorService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

func (s *ColorService) GetById(ctx context.Context, id int) (res *dto.ColorResponse, err error) {
	return s.base.GetById(ctx, id)
}

func (s *ColorService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (res *dto.PagedList[dto.ColorResponse], err error) {
	return s.base.GetByFilter(ctx, req)
}
