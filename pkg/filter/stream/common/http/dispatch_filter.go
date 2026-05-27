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

package http

import (
	"context"

	"mosn.io/api"
	"mosn.io/pkg/buffer"
)

type DispatchFilter struct {
	filterType     string
	requestHandler RequestHandler
	handler        api.StreamReceiverFilterHandler
}

func (dis *DispatchFilter) SetReceiveFilterHandler(handler api.StreamReceiverFilterHandler) {
	_ = "STUB: not implemented"
	return
}

func (dis *DispatchFilter) OnDestroy() { _ = "STUB: not implemented"; return }

func (dis *DispatchFilter) OnReceive(ctx context.Context, headers api.HeaderMap, buf buffer.IoBuffer, trailers api.HeaderMap) api.StreamFilterStatus {
	_ = "STUB: not implemented"
	// 1. log
	return *new(api.StreamFilterStatus)
}

// 2. validate path

// http path must be /{dis.filterType}/{endpoint_name}/{params}
// So we can return 404 directly if it does not start with {dis.filterType}

// illegal

// 3. process request

// illegal

func (dis *DispatchFilter) write404() { _ = "STUB: not implemented"; return }

func (dis *DispatchFilter) writeJsonResult(jsonObject map[string]interface{}, code int) {
	_ = "STUB: not implemented"
	return
}

// 0. marshal

// 1. header

// 2. body

// 3. write response
