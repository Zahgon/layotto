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

package mosn

import (
	"context"

	// bridge to mosn
	_ "mosn.io/mosn/pkg/filter/network/proxy"

	"mosn.io/layotto/components/rpc"
	"mosn.io/layotto/components/rpc/invoker/mosn/channel"
)

const (
	Name = "mosn"
)

// mosnInvoker is Invoker implement
type mosnInvoker struct {
	channel rpc.Channel
	cb      rpc.Callback
}

// mosnConfig is mosn config
type mosnConfig struct {
	Before  []rpc.CallbackFunc      `json:"before_invoke"`
	After   []rpc.CallbackFunc      `json:"after_invoke"`
	Channel []channel.ChannelConfig `json:"channel"`
}

// NewMosnInvoker is init mosnInvoker
func NewMosnInvoker() rpc.Invoker { _ = "STUB: not implemented"; return *new(rpc.Invoker) }

// Init is init mosn RpcConfig
func (m *mosnInvoker) Init(conf rpc.RpcConfig) error { _ = "STUB: not implemented"; return nil }

// todo support multiple channel

// Invoke is invoke mosn RPCRequest and Context to RPCResponse
func (m *mosnInvoker) Invoke(ctx context.Context, req *rpc.RPCRequest) (resp *rpc.RPCResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. validate request

// 2. beforeInvoke callback

// 3. do invocation

// 4. afterInvoke callback
