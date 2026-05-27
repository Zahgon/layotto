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

package tcpcopy

import (
	"context"
	"errors"

	"mosn.io/api"
	"mosn.io/mosn/pkg/types"
)

func init() {
	api.RegisterNetwork("tcpcopy", CreateTcpcopyFactory)
}

var (
	ErrInvalidConfig = errors.New("invalid config for tcpcopy")
)

type config struct {
	port string
}

type tcpcopyFactory struct {
	cfg *config
}

func CreateTcpcopyFactory(cfg map[string]interface{}) (api.NetworkFilterChainFactory, error) {
	_ = "STUB: not implemented"
	return *

	// Parse static config for dump strategy
	new(api.NetworkFilterChainFactory), nil
}

// TODO extract some other fields

func (f *tcpcopyFactory) Init(param interface{}) error {
	_ = "STUB: not implemented"
	// 1. get listener config
	return nil
}

// 2. parse listener port

// 3. set config

func (f *tcpcopyFactory) CreateFilterChain(context context.Context, callbacks api.NetWorkFilterChainFactoryCallbacks) {
	_ = "STUB: not implemented"
	return
}

func (f *tcpcopyFactory) OnData(data types.IoBuffer) (res api.FilterStatus) {
	_ = "STUB: not implemented"
	// Determine whether to continue sampling
	return *new(api.FilterStatus)
}

// Asynchronous sampling

func (f *tcpcopyFactory) OnNewConnection() api.FilterStatus {
	_ = "STUB: not implemented"
	return *new(api.FilterStatus)
}

func (f *tcpcopyFactory) InitializeReadFilterCallbacks(cb api.ReadFilterCallbacks) {
	_ = "STUB: not implemented"
	return
}
