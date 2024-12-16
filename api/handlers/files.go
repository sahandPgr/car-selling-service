package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/constatns"
	"github.com/sahandPgr/car-selling-service/internal/services"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
)

var log = logger.NewLogger(config.GetConfig())

type FileHandler struct {
	service *services.FileService
}

func NewFileHandler(cfg *config.Config) *FileHandler {
	return &FileHandler{
		service: services.NewFileService(cfg),
	}
}

// CreateFile godoc
// @Summary Create a File
// @Description Create a File
// @Tags Files
// @Accept x-www-form-urlencoded
// @Produce json
// @Param file formData dto.UploadFileRequest true "Create a files"
// @Param file formData file true "Create a files"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.FileResponse} "Created"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/files/ [post]
// @Security AuthBearer
func (h *FileHandler) Create(ctx *gin.Context) {
	dt := new(dto.UploadFileRequest)
	err := ctx.ShouldBind(&dt)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.BadRequest, helper.GetBaseHttpResponseWithValidation(nil, false, helper.BadRequest, err))
		return
	}
	req := new(dto.CreateFileRequest)
	req.Description = dt.Description
	req.Directory = constatns.UploadDirName
	req.MediaType = dt.File.Header.Get("Content-Type")
	req.Name, err = prepareUpload(dt.File, req.Directory)
	if err != nil {
		log.Error(logger.Io, logger.UploadFile, nil, err.Error())
		ctx.AbortWithStatusJSON(helper.NotFound, helper.GetBaseHttpResponseWithError(nil, false, helper.NotFound, err))
		return
	}

	res, err := h.service.Create(ctx, req)
	if err != nil {
		stc := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(stc, helper.GetBaseHttpResponseWithError(nil, false, stc, err))
		return
	}

	ctx.JSON(helper.Created, helper.GetBaseHttpResponse(res, true, helper.Created))
}

// UpdateFile godoc
// @Summary Update a File
// @Description Update a File
// @Tags Files
// @Accept json
// @Produce json
// @Param id path int true "Update a files"
// @Param request body dto.UpdateFileRequest true "Update a files"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.FileResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/files/{id} [put]
// @Security AuthBearer
func (h *FileHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	req := new(dto.UpdateFileRequest)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.BadRequest, helper.GetBaseHttpResponseWithValidation(nil, false, helper.BadRequest, err))
		return
	}
	res, err := h.service.Update(ctx, id, req)
	if err != nil {
		stc := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(stc, helper.GetBaseHttpResponseWithError(nil, false, stc, err))
		return
	}
	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(res, true, helper.Success))
}

// DeleteFile godoc
// @Summary Delete a File
// @Description Delete a File
// @Tags Files
// @Accept json
// @Produce json
// @Param id path int true "Delete a files"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/files/{id} [delete]
// @Security AuthBearer
func (h *FileHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if id == 0 {
		ctx.AbortWithStatusJSON(helper.NotFound, helper.GetBaseHttpResponse(nil, false, helper.NotFound))
		return
	}
	file, _ := h.service.GetById(ctx, id)
	if err := os.Remove(fmt.Sprintf("%s/%s", file.Directory, file.Name)); err != nil {
		log.Error(logger.Io, logger.RmoveFile, nil, err.Error())
		ctx.AbortWithStatusJSON(helper.NotFound, helper.GetBaseHttpResponseWithError(nil, false, helper.NotFound, err))
		return
	}

	err := h.service.Delete(ctx, id)
	if err != nil {
		stc := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(stc, helper.GetBaseHttpResponseWithError(nil, false, stc, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(nil, true, helper.Success))
}

// GetByIdFile godoc
// @Summary GetById a File
// @Description GetById a File
// @Tags Files
// @Accept json
// @Produce json
// @Param id path int true "GetById a files"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.FileResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/files/{id} [get]
// @Security AuthBearer
func (h *FileHandler) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(helper.NotFound, helper.GetBaseHttpResponse(nil, false, helper.NotFound))
		return
	}

	res, err := h.service.GetById(ctx, id)
	if err != nil {
		stc := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(stc, helper.GetBaseHttpResponseWithError(nil, false, stc, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(res, true, helper.Success))
}

// GetFiles godoc
// @Summary Get Files
// @Description Get Files
// @Tags Files
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.FileResponse]} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/files/get-by-filter [post]
// @Security AuthBearer
func (h *FileHandler) GetByFilter(ctx *gin.Context) {
	req := new(dto.PaginationInputWithFilter)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.BadRequest, helper.GetBaseHttpResponseWithValidation(nil, false, helper.BadRequest, err))
		return
	}

	res, err := h.service.GetByFilter(ctx, req)
	if err != nil {
		stc := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(stc, helper.GetBaseHttpResponseWithError(nil, false, stc, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(res, true, helper.Success))
}

func prepareUpload(file *multipart.FileHeader, dir string) (string, error) {
	randFileName := uuid.New()
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	fileName := file.Filename
	fileArr := strings.Split(fileName, ".")
	fileExt := fileArr[len(fileArr)-1]
	fileName = fmt.Sprintf("%s.%s", randFileName, fileExt)
	dst := fmt.Sprintf("%s/%s", dir, fileName)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	if _, err = io.Copy(out, src); err != nil {
		return "", err
	}

	return fileName, nil
}
