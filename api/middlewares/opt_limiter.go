package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/pkg/limiter"
	"golang.org/x/time/rate"
)

// OptLimiter is a middleware that limits the number of requests for each IP address.
func OptLimiter(config *config.Config) gin.HandlerFunc {
	var optLimiter = limiter.NewIPRateLimiter(rate.Every(config.Otp.Limiter*time.Second), 1)
	return func(ctx *gin.Context) {
		limiter := optLimiter.GetLimiter(ctx.Request.RemoteAddr)
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests,
				helper.GetBaseHttpResponseWithError(nil, false, helper.TooManyRequests, errors.New("Not Allowed")))
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}
