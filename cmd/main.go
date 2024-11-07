package main

import (
	"github.com/sahandPgr/car-selling-service/api"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/cache"
	"github.com/sahandPgr/car-selling-service/internal/db"
)

func main() {
	config := config.GetConfig()
	cache.InitialRedisClient(config)
	defer cache.CloseRedisClient()
	db.InitialDB(config)
	defer db.CloseDB()
	api.InitServer(config)
}
