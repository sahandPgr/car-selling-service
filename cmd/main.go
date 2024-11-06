package main

import (
	"github.com/sahandPgr/car-selling-service/api"
	"github.com/sahandPgr/car-selling-service/config"
	internal "github.com/sahandPgr/car-selling-service/internal/cache"
)

func main() {
	config := config.GetConfig()
	internal.InitialRedisClient(config)
	defer internal.CloseRedisClient()
	api.InitServer(config)
}
