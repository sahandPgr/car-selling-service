package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/models"
)

type GearboxService struct {
	base *BaseService[models.Gearbox, dto.CreateGearboxRequest, dto.UpdateGearboxRequest, dto.GearboxResponse]
}

func NewGearboxService(cfg *config.Config) *GearboxService {
	return &GearboxService{
		base: NewBaseService[models.Gearbox, dto.CreateGearboxRequest, dto.UpdateGearboxRequest, dto.GearboxResponse](cfg,
			[]Preload{}),
	}
}

// Create
func (s *GearboxService) Create(ctx context.Context, req *dto.CreateGearboxRequest) (res *dto.GearboxResponse, err error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *GearboxService) Update(ctx context.Context, id int, req *dto.UpdateGearboxRequest) (res *dto.GearboxResponse, err error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *GearboxService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

// Get by Id
func (s *GearboxService) GetById(ctx context.Context, id int) (res *dto.GearboxResponse, err error) {
	return s.base.GetById(ctx, id)
}

func (s *GearboxService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GearboxResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
