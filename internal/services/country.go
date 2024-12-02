package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/constatns"
	"github.com/sahandPgr/car-selling-service/internal/db"
	"github.com/sahandPgr/car-selling-service/internal/models"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
	"gorm.io/gorm"
)

type CountryService struct {
	database *gorm.DB
	log      logger.Logger
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		database: db.GetDB(),
		log:      logger.NewLogger(cfg),
	}
}

// Create
func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (res *dto.CountryResponse, err error) {
	country := models.Country{Name: req.Name}
	country.CreatedBy = int(ctx.Value(constatns.UserIdKey).(float64))
	country.CreatedAt = time.Now().UTC()

	tx := s.database.WithContext(ctx).Begin()
	if err = tx.Create(&country).Error; err != nil {
		tx.Rollback()
		s.log.Error(logger.Potgres, logger.Insert, nil, err.Error())
		return nil, err
	}
	tx.Commit()
	res = &dto.CountryResponse{Name: country.Name}

	return res, nil
}

// Update
func (s *CountryService) Update(ctx context.Context, id int, req *dto.CreateUpdateCountryRequest) (res *dto.CountryResponse, err error) {
	updateMap := map[string]interface{}{
		"Name":        req.Name,
		"modified_by": &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constatns.UserIdKey).(float64))},
		"modified_at": &sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}

	tx := s.database.WithContext(ctx).Begin()

	if err = tx.Model(&models.Country{}).
		Where("id = ?", id).
		Updates(updateMap).
		Error; err != nil {
		tx.Rollback()
		s.log.Error(logger.Potgres, logger.Update, nil, err.Error())
		return nil, err
	}

	country := &models.Country{}
	if err = tx.Model(&models.Country{}).
		Where("id = ? AND deleted_by is null", id).
		First(&country).Error; err != nil {
		tx.Rollback()
		s.log.Error(logger.Potgres, logger.Select, nil, err.Error())
		return nil, err
	}
	tx.Commit()
	res = &dto.CountryResponse{Id: country.ID, Name: country.Name}

	return res, nil
}

// Delete
func (s *CountryService) Delete(ctx context.Context, id int) (err error) {
	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constatns.UserIdKey).(float64))},
		"deleted_at": &sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}
	tx := s.database.WithContext(ctx).Begin()

	if err = tx.Model(&models.Country{}).
		Where("id = ?", id).Updates(deleteMap).Error; err != nil {
		tx.Rollback()
		s.log.Error(logger.Potgres, logger.Delete, nil, err.Error())
		return err
	}
	tx.Commit()

	return nil
}

// Get by Id
func (s *CountryService) GetById(ctx context.Context, id int) (res *dto.CountryResponse, err error) {
	country := &models.Country{}

	tx := s.database.WithContext(ctx).Begin()
	if err = tx.Model(&models.Country{}).
		Where("id = ? AND deleted_by is null", id).
		First(&country).Error; err != nil {
		tx.Rollback()
		s.log.Error(logger.Potgres, logger.Select, nil, err.Error())
		return nil, err
	}
	res = &dto.CountryResponse{Name: country.Name, Id: country.ID}

	return res, nil
}
