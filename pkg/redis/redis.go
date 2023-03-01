package redis

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	defaultConnAttempts = 3
	defaultConnTimeout  = time.Second
)

type redisClient struct {
	connAttempts int
	connTimeout  time.Duration

	client *redis.Client
}

var _ RedisEngine = (*redisClient)(nil)

func NewRedisClient(cfgRedis Config) (RedisEngine, error) {
	options := &redis.Options{
		Addr:         cfgRedis.Addr,
		MinIdleConns: cfgRedis.MinIdleConns,
		PoolSize:     cfgRedis.PoolSize,
		PoolTimeout:  time.Duration(cfgRedis.PoolTimeout) * time.Second,
		Password:     cfgRedis.Password,
		DB:           cfgRedis.DB,
	}

	r := &redisClient{
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	for r.connAttempts > 0 {
		r.client = redis.NewClient(options)
		if r.client != nil {
			break
		}

		time.Sleep(r.connTimeout)
		r.connAttempts--
	}

	if r.client == nil {
		return nil, errors.New("can't connect to redis")
	}

	return r, nil
}

func (r *redisClient) GetRedisClient() *redis.Client {
	return r.client
}

func (r *redisClient) Close() {
	if r.client != nil {
		r.client.Close()
	}
}
