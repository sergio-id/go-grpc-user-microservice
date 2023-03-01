package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
	"io"
)

// NewJaegerTracer creates new jaeger tracer
func NewJaegerTracer(jaegerConfig *Config) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: jaegerConfig.ServiceName,

		// Set the sampler to the remote sampler.
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},

		// Set the reporter to the remote reporter.
		Reporter: &config.ReporterConfig{
			LogSpans:           jaegerConfig.LogSpans,
			LocalAgentHostPort: jaegerConfig.HostPort,
		},
	}

	// Set the Zipkin B3 Propagator
	p := zipkin.NewZipkinB3HTTPHeaderPropagator()

	return cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
		config.Injector(opentracing.HTTPHeaders, p),
		config.Injector(opentracing.TextMap, p),
		config.Injector(opentracing.Binary, p),
		config.Extractor(opentracing.HTTPHeaders, p),
		config.Extractor(opentracing.TextMap, p),
		config.Extractor(opentracing.Binary, p),
		config.ZipkinSharedRPCSpan(false),
	)
}
