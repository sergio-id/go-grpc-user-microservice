package grpc

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/types"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/grpc_client"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/interceptors"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/logger"
	"github.com/sergio-id/go-grpc-user-microservice/proto"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func SetupTestUserGrpcClient() (proto.UserServiceClient, func() error, error) {
	appLogger := logger.NewAppLogger(logger.Config{
		LogLevel: "debug",
		Console:  true,
	})
	interceptorManager := interceptors.NewInterceptorManager(appLogger, nil)
	return grpc_client.NewUserGrpcClient(context.Background(), ":5001", interceptorManager)
}

func Test_userService(t *testing.T) {
	t.Parallel()

	client, closeFunc, err := SetupTestUserGrpcClient()
	require.NoError(t, err)
	defer closeFunc() // nolint: errcheck

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// check create method
	createRequest := &proto.CreateRequest{
		Email:     gofakeit.Email(),
		Password:  gofakeit.Password(true, true, true, true, false, 16),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Gender:    string(types.Male),
		Status:    string(types.Active),
	}
	t.Logf("createRequest: %s", createRequest.String())

	createReply, err := client.Create(ctx, createRequest)

	checkCreatedUser(t, createReply, err)

	t.Logf("createReply: %s", createReply.String())

	// check get method
	getByIDReply, err := client.GetByID(ctx, &proto.GetByIDRequest{
		Id: createReply.GetUser().GetId(),
	})

	checkGetByIDUser(t, createRequest, getByIDReply, err)

	t.Logf("getByIDReply: %s", getByIDReply.String())

	// check update method
	updateRequest := &proto.UpdateRequest{
		Id:        createReply.GetUser().GetId(),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Gender:    string(types.Male),
		Status:    string(types.Active),
	}
	t.Logf("updateRequest: %s", updateRequest.String())

	updateReply, err := client.Update(ctx, updateRequest)

	checkUpdatedUser(t, updateRequest, updateReply, err)

	t.Logf("updateReply: %s", updateReply.String())

	// check delete method
	deleteReply, err := client.Delete(ctx, &proto.DeleteRequest{
		Id: createReply.GetUser().GetId(),
	})

	require.NoError(t, err)
	require.NotNil(t, deleteReply)

	t.Logf("deleteReply: %s", deleteReply.String())

	// check get method
	getByIDReply, err = client.GetByID(ctx, &proto.GetByIDRequest{
		Id: createReply.GetUser().GetId(),
	})

	require.Error(t, err)
	require.Nil(t, getByIDReply)

	t.Logf("getByIDReply: %s", getByIDReply.String())
}

func checkUpdatedUser(t *testing.T, request *proto.UpdateRequest, reply *proto.UpdateReply, err error) {
	require.NoError(t, err)
	require.NotNil(t, reply)
	require.NotNil(t, reply.GetUser())
	require.Equal(t, request.GetFirstName(), reply.GetUser().GetFirstName())
	require.Equal(t, request.GetLastName(), reply.GetUser().GetLastName())
}

func checkGetByIDUser(t *testing.T, request *proto.CreateRequest, reply *proto.GetByIDReply, err error) {
	require.NoError(t, err)
	require.NotNil(t, reply)
	require.NotNil(t, reply.GetUser())
	require.Equal(t, request.GetEmail(), reply.GetUser().GetEmail())
	require.Equal(t, request.GetFirstName(), reply.GetUser().GetFirstName())
	require.Equal(t, request.GetLastName(), reply.GetUser().GetLastName())
}

func checkCreatedUser(t *testing.T, reply *proto.CreateReply, err error) {
	require.NoError(t, err)
	require.NotNil(t, reply)
	require.NotNil(t, reply.GetUser())
	require.NotEqual(t, reply.GetUser().GetId(), 0)
}
