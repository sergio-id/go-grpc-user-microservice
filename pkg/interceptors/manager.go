package interceptors

import (
	"context"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

// GrpcMetricsCb is a callback function for metrics
type GrpcMetricsCb func(err error)

// InterceptorManager interface
type InterceptorManager interface {
	Logger(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error)
	ClientRequestLoggerInterceptor() func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error
}

// interceptorManager struct
type interceptorManager struct {
	log       logger.Logger
	metricsCb GrpcMetricsCb
}

// NewInterceptorManager constructor
func NewInterceptorManager(logger logger.Logger, metricsCb GrpcMetricsCb) *interceptorManager {
	return &interceptorManager{log: logger, metricsCb: metricsCb}
}

// Logger Interceptor for logging requests and responses to the server side of the gRPC connection.
func (im *interceptorManager) Logger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)
	im.log.GrpcMiddlewareAccessLogger(info.FullMethod, time.Since(start), md, err)
	if im.metricsCb != nil {
		im.metricsCb(err)
	}
	return reply, err
}

// ClientRequestLoggerInterceptor gRPC client interceptor for logging requests and responses to the server side of the gRPC connection.
func (im *interceptorManager) ClientRequestLoggerInterceptor() func(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)

		md, _ := metadata.FromIncomingContext(ctx)

		im.log.GrpcClientInterceptorLogger(method, req, reply, time.Since(start), md, err)

		return err
	}
}
