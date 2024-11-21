package handlers

import (
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
