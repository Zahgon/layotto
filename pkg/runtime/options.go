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

package runtime

import (
	"google.golang.org/grpc"

	"mosn.io/layotto/components/oss"

	"mosn.io/layotto/components/configstores"
	"mosn.io/layotto/components/custom"
	"mosn.io/layotto/components/file"
	"mosn.io/layotto/components/hello"
	"mosn.io/layotto/components/rpc"
	rgrpc "mosn.io/layotto/pkg/grpc"
	mbindings "mosn.io/layotto/pkg/runtime/bindings"
	runtime_lock "mosn.io/layotto/pkg/runtime/lock"
	"mosn.io/layotto/pkg/runtime/pubsub"
	msecretstores "mosn.io/layotto/pkg/runtime/secretstores"
	runtime_sequencer "mosn.io/layotto/pkg/runtime/sequencer"
	"mosn.io/layotto/pkg/runtime/state"
)

// services encapsulates the service to include in the runtime
type services struct {
	hellos        []*hello.HelloFactory
	configStores  []*configstores.StoreFactory
	rpcs          []*rpc.Factory
	files         []*file.Factory
	oss           []*oss.Factory
	pubSubs       []*pubsub.Factory
	states        []*state.Factory
	locks         []*runtime_lock.Factory
	sequencers    []*runtime_sequencer.Factory
	outputBinding []*mbindings.OutputBindingFactory
	inputBinding  []*mbindings.InputBindingFactory
	secretStores  []*msecretstores.Factory
	// Custom components.
	// The key is component kind
	custom map[string][]*custom.Factory
	extensionComponentFactorys
}

type runtimeOptions struct {
	// services
	services services
	// other config options
	srvMaker rgrpc.NewServer
	errInt   ErrInterceptor
	options  []grpc.ServerOption
	// new grpc api
	apiFactorys []rgrpc.NewGrpcAPI
}

func newRuntimeOptions() *runtimeOptions { _ = "STUB: not implemented"; return nil }

type Option func(o *runtimeOptions)

func WithNewServer(f rgrpc.NewServer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGrpcOptions(options ...grpc.ServerOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithGrpcAPI(apiFuncs ...rgrpc.NewGrpcAPI) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type ErrInterceptor func(err error, format string, args ...interface{})

func WithErrInterceptor(i ErrInterceptor) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCustomComponentFactory(kind string, factorys ...*custom.Factory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithHelloFactory(hellos ...*hello.HelloFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithConfigStoresFactory(configStores ...*configstores.StoreFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithRpcFactory(rpcs ...*rpc.Factory) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithOssFactory(oss ...*oss.Factory) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithFileFactory(files ...*file.Factory) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPubSubFactory(factorys ...*pubsub.Factory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLockFactory(factorys ...*runtime_lock.Factory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithStateFactory(factorys ...*state.Factory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithInputBindings adds input binding components to the runtime.
func WithInputBindings(factorys ...*mbindings.InputBindingFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithOutputBindings adds output binding components to the runtime.
func WithOutputBindings(factorys ...*mbindings.OutputBindingFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSequencerFactory(factorys ...*runtime_sequencer.Factory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSecretStoresFactory(factorys ...*msecretstores.Factory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
