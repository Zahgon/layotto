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

package zipkin

import (
	"context"
	"time"

	"github.com/openzipkin/zipkin-go"
	"mosn.io/api"
	"mosn.io/mosn/pkg/types"

	ltrace "mosn.io/layotto/components/trace"
)

const (
	service_name       = "service_name"
	reporter_endpoint  = "reporter_endpoint"
	recorder_host_post = "recorder_host_post"
)

type grpcZipTracer struct {
	*zipkin.Tracer
}

type grpcZipSpan struct {
	*ltrace.Span
	tracer *grpcZipTracer
	ctx    context.Context
	span   zipkin.Span
}

func NewGrpcZipTracer(traceCfg map[string]interface{}) (api.Tracer, error) {
	_ = "STUB: not implemented"
	return *new(api.Tracer), nil
}

func getRecorderHostPort(traceCfg map[string]interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getReporterEndpoint(traceCfg map[string]interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getServerName(traceCfg map[string]interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *grpcZipTracer) Start(ctx context.Context, request interface{}, _ time.Time) api.Span {
	_ = "STUB: not implemented"
	return *new(api.Span)
}

// start span

func (s *grpcZipSpan) TraceId() string { _ = "STUB: not implemented"; return "" }

func (s *grpcZipSpan) InjectContext(requestHeaders types.HeaderMap, requestInfo api.RequestInfo) {
	_ = "STUB: not implemented"
	return
}

func (s *grpcZipSpan) SetRequestInfo(requestInfo api.RequestInfo) {
	_ = "STUB: not implemented"
	return
}

func (s *grpcZipSpan) FinishSpan() { _ = "STUB: not implemented"; return }
