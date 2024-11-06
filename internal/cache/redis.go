package internal

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sahandPgr/car-selling-service/config"
)

var redisClient *redis.Client

// InitialRedisClient initializes the Redis client
func InitialRedisClient(config *config.Config) {

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

}

// GetRedisClient returns the Redis client
func GetRedisClient() *redis.Client {
	return redisClient
}

// CloseRedisClient closes the Redis client
func CloseRedisClient() {
	redisClient.Close()
}
