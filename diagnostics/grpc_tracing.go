package diagnostics

import (
	"context"

	"google.golang.org/grpc"
)

// UnaryInterceptorFilter is an implementation of grpc.UnaryServerInterceptor
func UnaryInterceptorFilter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get tracer

// start a span

// construct a new context which contains the span

// handle request

func StreamInterceptorFilter(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// get tracer

// start a span

// construct a new context which contains the span

// handle request
