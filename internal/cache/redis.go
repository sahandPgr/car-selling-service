package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
)

var redisClient *redis.Client

// InitialRedisClient initializes the Redis client
func InitialRedisClient(config *config.Config, log logger.Logger) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         config.Redis.Host + ":" + config.Redis.Port,
		Password:     config.Redis.Password,
		DB:           config.Redis.Db,
		PoolSize:     config.Redis.PoolSize,
		DialTimeout:  config.Redis.DialTimeout * time.Second,
		WriteTimeout: config.Redis.WriteTimeout * time.Second,
		ReadTimeout:  config.Redis.ReadTimeout * time.Second,
		PoolTimeout:  config.Redis.PoolTimeout * time.Second,
	})

	ctx := context.Background()
	if res := redisClient.Ping(ctx).String(); res != "ping: PONG" {
		log.Error(logger.Redis, logger.Startup, nil, "Failed to connect to Redis")
	}
}

// GetRedisClient returns the Redis client
func GetRedisClient() *redis.Client {
	return redisClient
}

// CloseRedisClient closes the Redis client
func CloseRedisClient() {
	redisClient.Close()
}
