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

package callback

import (
	"encoding/json"

	"mosn.io/layotto/components/rpc"
)

// RegisterBeforeInvoke is set BeforeFactory
func RegisterBeforeInvoke(f BeforeFactory) { _ = "STUB: not implemented"; return }

// RegisterAfterInvoke is set AfterFactory
func RegisterAfterInvoke(f AfterFactory) { _ = "STUB: not implemented"; return }

// BeforeFactory is handled RPCRequest
type BeforeFactory interface {
	// Name is create beforeFactory name
	Name() string
	// Init is init RawMessage
	Init(json.RawMessage) error
	// Create is exec specific logic
	Create() func(*rpc.RPCRequest) (*rpc.RPCRequest, error)
}

// AfterFactory is handled RPCResponse
type AfterFactory interface {
	// Name is create afterFactory name
	Name() string
	// Init is init RawMessage
	Init(json.RawMessage) error
	// Create is exec specific logic
	Create() func(*rpc.RPCResponse) (*rpc.RPCResponse, error)
}

var (
	// to storage BeforeFactory
	beforeInvokeRegistry = map[string]BeforeFactory{}
	// to storage AfterFactory
	afterInvokeRegistry = map[string]AfterFactory{}
)

// NewCallback is created Callback
func NewCallback() rpc.Callback { _ = "STUB: not implemented"; return *new(rpc.Callback) }

type callback struct {
	beforeInvoke []func(*rpc.RPCRequest) (*rpc.RPCRequest, error)
	afterInvoke  []func(*rpc.RPCResponse) (*rpc.RPCResponse, error)
}

// AddBeforeInvoke is add beforeInvoke into callback.beforeInvoke
func (c *callback) AddBeforeInvoke(conf rpc.CallbackFunc) { _ = "STUB: not implemented"; return }

// AddAfterInvoke is used to add beforeInvoke into callback.afterInvoke
func (c *callback) AddAfterInvoke(conf rpc.CallbackFunc) { _ = "STUB: not implemented"; return }

// BeforeInvoke is used to invoke beforeInvoke callbacks
func (c *callback) BeforeInvoke(request *rpc.RPCRequest) (*rpc.RPCRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AfterInvoke is used to invoke afterInvoke callbacks
func (c *callback) AfterInvoke(response *rpc.RPCResponse) (*rpc.RPCResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
