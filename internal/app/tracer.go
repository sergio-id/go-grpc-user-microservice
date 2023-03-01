package app

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/tracing"
	"io"
)

func (a *App) RunJaegerServer(cancel context.CancelFunc) io.Closer {
	a.logger.Infof("Jaeger app is running on port: %s", a.cfg.Jaeger.HostPort)

	tracer, closer, err := tracing.NewJaegerTracer(&a.cfg.Jaeger)
	if err != nil {
		a.logger.Errorf("failed to initialize jaeger: %v", err)
		cancel()
	}
	opentracing.SetGlobalTracer(tracer)

	return closer
}

func (a *App) CloseJaegerServer(closerTrace io.Closer) {
	a.logger.Infof("Closing jaeger server...")

	err := closerTrace.Close()
	if err != nil {
		a.logger.Warnf("failed to close tracer: %v", err)
	}
}
