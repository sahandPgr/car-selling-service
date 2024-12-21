package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/models"
)

type PropertyService struct {
	base *BaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]
}

func NewPropertyService(cfg *config.Config) *PropertyService {
	return &PropertyService{
		base: NewBaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse](cfg,
			[]Preload{{string: "Category"}}),
	}
}

func (s *PropertyService) Create(ctx context.Context, req *dto.CreatePropertyRequest) (res *dto.PropertyResponse, err error) {
	return s.base.Create(ctx, req)
}

func (s *PropertyService) Update(ctx context.Context, id int, req *dto.UpdatePropertyRequest) (res *dto.PropertyResponse, err error) {
	return s.base.Update(ctx, id, req)
}

func (s *PropertyService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

func (s *PropertyService) GetById(ctx context.Context, id int) (res *dto.PropertyResponse, err error) {
	return s.base.GetById(ctx, id)
}

func (s *PropertyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (res *dto.PagedList[dto.PropertyResponse], err error) {
	return s.base.GetByFilter(ctx, req)
}
