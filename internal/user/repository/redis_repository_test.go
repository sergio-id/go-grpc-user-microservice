package repository

import (
	"context"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/domain"
	"log"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
)

func SetupRedis() *userRedisRepository {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return NewUserRedisRepository(client)
}

func TestUserRedisRepo_SetUserCtx(t *testing.T) {
	t.Parallel()

	redisRepo := SetupRedis()

	t.Run("SetUserCtx", func(t *testing.T) {
		err := redisRepo.SetUserCtx(context.Background(), 10*time.Second, &domain.User{ID: 1})
		require.NoError(t, err)
	})
}

func TestUserRedisRepo_GetByIDCtx(t *testing.T) {
	t.Parallel()

	redisRepo := SetupRedis()

	t.Run("GetByIDCtx", func(t *testing.T) {
		user := &domain.User{ID: 1}

		err := redisRepo.SetUserCtx(context.Background(), 10*time.Second, user)
		require.NoError(t, err)

		user, err = redisRepo.GetByIDCtx(context.Background(), user.ID)
		require.NoError(t, err)
		require.NotNil(t, user)
	})
}

func TestUserRedisRepo_DeleteUserCtx(t *testing.T) {
	t.Parallel()

	redisRepo := SetupRedis()

	t.Run("DeleteUserCtx", func(t *testing.T) {
		user := &domain.User{ID: 1}

		err := redisRepo.DeleteUserCtx(context.Background(), user.ID)
		require.NoError(t, err)
	})
}
