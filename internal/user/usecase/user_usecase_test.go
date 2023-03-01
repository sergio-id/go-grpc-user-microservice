package usecase

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/domain"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/mock"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/types"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/logger"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserUseCase_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	appLogger := logger.NewAppLogger(logger.Config{})
	appLogger.InitLogger()

	mockPostgresqlRepository := mock.NewMockPostgresqlRepository(ctrl)
	mockRedisRepository := mock.NewMockRedisRepository(ctrl)
	userUseCase := NewUserUseCase(appLogger, mockPostgresqlRepository, mockRedisRepository)

	mockUser := &domain.User{
		ID:       1,
		Email:    "example@gmail.com",
		Password: "111111",
		Gender:   string(types.Male),
		Status:   string(types.Active),
	}

	ctx := context.Background()

	mockPostgresqlRepository.EXPECT().GetByEmail(gomock.Any(), mockUser.Email).Return(nil, sql.ErrNoRows)

	mockPostgresqlRepository.EXPECT().Create(gomock.Any(), mockUser).Return(&domain.User{
		ID:       mockUser.ID,
		Email:    "example@gmail.com",
		Password: "111111",
		Gender:   string(types.Male),
		Status:   string(types.Active),
	}, nil)

	createdUser, err := userUseCase.Create(ctx, mockUser)
	require.NoError(t, err)
	require.NotNil(t, createdUser)
	require.Equal(t, createdUser.ID, mockUser.ID)
}

func TestUserUseCase_GetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	appLogger := logger.NewAppLogger(logger.Config{})
	appLogger.InitLogger()

	mockPostgresqlRepository := mock.NewMockPostgresqlRepository(ctrl)
	mockRedisRepository := mock.NewMockRedisRepository(ctrl)
	userUseCase := NewUserUseCase(appLogger, mockPostgresqlRepository, mockRedisRepository)

	mockUser := &domain.User{
		ID:       1,
		Email:    "example@gmail.com",
		Password: "111111",
		Gender:   string(types.Male),
		Status:   string(types.Active),
	}

	ctx := context.Background()

	mockRedisRepository.EXPECT().GetByIDCtx(gomock.Any(), mockUser.ID).Return(nil, redis.Nil)
	mockPostgresqlRepository.EXPECT().GetById(gomock.Any(), mockUser.ID).Return(mockUser, nil)
	mockRedisRepository.EXPECT().SetUserCtx(gomock.Any(), gomock.Any(), mockUser).Return(nil)

	foundUser, err := userUseCase.GetById(ctx, mockUser.ID)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.ID, mockUser.ID)
}
