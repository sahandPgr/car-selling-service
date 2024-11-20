package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	services "github.com/sahandPgr/car-selling-service/internal/services/user_service"
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
func (h *UserHandler) SendOtp(c *gin.Context) {
	var dto = new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(dto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GetBaseHttpResponseWithValidation(nil, false, http.StatusBadRequest, err))
		return
	}

	err = h.userService.SendOtp(dto)
	if err != nil {
		statusCode := helper.ConvertServiceErrorToStatusCode(err)
		c.AbortWithStatusJSON(statusCode,
			helper.GetBaseHttpResponseWithError(nil, false, statusCode, err))
	}

	c.JSON(http.StatusOK, helper.GetBaseHttpResponse(nil, true, http.StatusOK))
}
