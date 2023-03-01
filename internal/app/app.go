package app

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sergio-id/go-grpc-user-microservice/config"
	"github.com/sergio-id/go-grpc-user-microservice/internal/metrics"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/repository"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/usecase"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/interceptors"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/logger"
	"google.golang.org/grpc"
	"net/http"
)

type App struct {
	cfg               config.Config
	logger            logger.Logger
	db                *sqlx.DB
	redisClient       *redis.Client
	grpcServer        *grpc.Server
	healthCheckServer *http.Server
	metricsServer     *echo.Echo
	im                interceptors.InterceptorManager
	metrics           *metrics.Metrics
	validate          *validator.Validate
}

// NewApp creates a new instance of the App struct with the provided logger and config.
func NewApp(
	logger logger.Logger,
	cfg config.Config,
	db *sqlx.DB,
	redisClient *redis.Client,
	grpcServer *grpc.Server,
	metricsServer *echo.Echo,
	healthCheckServer *http.Server,
	validator *validator.Validate,
) *App {
	return &App{
		logger:            logger,
		cfg:               cfg,
		db:                db,
		redisClient:       redisClient,
		grpcServer:        grpcServer,
		metricsServer:     metricsServer,
		healthCheckServer: healthCheckServer,
		validate:          validator,
	}
}

func InitApp(ctx context.Context, logger logger.Logger, cfg config.Config) (*App, func(), error) {
	// connect postgres
	db, closeDB, err := getDB(cfg.Postgres)
	if err != nil {
		return nil, nil, err
	}

	// connect redis
	redisClient, closeRedis, err := getRedisClient(cfg.Redis)
	if err != nil {
		return nil, nil, err
	}

	// init use case
	userPostgresRepo := repository.NewUserPostgresqlRepository(db)
	userRedisRepo := repository.NewUserRedisRepository(redisClient)
	userUseCase := usecase.NewUserUseCase(logger, userPostgresRepo, userRedisRepo)

	// init metrics
	m := metrics.NewMetrics(cfg)
	im := interceptors.NewInterceptorManager(logger, getGrpcMetricsCb(m))

	// init validator
	validate := newValidate()

	// init grpc server
	grpcServer, err := newUserGrpcServer(logger, cfg, userUseCase, validate, m, im)
	if err != nil {
		return nil, nil, err
	}

	// init metrics server
	metricsServer := newMetricsServer(cfg.Probes.PrometheusPath)

	// init health check server
	healthCheckServer := newHealthCheckServer(ctx, cfg, logger, db, redisClient)

	// init app
	a := NewApp(logger, cfg, db, redisClient, grpcServer, metricsServer, healthCheckServer, validate)

	return a, func() {
		closeRedis()
		closeDB()
	}, nil
}
