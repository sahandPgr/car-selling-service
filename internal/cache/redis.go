package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

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

// Set function sets a key-value pair in Redis
func Set[T any](c *redis.Client, key string, value T, expiration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = c.Set(ctx, key, v, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// Get function gets a value from Redis
func Get[T any](c *redis.Client, key string) (T, error) {
	var res = *new(T)
	v, err := c.Get(ctx, key).Result()
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(v), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetRedisClient returns the Redis client
func GetRedisClient() *redis.Client {
	return redisClient
}

// CloseRedisClient closes the Redis client
func CloseRedisClient() {
	redisClient.Close()
}
