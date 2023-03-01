package grpc_client

import (
	"context"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/interceptors"
	"github.com/sergio-id/go-grpc-user-microservice/proto"
)

// NewUserGrpcClient creates new user grpc client
func NewUserGrpcClient(ctx context.Context, port string, im interceptors.InterceptorManager) (proto.UserServiceClient, func() error, error) {
	// create grpc connection
	grpcServiceConn, err := NewGrpcConn(ctx, port, im)
	if err != nil {
		return nil, nil, err
	}
	// create user grpc client
	userGrpcClient := proto.NewUserServiceClient(grpcServiceConn)

	return userGrpcClient, grpcServiceConn.Close, nil
}
