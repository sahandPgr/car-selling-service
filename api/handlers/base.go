package handlers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/api/helper"
)

func Create[Ti any, To any](ctx *gin.Context, caller func(ctx context.Context, req *Ti) (*To, error)) {
	dt := new(Ti)
	err := ctx.ShouldBindJSON(dt)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	res, err := caller(ctx, dt)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Created, helper.GetBaseHttpResponse(res, true, helper.Created))
}

func Update[Ti any, To any](ctx *gin.Context, caller func(ctx context.Context, id int, req *Ti) (*To, error)) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(helper.BadRequest,
			helper.GetBaseHttpResponseWithError(nil, false, helper.BadRequest, err))
		return
	}
	dt := new(Ti)
	err = ctx.ShouldBindJSON(dt)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}
	res, err := caller(ctx, id, dt)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(res, true, helper.Success))
}

func Delete(ctx *gin.Context, caller func(ctx context.Context, id int) error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(helper.BadRequest,
			helper.GetBaseHttpResponseWithError(nil, false, helper.BadRequest, err))
		return
	}
	if id == 0 {
		ctx.AbortWithStatusJSON(helper.NotFound, helper.GetBaseHttpResponse(nil, false, helper.NotFound))
		return
	}

	err = caller(ctx, id)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(nil, true, helper.Success))
}

func GetById[To any](ctx *gin.Context, caller func(ctx context.Context, id int) (*To, error)) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(helper.BadRequest,
			helper.GetBaseHttpResponseWithError(nil, false, helper.BadRequest, err))
		return
	}
	if id == 0 {
		ctx.AbortWithStatusJSON(helper.NotFound, helper.GetBaseHttpResponse(nil, false, helper.NotFound))
		return
	}

	res, err := caller(ctx, id)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(res, true, helper.Success))
}

func GetByFilter[Ti any, To any](ctx *gin.Context, caller func(ctx context.Context, req *Ti) (*To, error)) {
	dt := new(Ti)
	err := ctx.ShouldBindJSON(dt)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	res, err := caller(ctx, dt)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
		return
	}

	ctx.JSON(helper.Success, helper.GetBaseHttpResponse(res, true, helper.Success))
}
