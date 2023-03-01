package app

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MetricsServer is a metrics server.
func newMetricsServer(path string) *echo.Echo {
	metricsServer := echo.New()
	metricsServer.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         stackSize,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))

	metricsServer.GET(path, echo.WrapHandler(promhttp.Handler()))

	return metricsServer
}

func (a *App) RunMetricsServer(cancel context.CancelFunc) {
	go func() {
		a.logger.Infof("Metrics server is running on port %v", a.cfg.Probes.PrometheusPort)

		err := a.metricsServer.Start(a.cfg.Probes.PrometheusPort)
		if err != nil {
			a.logger.Errorf("Metrics server failed to start: %v", err)
			cancel()
		}
	}()
}

func (a *App) CloseMetricsServer(ctx context.Context) {
	a.logger.Infof("Closing metrics server...")

	err := a.metricsServer.Shutdown(ctx)
	if err != nil {
		a.logger.Warnf("Error while closing metrics server: %v", err)
	}
}
