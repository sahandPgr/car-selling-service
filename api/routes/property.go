package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/api/handlers"
	"github.com/sahandPgr/car-selling-service/config"
)

func PropertyCategory(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewPropertiesCategoryHandler(cfg)
	r.POST("/", h.Create)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.POST("/get-by-filter", h.GetByFilter)
}

func Property(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewPropertiesHandler(cfg)
	r.POST("/", h.Create)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.POST("/get-by-filter", h.GetByFilter)
}
