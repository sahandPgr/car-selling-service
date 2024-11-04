package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/config"
)

func InitServer() {
	config := config.GetConfig()
	r := gin.Default()
	r.Run(fmt.Sprintf(":%s", config.Server.Port))
}
