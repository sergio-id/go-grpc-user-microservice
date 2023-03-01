package app

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/heptiolabs/healthcheck"
	"github.com/jmoiron/sqlx"
	"github.com/sergio-id/go-grpc-user-microservice/config"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/constants"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/logger"
	"net/http"
	"time"
)

const (
	stackSize    = 1 << 10 // 1 KB
	readTimeout  = 15 * time.Second
	writeTimeout = 15 * time.Second
)

// newHealthCheckServer creates a new health check server.
func newHealthCheckServer(
	ctx context.Context,
	cfg config.Config,
	logger logger.Logger,
	db *sqlx.DB,
	redisClient *redis.Client,
) *http.Server {
	health := healthcheck.NewHandler()

	mux := http.NewServeMux()

	healthCheckServer := &http.Server{
		Handler:      mux,
		Addr:         cfg.Probes.Port,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}

	// Maps the liveness and readiness endpoints to the corresponding health check functions.
	mux.HandleFunc(cfg.Probes.LivenessPath, health.LiveEndpoint)
	mux.HandleFunc(cfg.Probes.ReadinessPath, health.ReadyEndpoint)

	configureHealthCheckEndpoints(ctx, logger, health, cfg, db, redisClient)

	return healthCheckServer
}

// configureHealthCheckEndpoints configures the custom health check endpoints.
func configureHealthCheckEndpoints(
	ctx context.Context,
	logger logger.Logger,
	health healthcheck.Handler,
	cfg config.Config,
	db *sqlx.DB,
	redisClient *redis.Client,
) {
	// Add a readiness and liveness check for Postgres by pinging the database asynchronously with context.
	// If the ping returns an error, it is logged with a warning message.
	health.AddReadinessCheck(constants.Postgres, healthcheck.AsyncWithContext(ctx, func() error {
		if err := db.Ping(); err != nil {
			logger.Warnf("(Postgres Readiness Check) err: %v", err)
			return err
		}
		return nil
	}, time.Duration(cfg.Probes.CheckIntervalSeconds)*time.Second))

	health.AddLivenessCheck(constants.Postgres, healthcheck.AsyncWithContext(ctx, func() error {
		if err := db.Ping(); err != nil {
			logger.Warnf("(Postgres Liveness Check) err: %v", err)
			return err
		}
		return nil
	}, time.Duration(cfg.Probes.CheckIntervalSeconds)*time.Second))

	health.AddReadinessCheck(constants.Redis, healthcheck.AsyncWithContext(ctx, func() error {
		if err := redisClient.Ping(ctx).Err(); err != nil {
			logger.Warnf("(Redis Readiness Check) err: %v", err)
			return err
		}
		return nil
	}, time.Duration(cfg.Probes.CheckIntervalSeconds)*time.Second))

	health.AddLivenessCheck(constants.Redis, healthcheck.AsyncWithContext(ctx, func() error {
		if err := redisClient.Ping(ctx).Err(); err != nil {
			logger.Warnf("(Redis Liveness Check) err: %v", err)
			return err
		}
		return nil
	}, time.Duration(cfg.Probes.CheckIntervalSeconds)*time.Second))
}

func (a *App) RunHealthCheckServer(cancel context.CancelFunc) {
	go func() {
		a.logger.Infof("Health check server is listening on %s", a.healthCheckServer.Addr)

		err := a.healthCheckServer.ListenAndServe()
		if err != nil {
			a.logger.Errorf("Health check server failed to start: %v", err)
			cancel()
		}
	}()
}

func (a *App) CloseHealthCheckServer(ctx context.Context) {
	a.logger.Infof("Shutting down health check server...")

	err := a.healthCheckServer.Shutdown(ctx)
	if err != nil {
		a.logger.Warnf("Health check server failed to shutdown: %v", err)
	}
}
