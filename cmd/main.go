package main

import (
	"github.com/sahandPgr/car-selling-service/api"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/cache"
	"github.com/sahandPgr/car-selling-service/internal/db"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
)

func main() {
	config := config.GetConfig()
	log := logger.NewLogger(config)
	cache.InitialRedisClient(config, log)
	defer cache.CloseRedisClient()
	db.InitialDB(config, log)
	defer db.CloseDB()
	api.InitServer(config, log)
}
