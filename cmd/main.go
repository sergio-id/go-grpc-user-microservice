package main

import (
	"context"
	"flag"
	"github.com/sergio-id/go-grpc-user-microservice/config"
	"github.com/sergio-id/go-grpc-user-microservice/internal/app"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// main is the entry point for the user microservice.
// It initializes the microservice and starts the gRPC server.
func main() {
	log.Println("🚀Starting user microservice")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	appLog := logger.NewAppLogger(cfg.Logger)
	appLog.InitLogger()
	appLog.Named(cfg.ServiceName)
	appLog.Infof("CFG APP: %#v", cfg)

	a, cleanup, err := app.InitApp(ctx, appLog, *cfg)
	defer cleanup()
	if err != nil {
		appLog.Errorf("failed init app", err)
		cancel()
	}

	if err = a.RunMigrate(); err != nil {
		appLog.Errorf("failed run migrate", err)
		cancel()
	}

	if cfg.Jaeger.Enable {
		closerTrace := a.RunJaegerServer(cancel)
		defer a.CloseJaegerServer(closerTrace)
	}

	a.RunHealthCheckServer(cancel)
	defer a.CloseHealthCheckServer(ctx)

	a.RunMetricsServer(cancel)
	defer a.CloseMetricsServer(ctx)

	a.RunGRPCServer(cancel)
	defer a.StopGRPCServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		appLog.Infof("signal.Notify", v)
	case done := <-ctx.Done():
		appLog.Infof("ctx.Done", done)
	}
}
