package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/api/middlewares"
	"github.com/sahandPgr/car-selling-service/config"
)

func InitServer(config *config.Config) {
	r := gin.Default()
	r.Use(middlewares.Cors(config))
	r.Run(fmt.Sprintf(":%s", config.Server.Port))
}
