package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/models"
)

type CarModelService struct {
	base *BaseService[models.CarModel, dto.CreateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse]
}

func NewCarModelService(cfg *config.Config) *CarModelService {
	return &CarModelService{
		base: NewBaseService[models.CarModel, dto.CreateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse](cfg,
			[]Preload{{string: "Company.Country"}, {string: "CarType"}, {string: "Gearbox"},
				{string: "CarModelColors.Color"}, {string: "CarModelPersianYears.Year"},
				{string: "CarModelProperties.Property"}, {string: "CarModelImages.Image"},
				{string: "CarModelComments.User"}}),
	}
}

// Create
func (s *CarModelService) Create(ctx context.Context, req *dto.CreateCarModelRequest) (res *dto.CarModelResponse, err error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelService) Update(ctx context.Context, id int, req *dto.UpdateCarModelRequest) (res *dto.CarModelResponse, err error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

// Get by Id
func (s *CarModelService) GetById(ctx context.Context, id int) (res *dto.CarModelResponse, err error) {
	return s.base.GetById(ctx, id)
}

func (s *CarModelService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
