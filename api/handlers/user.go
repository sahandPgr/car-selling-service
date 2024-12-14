package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/services"
)

// UserHandler is a struct for user handler
type UserHandler struct {
	userService *services.UserService
}

// This function creates a new user handler
func NewUserHandler(config *config.Config) *UserHandler {
	return &UserHandler{
		userService: services.NewUserService(config),
	}
}

// This function sends otp Handler
// SendOtp godoc
// @Summary Send otp
// @Description Send otp
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/user/send-otp [post]
func (h *UserHandler) SendOtp(c *gin.Context) {
	var dto = new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(dto)
	if err != nil {
		c.AbortWithStatusJSON(int(helper.BadRequest),
			helper.GetBaseHttpResponseWithValidation(nil, false, helper.BadRequest, err))
		return
	}

	err = h.userService.SendOtp(dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		c.AbortWithStatusJSON(statusCode,
			helper.GetBaseHttpResponseWithError(nil, false, helper.Conflict, err))
		return
	}

	c.JSON(int(helper.Created), helper.GetBaseHttpResponse(nil, true, helper.Created))
}

// This function Register By Username Handler
// RegisterByUsername godoc
// @Summary RegisterByUsername
// @Description RegisterByUsername
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.RegisterUserByUsernameRequest true "RegisterUserByUsernameRequest"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/user/register-by-username [post]
func (h *UserHandler) RegisterByUsername(c *gin.Context) {
	var dto = new(dto.RegisterUserByUsernameRequest)
	err := c.ShouldBindJSON(dto)
	if err != nil {
		c.AbortWithStatusJSON(int(helper.BadRequest), helper.GetBaseHttpResponseWithValidation(nil, false, helper.BadRequest, err))
		return
	}
	err = h.userService.RegisterByUsername(dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		c.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, helper.Conflict, err))
		return
	}
	c.JSON(int(helper.Created), helper.GetBaseHttpResponse(nil, true, helper.Created))
}

// This function Register/Login By Phone number Handler
// RegisterLoginByMobile godoc
// @Summary RegisterLoginByMobile
// @Description RegisterLoginByMobile
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.RegisterLoginByMobileRequest true "RegisterLoginByMobileRequest"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/user/register-login-by-phone [post]
func (h *UserHandler) RegisterLoginByMobile(c *gin.Context) {
	var dto = new(dto.RegisterLoginByMobileRequest)
	err := c.ShouldBindJSON(dto)
	if err != nil {
		c.AbortWithStatusJSON(int(helper.BadRequest), helper.GetBaseHttpResponseWithValidation(nil, false, helper.BadRequest, err))
		return
	}
	res, err := h.userService.RegisterLoginByMobile(dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		c.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, helper.InternalServerError, err))
		return
	}
	c.JSON(int(helper.Created), helper.GetBaseHttpResponse(res, true, helper.Created))
}

// This function Login By Username Handler
// LoginByUsername godoc
// @Summary LoginByUsername
// @Description LoginByUsername
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.LoginByUsernameRequest true "LoginByUsernameRequest"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/user/login-by-username [post]
func (h *UserHandler) LoginByUsername(c *gin.Context) {
	var dto = new(dto.LoginByUsernameRequest)
	err := c.ShouldBindJSON(dto)
	if err != nil {
		c.AbortWithStatusJSON(int(helper.BadRequest), helper.GetBaseHttpResponseWithValidation(nil, false, helper.BadRequest, err))
		return
	}
	res, err := h.userService.LoginByUsername(dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		c.AbortWithStatusJSON(statusCode, helper.GetBaseHttpResponseWithError(nil, false, helper.Conflict, err))
		return
	}
	c.JSON(int(helper.Created), helper.GetBaseHttpResponse(res, true, helper.Created))
}
