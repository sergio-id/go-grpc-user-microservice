package redis

import (
	"github.com/go-redis/redis/v8"
)

// RedisEngine is an interface for Redis client.
type RedisEngine interface {
	GetRedisClient() *redis.Client
	Close()
}
