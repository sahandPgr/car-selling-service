package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/models"
)

type PropertyCategoryService struct {
	base *BaseService[models.PropertyCategory, dto.CreatePropertyCategoryRequest, dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse]
}

func NewPropertyCategoryService(cfg *config.Config) *PropertyCategoryService {
	return &PropertyCategoryService{
		base: NewBaseService[models.PropertyCategory, dto.CreatePropertyCategoryRequest, dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse](cfg,
			[]Preload{{string: "Properties"}}),
	}
}

func (s *PropertyCategoryService) Create(ctx context.Context, req *dto.CreatePropertyCategoryRequest) (res *dto.PropertyCategoryResponse, err error) {
	return s.base.Create(ctx, req)
}

func (s *PropertyCategoryService) Update(ctx context.Context, id int, req *dto.UpdatePropertyCategoryRequest) (res *dto.PropertyCategoryResponse, err error) {
	return s.base.Update(ctx, id, req)
}

func (s *PropertyCategoryService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

func (s *PropertyCategoryService) GetById(ctx context.Context, id int) (res *dto.PropertyCategoryResponse, err error) {
	return s.base.GetById(ctx, id)
}

func (s *PropertyCategoryService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (res *dto.PagedList[dto.PropertyCategoryResponse], err error) {
	return s.base.GetByFilter(ctx, req)
}
