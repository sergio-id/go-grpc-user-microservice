package grpc_client

import (
	"context"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

const (
	backoffLinear  = 100 * time.Millisecond
	backoffRetries = 3
)

// NewGrpcConn creates new grpc connection
func NewGrpcConn(ctx context.Context, port string, im interceptors.InterceptorManager) (*grpc.ClientConn, error) {
	// retry options
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(backoffLinear)), // backoff
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),             // retry on codes
		grpc_retry.WithMax(backoffRetries),                              // max retries
	}

	// create grpc connection
	conn, err := grpc.DialContext(
		ctx,
		port,
		grpc.WithUnaryInterceptor(im.ClientRequestLoggerInterceptor()),        // log requests
		grpc.WithTransportCredentials(insecure.NewCredentials()),              // use insecure credentials
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)), // retry
	)
	if err != nil {
		return nil, errors.Wrap(err, "grpc.DialContext")
	}

	return conn, nil
}
