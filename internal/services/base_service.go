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

// NewBaseService initializes a new BaseService instance.
func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config, preloads []Preload) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: db.GetDB(),
		Log:      logger.NewLogger(cfg),
		Preloads: preloads,
	}
}

// Create inserts a new record in the database and returns the created entity.
func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, _ := utils.TypeConverter[T](req)   // Convert request to model type.
	tx := s.Database.WithContext(ctx).Begin() // Begin a new database transaction.
	if err := tx.Create(model).Error; err != nil {
		tx.Rollback() // Rollback transaction in case of error.
		s.Log.Error(logger.Potgres, logger.Insert, nil, err.Error())
		return nil, err
	}
	tx.Commit() // Commit the transaction.
	u, _ := utils.TypeConverter[models.BaseModel](model)
	return s.GetById(ctx, u.ID) // Retrieve the created entity by ID.
}

// Update modifies an existing record based on the provided ID and request data.
func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {
	updateMap, _ := utils.TypeConverter[map[string]interface{}](req) // Convert request to map.
	snakeMap := make(map[string]interface{})
	for k, v := range *updateMap {
		snakeMap[utils.ToSnakeCase(k)] = v // Convert keys to snake_case.
	}
	snakeMap["modified_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constatns.UserIdKey).(float64))}
	snakeMap["modified_at"] = &sql.NullTime{Valid: true, Time: time.Now().UTC()}

	model := new(T)
	tx := s.Database.WithContext(ctx).Begin() // Begin transaction.
	if err := tx.Model(model).
		Where(constatns.NotNullQuery, id).
		Updates(snakeMap).Error; err != nil {
		tx.Rollback()
		s.Log.Error(logger.Potgres, logger.Update, nil, err.Error())
		return nil, err
	}
	tx.Commit()               // Commit transaction.
	return s.GetById(ctx, id) // Retrieve the updated entity by ID.
}

// Delete performs a soft delete by setting the `deleted_by` and `deleted_at` fields.
func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	tx := s.Database.WithContext(ctx).Begin() // Begin transaction.
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
		Where(constatns.NotNullQuery, id).
		Updates(deleteMap).RowsAffected; count == 0 {
		s.Log.Error(logger.Potgres, logger.Update, nil, serviceerrors.RecordNotFound)
		tx.Rollback()
		return &serviceerrors.ServiceError{EndUserMessage: serviceerrors.RecordNotFound}
	}
	tx.Commit() // Commit transaction.
	return nil
}

// GetById retrieves a record by its ID if not deleted.
func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	tx := s.Database.WithContext(ctx).Begin() // Begin transaction.
	model := new(T)

	if err := tx.Model(model).
		Where(constatns.NotNullQuery, id).
		First(model).Error; err != nil {
		s.Log.Error(logger.Potgres, logger.Select, nil, err.Error())
		tx.Rollback()
		return nil, err
	}

	return utils.TypeConverter[Tr](model) // Convert model to response type.
}

// GetByFilter retrieves records based on filters with pagination.
func (s *BaseService[T, Tc, Tu, Tr]) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[Tr], error) {
	return pagination[T, Tr](req, s.Preloads, s.Database)
}

// getQuery generates the SQL WHERE clause based on dynamic filters.
func getQuery[T any](filter *dto.DynaimcFilter) string {
	model := new(T)
	modelType := reflect.TypeOf(*model)
	query := make([]string, 0)
	query = append(query, "deleted_by is null") // Ensure only non-deleted records are fetched.

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

	return strings.Join(query, " AND ") // Combine conditions with AND.
}

// getSort generates the SQL ORDER BY clause based on dynamic sorting options.
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
	return strings.Join(sort, ", ") // Combine sorting options with commas.
}

// pagination handles pagination, filtering, and sorting logic for queries.
func pagination[T any, Tr any](req *dto.PaginationInputWithFilter, preloads []Preload, db *gorm.DB) (*dto.PagedList[Tr], error) {
	model := new(T)
	items := new([]T)
	itemResponse := new([]Tr)
	var totalRows int64 = 0
	var err error

	db = preload(db, preloads)               // Apply preloads.
	query := getQuery[T](&req.DynaimcFilter) // Generate query conditions.
	sort := getSort[T](&req.DynaimcFilter)   // Generate sorting options.

	db.Model(model).
		Where(query).
		Count(&totalRows) // Count total matching records.

	if err = db.Model(model).
		Where(query).
		Offset(req.GetOffset()).
		Limit(req.GetPageSize()).
		Order(sort).
		Find(&items).Error; err != nil {
		return nil, err
	}

	itemResponse, err = utils.TypeConverter[[]Tr](items) // Convert items to response type.

	return getPageList(itemResponse, totalRows, req.GetPageNumber(), req.GetPageSize()), err
}

// getPageList creates a PagedList object with metadata for pagination.
func getPageList[T any](items *[]T, totalRow int64, pageNumber int, pageSize int) *dto.PagedList[T] {
	pl := &dto.PagedList[T]{
		PageNumber: pageNumber,
		TotalRows:  totalRow,
		Items:      items,
	}
	pl.TotalPages = int(math.Ceil(float64(totalRow) / float64(pageSize))) // Calculate total pages.
	pl.HasNextPage = pl.TotalPages > pageNumber                           // Check if next page exists.
	pl.HasPerviousPage = pageNumber > 1                                   // Check if previous page exists.

	return pl
}

// preload applies preload options to the database query.
func preload(db *gorm.DB, preloads []Preload) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.string)
	}
	return db
}
