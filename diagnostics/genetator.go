package diagnostics

import (
	"context"

	"mosn.io/api"

	"mosn.io/layotto/components/trace"
)

func init() {
	trace.RegisterGenerator("mosntracing", &OpenGenerator{})
}

// OpenGenerator is the default implementation of Generator
type OpenGenerator struct {
}

func (o *OpenGenerator) GetTraceId(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

func (o *OpenGenerator) GetSpanId(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// GetParentSpanId returns the same id as GetSpanId.
// It's because currently Layotto don't know the parent id.
func (o *OpenGenerator) GetParentSpanId(ctx context.Context) string {
	_ = "STUB: not implemented"
	// TODO: need some design to get the parent id
	return ""
}

func (o *OpenGenerator) GenerateNewContext(ctx context.Context, span api.Span) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
