package main

import (
	"github.com/sahandPgr/car-selling-service/api"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/cache"
	"github.com/sahandPgr/car-selling-service/internal/db"
	"github.com/sahandPgr/car-selling-service/internal/db/migrations"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	config := config.GetConfig()
	log := logger.NewLogger(config)
	cache.InitialRedisClient(config, log)
	defer cache.CloseRedisClient()
	db.InitialDB(config, log)
	defer db.CloseDB()
	migrations.Up_1(log)
	api.InitServer(config, log)
}
