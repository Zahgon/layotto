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
	"mosn.io/mosn/pkg/wasm/abi/proxywasm010"
	"mosn.io/proxy-wasm-go-host/proxywasm/common"
	proxywasm "mosn.io/proxy-wasm-go-host/proxywasm/v1"
)

// LayottoHandler implement proxywasm.ImportsHandler
type LayottoHandler struct {
	proxywasm010.DefaultImportsHandler

	IoBuffer common.IoBuffer
}

var _ proxywasm.ImportsHandler = &LayottoHandler{}

// GetState Obtains the state for a specific key
func (d *LayottoHandler) GetState(storeName string, key string) (string, proxywasm.WasmResult) {
	_ = "STUB: not implemented"
	return "", *new(proxywasm.WasmResult)
}

// InvokeService Do rpc calls
func (d *LayottoHandler) InvokeService(id string, method string, param string) (string, proxywasm.WasmResult) {
	_ = "STUB: not implemented"
	return "", *new(proxywasm.WasmResult)
}

// GetFuncCallData Get the IoBuffer of LayottoHandler
func (d *LayottoHandler) GetFuncCallData() common.IoBuffer {
	_ = "STUB: not implemented"
	return *new(common.IoBuffer)
}
