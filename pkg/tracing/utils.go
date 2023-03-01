package tracing

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc/metadata"
)

// StartGrpcServerTracerSpan starts grpc server tracer span
func StartGrpcServerTracerSpan(ctx context.Context, operationName string) (context.Context, opentracing.Span) {
	// extract metadata from context
	metadataMap := make(opentracing.TextMapCarrier)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		for key := range md.Copy() {
			metadataMap.Set(key, md.Get(key)[0])
		}
	}

	// extract span from metadata
	span, err := opentracing.GlobalTracer().Extract(opentracing.TextMap, metadataMap)
	if err != nil {
		// if error occurred, start new span
		sp := opentracing.GlobalTracer().StartSpan(operationName)
		// set span as server span
		ctx = opentracing.ContextWithSpan(ctx, sp)
		return ctx, sp
	}

	// start new span with extracted span
	sp := opentracing.GlobalTracer().StartSpan(operationName, ext.RPCServerOption(span))
	// set span as server span
	ctx = opentracing.ContextWithSpan(ctx, sp)

	return ctx, sp
}
