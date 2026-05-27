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

package skywalking

import (
	"context"
	"time"

	"github.com/SkyAPM/go2sky"
	"mosn.io/api"
	"mosn.io/mosn/pkg/trace/skywalking"
	"mosn.io/mosn/pkg/types"

	ltrace "mosn.io/layotto/components/trace"
)

func NewGrpcSkyTracer(_ map[string]interface{}) (api.Tracer, error) {
	_ = "STUB: not implemented"
	return *new(api.Tracer), nil
}

type grpcSkyTracer struct {
	*go2sky.Tracer
}

func (tracer *grpcSkyTracer) SetGO2SkyTracer(t *go2sky.Tracer) { _ = "STUB: not implemented"; return }

func (tracer *grpcSkyTracer) Start(ctx context.Context, request interface{}, _ time.Time) api.Span {
	_ = "STUB: not implemented"
	return *new(api.Span)
}

// create entry span (downstream)

type grpcSkySpan struct {
	*ltrace.Span
	tracer  *grpcSkyTracer
	ctx     context.Context
	carrier *skywalking.SpanCarrier
}

func (h *grpcSkySpan) TraceId() string { _ = "STUB: not implemented"; return "" }

func (h *grpcSkySpan) InjectContext(requestHeaders types.HeaderMap, requestInfo api.RequestInfo) {
	_ = "STUB: not implemented"
	return
}

func (h *grpcSkySpan) SetRequestInfo(requestInfo api.RequestInfo) {
	_ = "STUB: not implemented"
	return
}

func (h *grpcSkySpan) FinishSpan() { _ = "STUB: not implemented"; return }
