package redis

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// Defaults.
const (
  PoolSize = 10
	MaxActiveConns = 5
	MinIdleConns = 2

	PoolTimeoutSec = 30
)


func NewConnectionFromEnvVar() (*redis.Client, error) {
  address := os.Getenv("REDIS_ADDRESS")
  password := os.Getenv("REDIS_PASSWORD")
  db, err := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))

  if err != nil {
    return nil, fmt.Errorf("Redis DB number error: %s", err.Error())
  }

  pool := getRedisConnectionPool(address, password, db)

  ctx := context.Background()
  _, err = pool.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Redis: %s", err.Error())
	}

  return pool, nil
}

func getRedisConnectionPool(address, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:           address, // Redis server address
		Password:       password,               // No password by default
		DB:             db,                // Use default DB
		PoolSize:       PoolSize,               // Maximum number of connections
    MaxActiveConns: MaxActiveConns,
		MinIdleConns:   MinIdleConns,               // Minimum number of idle connections
		PoolTimeout:    time.Duration(PoolTimeoutSec) * time.Second, // Timeout for getting a connection from the pool
	})
}
