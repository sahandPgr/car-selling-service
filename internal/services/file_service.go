package services

import (
	"context"

	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/models"
)

type FileService struct {
	// Base service to handle CRUD operations for files.
	base *BaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.FileResponse]
}

// NewFileService creates a new instance of FileService with the given config and preloads.
func NewFileService(cfg *config.Config) *FileService {
	return &FileService{
		base: NewBaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.FileResponse](cfg, []Preload{}),
	}
}

// Create a new file record.
func (s *FileService) Create(ctx context.Context, req *dto.CreateFileRequest) (res *dto.FileResponse, err error) {
	return s.base.Create(ctx, req)
}

// Update an existing file record by ID.
func (s *FileService) Update(ctx context.Context, id int, req *dto.UpdateFileRequest) (res *dto.FileResponse, err error) {
	return s.base.Update(ctx, id, req)
}

// Delete a file record by ID.
func (s *FileService) Delete(ctx context.Context, id int) (err error) {
	return s.base.Delete(ctx, id)
}

// GetById retrieves a file record by ID.
func (s *FileService) GetById(ctx context.Context, id int) (res *dto.FileResponse, err error) {
	return s.base.GetById(ctx, id)
}

// GetByFilter retrieves file records based on filter and pagination input.
func (s *FileService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (res *dto.PagedList[dto.FileResponse], err error) {
	return s.base.GetByFilter(ctx, req)
}
