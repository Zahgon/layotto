/*
 * Copyright 2021 Layotto Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package jaeger

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	jaegerc "github.com/uber/jaeger-client-go"
	"mosn.io/api"
	"mosn.io/mosn/pkg/types"

	ltrace "mosn.io/layotto/components/trace"
)

const (
	serviceName              = "service_name"
	strategy                 = "strategy"
	agentHost                = "agent_host"
	collectorEndpoint        = "collector_endpoint"
	defaultServiceName       = "layotto"
	defaultJaegerAgentHost   = "127.0.0.1:6831"
	jaegerAgentHostKey       = "TRACE"
	appIDKey                 = "APP_ID"
	defaultCollectorEndpoint = "http://127.0.0.1:14268/api/traces"
	defaultStrategy          = "collector"
)

type grpcJaegerTracer struct {
	tracer opentracing.Tracer
}

type grpcJaegerSpan struct {
	*ltrace.Span
	ctx        context.Context
	trace      *grpcJaegerTracer
	jaegerSpan opentracing.Span
	spanCtx    jaegerc.SpanContext
}

func NewGrpcJaegerTracer(traceCfg map[string]interface{}) (api.Tracer, error) {
	_ = "STUB: not implemented"
	// 1. construct the ReporterConfig, which is used to communicate with jaeger
	return *new(api.Tracer), nil
}

// Determining whether to start the agent

// 2. construct the Configuration

// 3. use the Configuration to construct a new tracer

// 4. adapt to the `api.Tracer`

func getAgentHost(traceCfg map[string]interface{}) string { _ = "STUB: not implemented"; return "" }

//if TRACE is not set, get it from the env variable

func getStrategy(traceCfg map[string]interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getCollectorEndpoint(traceCfg map[string]interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func getServiceName(traceCfg map[string]interface{}) string { _ = "STUB: not implemented"; return "" }

//if service_name is not set, get it from the env variable

func (t *grpcJaegerTracer) Start(ctx context.Context, request interface{}, startTime time.Time) api.Span {
	_ = "STUB: not implemented"
	return *new(api.Span)
}

//create entry span (downstream)

//renew span context

func (s *grpcJaegerSpan) TraceId() string { _ = "STUB: not implemented"; return "" }

func (s *grpcJaegerSpan) InjectContext(requestHeaders types.HeaderMap, requestInfo api.RequestInfo) {
	_ = "STUB: not implemented"
	return
}

func (s *grpcJaegerSpan) SetRequestInfo(requestInfo api.RequestInfo) {
	_ = "STUB: not implemented"
	return
}

func (s *grpcJaegerSpan) FinishSpan() { _ = "STUB: not implemented"; return }
