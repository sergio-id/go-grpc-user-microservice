package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/domain"
	"time"

	"github.com/sergio-id/go-grpc-user-microservice/pkg/grpc_errors"
)

const (
	basePrefix = "cache_user:user_id:"
)

// RedisRepository is an interface for the redis repository.
type userRedisRepository struct {
	redisClient *redis.Client
	prefix      string
}

// NewUserRedisRepository creates a new redis repository.
func NewUserRedisRepository(redisClient *redis.Client) *userRedisRepository {
	return &userRedisRepository{redisClient: redisClient, prefix: basePrefix}
}

// GetByIDCtx method gets user by id.
func (r *userRedisRepository) GetByIDCtx(ctx context.Context, id uint64) (*domain.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRedisRepository.GetByIDCtx")
	defer span.Finish()

	userBytes, err := r.redisClient.Get(ctx, r.getFullKey(id)).Bytes()
	if err != nil {
		if err != redis.Nil {
			return nil, grpc_errors.ErrNotFound
		}
		return nil, err
	}

	var u domain.User
	if err = json.Unmarshal(userBytes, &u); err != nil {
		return nil, err
	}

	return &u, nil
}

// SetUserCtx method sets user.
func (r *userRedisRepository) SetUserCtx(ctx context.Context, duration time.Duration, user *domain.User) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRedisRepository.SetUserCtx")
	defer span.Finish()

	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.redisClient.
		Set(ctx, r.getFullKey(user.ID), userBytes, duration).
		Err()
}

// DeleteUserCtx method deletes user.
func (r *userRedisRepository) DeleteUserCtx(ctx context.Context, id uint64) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRedisRepository.DeleteUserCtx")
	defer span.Finish()

	return r.redisClient.
		Del(ctx, r.getFullKey(id)).
		Err()
}

// getFullKey method returns a full key.
func (r *userRedisRepository) getFullKey(id uint64) string {
	return fmt.Sprintf("%s:%d", r.prefix, id)
}
