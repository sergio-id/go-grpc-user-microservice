package usecase

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/domain"
	"time"

	"github.com/pkg/errors"

	"github.com/sergio-id/go-grpc-user-microservice/pkg/grpc_errors"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/logger"
)

const (
	baseUserCacheDuration = 1 * time.Hour // 1 hour
)

// UserUseCase is an interface for the user usecase.
type UserUseCase struct {
	logger       logger.Logger
	postgresRepo domain.PostgresqlRepository
	redisRepo    domain.RedisRepository
}

// NewUserUseCase creates a new user usecase.
func NewUserUseCase(logger logger.Logger, postgresRepo domain.PostgresqlRepository, redisRepo domain.RedisRepository) *UserUseCase {
	return &UserUseCase{logger: logger, postgresRepo: postgresRepo, redisRepo: redisRepo}
}

// Create method creates a new user.
func (u *UserUseCase) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.Create")
	defer span.Finish()

	existsUser, err := u.postgresRepo.GetByEmail(ctx, user.Email)
	if existsUser != nil || err == nil {
		return nil, grpc_errors.ErrEmailExists
	}

	return u.postgresRepo.Create(ctx, user)
}

// Update method updates user.
func (u *UserUseCase) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.Update")
	defer span.Finish()

	updatedUser, err := u.postgresRepo.Update(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "postgresRepo.Update")
	}

	if err := u.redisRepo.DeleteUserCtx(ctx, updatedUser.ID); err != nil {
		u.logger.Errorf("UserRedisRepository.DeleteUserCtx", err)
	}

	return updatedUser, nil
}

// Delete method deletes user.
func (u *UserUseCase) Delete(ctx context.Context, id uint64) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.Delete")
	defer span.Finish()

	err := u.postgresRepo.Delete(ctx, id)
	if err != nil {
		return errors.Wrap(err, "postgresRepo.Delete")
	}

	return u.redisRepo.DeleteUserCtx(ctx, id) //todo here
}

// GetById method returns user by id.
func (u *UserUseCase) GetById(ctx context.Context, id uint64) (*domain.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.GetById")
	defer span.Finish()

	cachedUser, err := u.redisRepo.GetByIDCtx(ctx, id)
	if err != nil && !errors.Is(err, redis.Nil) {
		u.logger.Errorf("redisRepo.GetByIDCtx", err)
	}
	if cachedUser != nil {
		return cachedUser, nil
	}

	foundUser, err := u.postgresRepo.GetById(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "postgresRepo.GetById")
	}

	if err := u.redisRepo.SetUserCtx(ctx, baseUserCacheDuration, foundUser); err != nil {
		u.logger.Errorf("UserRedisRepository.SetUserCtx", err)
	}

	return foundUser, nil
}
