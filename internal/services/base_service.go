package services

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/constatns"
	"github.com/sahandPgr/car-selling-service/internal/db"
	"github.com/sahandPgr/car-selling-service/internal/models"
	serviceerrors "github.com/sahandPgr/car-selling-service/internal/services/service_errors"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
	"github.com/sahandPgr/car-selling-service/utils"
	"gorm.io/gorm"
)

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Log      logger.Logger
	Preloads []Preload
}

type Preload struct {
	string
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
	u, _ := utils.TypeConverter[models.BaseModel](model)
	return s.GetById(ctx, u.ID)
}

// Update
func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {
	updateMap, _ := utils.TypeConverter[map[string]interface{}](req)
	snakeMap := make(map[string]interface{})
	for k, v := range *updateMap {
		snakeMap[utils.ToSnakeCase(k)] = v
	}
	snakeMap["modified_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constatns.UserIdKey).(float64))}
	snakeMap["modified_at"] = &sql.NullTime{Valid: true, Time: time.Now().UTC()}

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

// GetByFilter
func (s *BaseService[T, Tc, Tu, Tr]) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[Tr], error) {
	return pagination[T, Tr](req, s.Preloads, s.Database)
}

// GetQuery
func getQuery[T any](filter *dto.DynaimcFilter) string {
	model := new(T)
	modelType := reflect.TypeOf(*model)
	query := make([]string, 0)
	query = append(query, "deleted_by is null")

	if filter.Filter != nil {
		for name, filter := range *&filter.Filter {
			fld, ok := modelType.FieldByName(name)
			if ok {
				fld.Name = utils.ToSnakeCase(fld.Name)
				switch filter.Type {

				case "contains":
					query = append(query, fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From))

				case "notContains":
					query = append(query, fmt.Sprintf("%s not ILike '%%%s%%'", fld.Name, filter.From))

				case "startsWith":
					query = append(query, fmt.Sprintf("%s ILike '%s%%'", fld.Name, filter.From))

				case "endsWith":
					query = append(query, fmt.Sprintf("%s ILike '%%%s'", fld.Name, filter.From))

				case "equals":
					query = append(query, fmt.Sprintf("%s = '%s'", fld.Name, filter.From))

				case "notEqual":
					query = append(query, fmt.Sprintf("%s != '%s'", fld.Name, filter.From))

				case "lessThan":
					query = append(query, fmt.Sprintf("%s < %s", fld.Name, filter.From))

				case "lessThanOrEqual":
					query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.From))

				case "greaterThan":
					query = append(query, fmt.Sprintf("%s > %s", fld.Name, filter.From))

				case "greaterThanOrEqual":
					query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))

				case "inRange":
					if fld.Type.Kind() == reflect.String {
						query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))
					} else {
						query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.To))
					}
				}
			}
		}
	}

	return strings.Join(query, " AND ")
}

// GetSort
func getSort[T any](filter *dto.DynaimcFilter) string {
	model := new(T)
	modelType := reflect.TypeOf(*model)
	sort := make([]string, 0)
	if filter.Sort != nil {
		for _, tp := range *filter.Sort {
			fld, ok := modelType.FieldByName(tp.ColId)
			if ok && (tp.Sort == "asc" || tp.Sort == "desc") {
				fld.Name = utils.ToSnakeCase(fld.Name)
				sort = append(sort, fmt.Sprintf("%s %s", fld.Name, tp.Sort))
			}
		}
	}
	return strings.Join(sort, ", ")
}

// Paginate
func pagination[T any, Tr any](req *dto.PaginationInputWithFilter, preloads []Preload, db *gorm.DB) (*dto.PagedList[Tr], error) {
	model := new(T)
	items := new([]T)
	itemResponse := new([]Tr)
	var totalRows int64 = 0
	var err error

	db = preload(db, preloads)
	query := getQuery[T](&req.DynaimcFilter)
	sort := getSort[T](&req.DynaimcFilter)

	db.Model(model).
		Where(query).
		Count(&totalRows)

	if err = db.Model(model).
		Where(query).
		Offset(req.GetOffset()).
		Limit(req.GetPageSize()).
		Order(sort).
		Find(&items).Error; err != nil {
		return nil, err
	}

	itemResponse, err = utils.TypeConverter[[]Tr](items)

	return getPageList(itemResponse, totalRows, req.PageNumber, req.PageSize), err
}

func getPageList[T any](items *[]T, totalRow int64, pageNumber int, pageSize int) *dto.PagedList[T] {
	pl := &dto.PagedList[T]{
		PageNumber: pageNumber,
		TotalRows:  totalRow,
		Items:      items,
	}
	pl.TotalPages = int(math.Ceil(float64(totalRow) / float64(pageSize)))
	pl.HasNextPage = pl.TotalPages > pageNumber
	pl.HasPerviousPage = pageNumber > 1

	return pl
}

// Preload
func preload(db *gorm.DB, preloads []Preload) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.string)
	}
	return db
}
