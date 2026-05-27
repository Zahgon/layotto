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

package wasm

import (
	"context"
	"sync"

	"mosn.io/api"
	"mosn.io/mosn/pkg/types"
	"mosn.io/pkg/buffer"
	"mosn.io/proxy-wasm-go-host/proxywasm/common"
)

type Filter struct {
	LayottoHandler

	ctx     context.Context
	factory *FilterConfigFactory

	router  *Router
	plugins map[string]*WasmPlugin

	contextID  int32
	pluginUsed *WasmPlugin
	instance   types.WasmInstance
	abi        types.ABI
	exports    Exports

	receiverFilterHandler api.StreamReceiverFilterHandler
	senderFilterHandler   api.StreamSenderFilterHandler

	destroyOnce sync.Once

	requestBuffer  api.IoBuffer
	responseBuffer api.IoBuffer
}

type WasmPlugin struct {
	pluginName string
	plugin     types.WasmPlugin

	// useless for now
	rootContextID     int32
	config            *filterConfigItem
	vmConfigBytes     buffer.IoBuffer
	pluginConfigBytes buffer.IoBuffer
}

// GetVmConfig Get the VmConfig of WasmPlugin
func (p *WasmPlugin) GetVmConfig() common.IoBuffer {
	_ = "STUB: not implemented"
	return *new(common.IoBuffer)
}

// GetPluginConfig Get the plugin config of WasmPlugin
func (p *WasmPlugin) GetPluginConfig() common.IoBuffer {
	_ = "STUB: not implemented"
	return *new(common.IoBuffer)
}

var contextIDGenerator int32

// new context's ID
func newContextID(rootContextID int32) int32 { _ = "STUB: not implemented"; return 0 }

// NewFilter create the filter for a request
func NewFilter(ctx context.Context, factory *FilterConfigFactory) *Filter {
	_ = "STUB: not implemented"
	return nil
}

func (f *Filter) releaseUsedInstance() error { _ = "STUB: not implemented"; return nil }

// OnDestroy Destruction of filters
func (f *Filter) OnDestroy() { _ = "STUB: not implemented"; return }

// SetReceiveFilterHandler Set ReceiveFilterHandler of filter
func (f *Filter) SetReceiveFilterHandler(handler api.StreamReceiverFilterHandler) {
	_ = "STUB: not implemented"
	return
}

// SetSenderFilterHandler Set SenderFilterHandler of filter
func (f *Filter) SetSenderFilterHandler(handler api.StreamSenderFilterHandler) {
	_ = "STUB: not implemented"
	return
}

// Calculate the size of headerMap
func headerMapSize(headers api.HeaderMap) int { _ = "STUB: not implemented"; return 0 }

// OnReceive Reset the filter when receiving then return StreamFilter status
func (f *Filter) OnReceive(ctx context.Context, headers api.HeaderMap, buf buffer.IoBuffer, trailers api.HeaderMap) api.StreamFilterStatus {
	_ = "STUB: not implemented"
	return *new(api.StreamFilterStatus)
}

// Append ResponseData of filter
func (f *Filter) Append(ctx context.Context, headers api.HeaderMap, buf buffer.IoBuffer, trailers api.HeaderMap) api.StreamFilterStatus {
	_ = "STUB: not implemented"
	return *new(api.StreamFilterStatus)
}

// GetRootContextID Get RootContext ID of filter's FilterConfigFactory
func (f *Filter) GetRootContextID() int32 { _ = "STUB: not implemented"; return 0 }

// GetVmConfig Get the used WasmPlugin VmConfig of filter
func (f *Filter) GetVmConfig() common.IoBuffer {
	_ = "STUB: not implemented"
	return *new(common.IoBuffer)
}

// GetPluginConfig Get the used WasmPlugin config of filter
func (f *Filter) GetPluginConfig() common.IoBuffer {
	_ = "STUB: not implemented"
	return *new(common.IoBuffer)
}

// GetHttpRequestHeader Get the HttpRequest header of proxy-wasm
func (f *Filter) GetHttpRequestHeader() common.HeaderMap {
	_ = "STUB: not implemented"
	return *new(common.HeaderMap)
}

// GetHttpRequestBody Get the HttpRequest body of proxy-wasm
func (f *Filter) GetHttpRequestBody() common.IoBuffer {
	_ = "STUB: not implemented"
	return *new(common.IoBuffer)
}

// GetHttpRequestTrailer Get the HttpRequest trailer of proxy-wasm
func (f *Filter) GetHttpRequestTrailer() common.HeaderMap {
	_ = "STUB: not implemented"
	return *new(common.HeaderMap)
}

// GetHttpResponseHeader Get the HttpResponse header of proxy-wasm
func (f *Filter) GetHttpResponseHeader() common.HeaderMap {
	_ = "STUB: not implemented"
	return *new(common.HeaderMap)
}

// GetHttpResponseBody Get the HttpResponse body of proxy-wasm
func (f *Filter) GetHttpResponseBody() common.IoBuffer {
	_ = "STUB: not implemented"
	return *new(common.IoBuffer)
}

// GetHttpResponseTrailer Get the HttpResponse trailer of proxy-wasm
func (f *Filter) GetHttpResponseTrailer() common.HeaderMap {
	_ = "STUB: not implemented"
	return *new(common.HeaderMap)
}
