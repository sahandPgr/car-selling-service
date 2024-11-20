package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/api/handlers"
	"github.com/sahandPgr/car-selling-service/config"
)

// This function creates a new user router
func User(r *gin.RouterGroup, config *config.Config) {
	h := handlers.NewUserHandler(config)
	r.POST("/send-otp", h.SendOtp)
}
