package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/models"
)

type CarTypeService struct {
	base *BaseService[models.CarType, dto.CreateCarTypeRequest, dto.UpdateCarTypeRequest, dto.CarTypeResponse]
}

func NewCarTypeService(cfg *config.Config) *CarTypeService {
	return &CarTypeService{
		base: NewBaseService[models.CarType, dto.CreateCarTypeRequest, dto.UpdateCarTypeRequest, dto.CarTypeResponse](cfg,
			[]Preload{}),
	}
}

// Create
func (s *CarTypeService) Create(ctx context.Context, req *dto.CreateCarTypeRequest) (res *dto.CarTypeResponse, err error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarTypeService) Update(ctx context.Context, id int, req *dto.UpdateCarTypeRequest) (res *dto.CarTypeResponse, err error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarTypeService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

// Get by Id
func (s *CarTypeService) GetById(ctx context.Context, id int) (res *dto.CarTypeResponse, err error) {
	return s.base.GetById(ctx, id)
}

func (s *CarTypeService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarTypeResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
