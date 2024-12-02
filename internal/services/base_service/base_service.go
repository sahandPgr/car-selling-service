package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/constatns"
	"github.com/sahandPgr/car-selling-service/internal/db"
	serviceerrors "github.com/sahandPgr/car-selling-service/internal/services/service_errors"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
	"github.com/sahandPgr/car-selling-service/utils"
	"gorm.io/gorm"
)

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Log      logger.Logger
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: db.GetDB(),
		Log:      logger.NewLogger(cfg),
	}
}

// Create
func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, _ := utils.TypeConverter[T](req)
	tx := s.Database.WithContext(ctx).Begin()
	if err := tx.Create(model).Error; err != nil {
		tx.Rollback()
		s.Log.Error(logger.Potgres, logger.Insert, nil, err.Error())
		return nil, err
	}
	tx.Commit()
	return utils.TypeConverter[Tr](model)
}

// Update
func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {
	updateMap, _ := utils.TypeConverter[map[string]interface{}](req)
	(*updateMap)["modified_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constatns.UserIdKey).(float64))}
	(*updateMap)["modified_at"] = &sql.NullTime{Valid: true, Time: time.Now().UTC()}

	model := new(T)
	tx := s.Database.WithContext(ctx).Begin()
	if err := tx.Model(model).
		Where("id = ? AND deleted_by is null", id).
		Updates(*updateMap).Error; err != nil {
		tx.Rollback()
		s.Log.Error(logger.Potgres, logger.Update, nil, err.Error())
		return nil, err
	}
	tx.Commit()
	return s.GetById(ctx, id)
}

// Delete
func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	tx := s.Database.WithContext(ctx).Begin()
	model := new(T)
	userId := ctx.Value(constatns.UserIdKey)
	if userId == nil {
		return &serviceerrors.ServiceError{EndUserMessage: serviceerrors.PermissionDenied}
	}
	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Valid: true, Int64: int64(userId.(float64))},
		"deleted_at": &sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}
	if count := tx.Model(model).
		Where("id = ? AND deleted_by is null", id).
		Updates(deleteMap).RowsAffected; count == 0 {
		s.Log.Error(logger.Potgres, logger.Update, nil, serviceerrors.RecordNotFound)
		tx.Rollback()
		return &serviceerrors.ServiceError{EndUserMessage: serviceerrors.RecordNotFound}
	}
	tx.Commit()
	return nil
}

// GetById
func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	tx := s.Database.WithContext(ctx).Begin()
	model := new(T)

	if err := tx.Model(model).
		Where("id = ? AND deleted_by is null", id).
		First(model).Error; err != nil {
		s.Log.Error(logger.Potgres, logger.Select, nil, err.Error())
		tx.Rollback()
		return nil, err
	}

	return utils.TypeConverter[Tr](model)
}
