package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/api/handlers"
	"github.com/sahandPgr/car-selling-service/config"
)

func CarTypes(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarTypesHandler(cfg)
	r.POST("/", h.Create)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.POST("/get-by-filter", h.GetByFilter)
}

func Gearbox(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewGearboxesHandler(cfg)
	r.POST("/", h.Create)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.POST("/get-by-filter", h.GetByFilter)
}

func CarModels(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelsHandler(cfg)
	r.POST("/", h.Create)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.POST("/get-by-filter", h.GetByFilter)
}
