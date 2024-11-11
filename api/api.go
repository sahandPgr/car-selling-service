package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/api/middlewares"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
)

func InitServer(config *config.Config, log logger.Logger) {
	r := gin.Default()
	r.Use(middlewares.Cors(config))
	err := r.Run(fmt.Sprintf(":%s", config.Server.Port))
	if err != nil {
		log.Fatal(logger.General, logger.Startup, nil, "Failed to start server")
	}
}
