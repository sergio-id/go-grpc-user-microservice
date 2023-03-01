package app

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sergio-id/go-grpc-user-microservice/internal/metrics"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/migrations"
)

// GetGrpcMetricsCb returns a callback function that increments the metric counters for gRPC requests.
// The returned function takes an error argument and increments the error or success metric accordingly.
func getGrpcMetricsCb(m *metrics.Metrics) func(err error) {
	return func(err error) {
		if err != nil {
			m.ErrorGrpcRequests.Inc()
		} else {
			m.SuccessGrpcRequests.Inc()
		}
	}
}

// RunMigrate runs migrations.
func (a *App) RunMigrate() error {
	a.logger.Infof("Run migrations with config: %+v", a.cfg.Migrations)

	version, dirty, err := migrations.RunMigrations(a.cfg.Migrations)
	if err != nil {
		a.logger.Errorf("RunMigrations err: %v", err)
		return err
	}

	a.logger.Infof("Migrations successfully created: version: %d, dirty: %v", version, dirty)
	return nil
}
