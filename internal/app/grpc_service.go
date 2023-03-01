package app

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/sergio-id/go-grpc-user-microservice/config"
	"github.com/sergio-id/go-grpc-user-microservice/internal/metrics"
	grpc2 "github.com/sergio-id/go-grpc-user-microservice/internal/user/delivery/grpc"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/usecase"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/constants"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/interceptors"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/logger"
	"github.com/sergio-id/go-grpc-user-microservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

// newUserGrpcServer creates and starts a new gRPC server for the UserService.
// It returns a function to close the server, the server instance and an error if there is any.
func newUserGrpcServer(
	logger logger.Logger,
	cfg config.Config,
	userUseCase *usecase.UserUseCase,
	validate *validator.Validate,
	metrics *metrics.Metrics,
	im interceptors.InterceptorManager,
) (
	*grpc.Server,
	error,
) {
	// Create a new gRPC server with specified keepalive parameters and unary interceptors
	grpcServer := grpc.NewServer(
		// Keepalive parameters for the gRPC server to detect and mitigate deadlocks in the connection and to close
		// idle connections after a specified time period
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),
		// Unary interceptors for the gRPC server to log requests and responses to the server side of the gRPC
		// connection and to recover from panics in the gRPC server
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),  // Set gRPC context tags
			grpc_prometheus.UnaryServerInterceptor, // Prometheus monitoring for Unary RPCs
			grpc_recovery.UnaryServerInterceptor(), // Recover from panics in the gRPC server
			im.Logger,                              // Log requests and responses to the server side of the gRPC connection
		),
		),
	)

	// Create a new UserService gRPC server implementation
	userGrpcService := grpc2.NewUserServerGRPC(logger, cfg, userUseCase, validate, metrics)

	// Register the UserService gRPC server implementation with the gRPC server
	proto.RegisterUserServiceServer(grpcServer, userGrpcService)
	grpc_prometheus.Register(grpcServer)

	if cfg.GRPC.Development {
		// This allows us to use grpcurl to call the gRPC server
		reflection.Register(grpcServer)
	}

	return grpcServer, nil
}

func (a *App) RunGRPCServer(cancel context.CancelFunc) {
	l, err := net.Listen(constants.Tcp, a.cfg.GRPC.Port)
	if err != nil {
		a.logger.Errorf("net.Listen", err)
		cancel()
	}

	go func() {
		a.logger.Infof("gRPC server is running on port %s", a.cfg.GRPC.Port)

		err = a.grpcServer.Serve(l)
		if err != nil {
			a.logger.Errorf("gRPC server failed to start", err)
			cancel()
		}
	}()
}

func (a *App) StopGRPCServer() {
	a.logger.Infof("gRPC server is stopping...")
	a.grpcServer.GracefulStop()
}
