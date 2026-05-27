package diagnostics

import (
	"context"
	"time"

	"mosn.io/api"
)

const (
	generatorConfigKey = "generator"
	exporterConfigKey  = "exporter"
	defaultGenerator   = "mosntracing"
)

// grpcTracer  is used to start a new Span
type grpcTracer struct {
	config map[string]interface{}
}

func NewTracer(config map[string]interface{}) (api.Tracer, error) {
	_ = "STUB: not implemented"
	return *new(api.Tracer), nil
}

func getActiveExportersFromConfig(config map[string]interface{}) []string {
	_ = "STUB: not implemented"
	return nil
}

func (tracer *grpcTracer) Start(ctx context.Context, request interface{}, startTime time.Time) api.Span {
	_ = "STUB: not implemented"
	return *new(api.Span)
}

// NewSpan constructs a span and tag it with span/trace/parentSpan IDs.
// These IDs are generated using the Generator
func NewSpan(ctx context.Context, startTime time.Time, config map[string]interface{}) api.Span {
	_ = "STUB: not implemented"
	// construct span
	return *new(api.Span)
}

// get generator according to configuration

// use generator to extract the span/trace/parentSpan IDs

// tagging generator type

func GetNewContext(ctx context.Context, span api.Span) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

//if no implement generator, return old ctx
