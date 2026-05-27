// Copyright 2021 Layotto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package trace

import (
	"time"

	"mosn.io/api"

	"mosn.io/mosn/pkg/trace/sofa"
	"mosn.io/mosn/pkg/types"
)

type Span struct {
	StartTime     time.Time
	EndTime       time.Time
	traceId       string
	spanId        string
	parentSpanId  string
	tags          [sofa.TRACE_END]string
	operationName string
}

func (span *Span) SetTraceId(id string) { _ = "STUB: not implemented"; return }

func (span *Span) TraceId() string { _ = "STUB: not implemented"; return "" }

func (span *Span) SetSpanId(id string) { _ = "STUB: not implemented"; return }

func (span *Span) SpanId() string { _ = "STUB: not implemented"; return "" }

func (span *Span) SetParentSpanId(id string) { _ = "STUB: not implemented"; return }

func (span *Span) ParentSpanId() string { _ = "STUB: not implemented"; return "" }

func (span *Span) SetOperation(operation string) { _ = "STUB: not implemented"; return }

func (span *Span) SetTag(key uint64, value string) { _ = "STUB: not implemented"; return }

func (span *Span) SetRequestInfo(reqInfo types.RequestInfo) { _ = "STUB: not implemented"; return }

func (span *Span) Tag(key uint64) string { _ = "STUB: not implemented"; return "" }

func (span *Span) FinishSpan() { _ = "STUB: not implemented"; return }

func (span *Span) InjectContext(requestHeaders types.HeaderMap, requestInfo types.RequestInfo) {
	_ = "STUB: not implemented"
	return
}

func (span *Span) SpawnChild(operationName string, startTime time.Time) api.Span {
	_ = "STUB: not implemented"
	return *new(api.Span)
}
