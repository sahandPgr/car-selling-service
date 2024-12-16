package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/api/handlers"
	"github.com/sahandPgr/car-selling-service/config"
)

func Countries(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCountriesHandler(cfg)
	r.POST("/", h.Create)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.POST("/get-by-filter", h.GetByFilter)
}

func Cities(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCityHandler(cfg)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func Files(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewFileHandler(cfg)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
