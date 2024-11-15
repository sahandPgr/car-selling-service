package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/config"
)

// Cors function is a middleware for cors
func Cors(config *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", config.Cors.AllowOrigins)
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		ctx.Header("Access-Control-Max-Age", "3600")
		ctx.Set("Content-Type", "application/json")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
