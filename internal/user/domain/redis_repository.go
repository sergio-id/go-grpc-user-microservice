//go:generate mockgen -source redis_repository.go -destination ../mock/redis_repository.go -package mock
package domain

import (
	"context"
	"time"
)

// RedisRepository is an interface for the redis repository.
type RedisRepository interface {
	GetByIDCtx(ctx context.Context, id uint64) (*User, error)
	SetUserCtx(ctx context.Context, duration time.Duration, user *User) error
	DeleteUserCtx(ctx context.Context, id uint64) error
}
