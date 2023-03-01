package app

import (
	"github.com/go-redis/redis/v8"
	pkg_redis "github.com/sergio-id/go-grpc-user-microservice/pkg/redis"
)

// ConnectRedis connects to the Redis server.
func getRedisClient(cfgRedis pkg_redis.Config) (*redis.Client, func(), error) {
	r, err := pkg_redis.NewRedisClient(cfgRedis)
	if err != nil {
		return nil, nil, err
	}
	return r.GetRedisClient(), func() { r.Close() }, nil
}
