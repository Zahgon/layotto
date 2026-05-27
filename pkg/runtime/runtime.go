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
	"mosn.io/layotto/components/pkg/common"
	"mosn.io/layotto/pkg/runtime/lifecycle"

	"mosn.io/layotto/components/oss"

	"mosn.io/layotto/pkg/runtime/ref"

	refconfig "mosn.io/layotto/components/ref"

	"github.com/dapr/components-contrib/secretstores"

	"mosn.io/layotto/components/custom"
	msecretstores "mosn.io/layotto/pkg/runtime/secretstores"

	"github.com/dapr/components-contrib/bindings"

	mbindings "mosn.io/layotto/pkg/runtime/bindings"

	"mosn.io/layotto/components/file"

	"github.com/dapr/components-contrib/pubsub"
	"github.com/dapr/components-contrib/state"
	rawGRPC "google.golang.org/grpc"
	mgrpc "mosn.io/mosn/pkg/filter/network/grpc"

	"mosn.io/layotto/components/configstores"
	"mosn.io/layotto/components/hello"
	"mosn.io/layotto/components/lock"
	"mosn.io/layotto/components/pkg/info"
	"mosn.io/layotto/components/rpc"
	"mosn.io/layotto/components/sequencer"
	runtime_lock "mosn.io/layotto/pkg/runtime/lock"
	runtime_pubsub "mosn.io/layotto/pkg/runtime/pubsub"
	runtime_sequencer "mosn.io/layotto/pkg/runtime/sequencer"
	runtime_state "mosn.io/layotto/pkg/runtime/state"
)

type MosnRuntime struct {
	// configs
	runtimeConfig *MosnRuntimeConfig
	info          *info.RuntimeInfo
	srv           mgrpc.RegisteredServer
	// component registry
	helloRegistry           hello.Registry
	configStoreRegistry     configstores.Registry
	rpcRegistry             rpc.Registry
	pubSubRegistry          runtime_pubsub.Registry
	stateRegistry           runtime_state.Registry
	lockRegistry            runtime_lock.Registry
	sequencerRegistry       runtime_sequencer.Registry
	fileRegistry            file.Registry
	ossRegistry             oss.Registry
	bindingsRegistry        mbindings.Registry
	secretStoresRegistry    msecretstores.Registry
	customComponentRegistry custom.Registry
	Injector                *ref.DefaultInjector
	// component pool
	hellos map[string]hello.HelloService
	// config management system component
	configStores map[string]configstores.Store
	rpcs         map[string]rpc.Invoker
	pubSubs      map[string]pubsub.PubSub
	// state implementations store here are already initialized
	states            map[string]state.Store
	files             map[string]file.File
	oss               map[string]oss.Oss
	locks             map[string]lock.LockStore
	sequencers        map[string]sequencer.Store
	outputBindings    map[string]bindings.OutputBinding
	secretStores      map[string]secretstores.SecretStore
	customComponent   map[string]map[string]custom.Component
	dynamicComponents map[lifecycle.ComponentKey]common.DynamicComponent
	extensionComponents
	// app callback
	AppCallbackConn *rawGRPC.ClientConn
	// extend
	errInt            ErrInterceptor
	started           bool
	initRuntimeStages []initRuntimeStage
}

func (m *MosnRuntime) RuntimeConfig() *MosnRuntimeConfig { _ = "STUB: not implemented"; return nil }

type initRuntimeStage func(o *runtimeOptions, m *MosnRuntime) error

func NewMosnRuntime(runtimeConfig *MosnRuntimeConfig) *MosnRuntime {
	_ = "STUB: not implemented"
	return nil
}

func (m *MosnRuntime) GetInfo() *info.RuntimeInfo { _ = "STUB: not implemented"; return nil }

func (m *MosnRuntime) sendToOutputBinding(name string, req *bindings.InvokeRequest) (*bindings.InvokeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MosnRuntime) Run(opts ...Option) (mgrpc.RegisteredServer, error) {
	_ = "STUB: not implemented"
	// 0. mark already started
	return *new(mgrpc.RegisteredServer), nil
}

// 1. init runtime stage
// prepare runtimeOptions

// set ErrInterceptor

// init runtime with runtimeOptions

// prepare grpcOpts

// 2. init GrpcAPI stage

// init the GrpcAPI

// put them into grpc options

// 3. create grpc server

func (m *MosnRuntime) Stop() { _ = "STUB: not implemented"; return }

func (m *MosnRuntime) storeDynamicComponent(kind string, name string, store interface{}) {
	_ = "STUB: not implemented"
	return
}

// put it in the components map

func DefaultInitRuntimeStage(o *runtimeOptions, m *MosnRuntime) error {
	_ = "STUB: not implemented"
	return nil
}

// init callback connection

// init all kinds of components with config
//init secret & config first

func (m *MosnRuntime) initHellos(hellos ...*hello.HelloFactory) error {
	_ = "STUB: not implemented"
	return nil
}

// register all hello services implementation

//inject component

// register this component

func (m *MosnRuntime) initConfigStores(configStores ...*configstores.StoreFactory) error {
	_ = "STUB: not implemented"
	return nil
}

// register all config store services implementation

// register this component

func (m *MosnRuntime) initRpcs(rpcs ...*rpc.Factory) error { _ = "STUB: not implemented"; return nil }

// register all rpc components

// register this component

func (m *MosnRuntime) initPubSubs(factorys ...*runtime_pubsub.Factory) error {
	_ = "STUB: not implemented"
	// 1. init components
	return nil
}

// register all components implementation

// create component

// check consumerID

//inject secret to component

//inject component

// init this component with the config

// register this component

func (m *MosnRuntime) initStates(factorys ...*runtime_state.Factory) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. register all the implementation

// 2. loop initializing

// 2.1. create and store the component

//inject secret to component

//inject component

// 2.2. save prefix strategy

func (m *MosnRuntime) initOss(factorys ...*oss.Factory) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. register all oss store services implementation

// 2. loop initializing

// 2.1. create the component

//inject component

// 2.2. init

// register this component

func (m *MosnRuntime) initFiles(files ...*file.Factory) error {
	_ = "STUB: not implemented"
	return nil
}

// register all files store services implementation

//inject component

func (m *MosnRuntime) initLocks(factorys ...*runtime_lock.Factory) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. register all the implementation

// 2. loop initializing

// 2.1. create the component

//inject secret to component

//inject component

// 2.2. init

// 2.3. save runtime related configs

func (m *MosnRuntime) initSequencers(factorys ...*runtime_sequencer.Factory) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. register all the implementation

// 2. loop initializing

// 2.1. create the component

//inject secret to component

//inject component

// 2.2. init

// 2.3. save runtime related configs

// register this component

func (m *MosnRuntime) initAppCallbackConnection() error {
	_ = "STUB: not implemented"
	// init the client connection for calling app
	return nil
}

// get callback address

// dial

func (m *MosnRuntime) initOutputBinding(factorys ...*mbindings.OutputBindingFactory) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. register all factory methods.

// 2. loop initializing

// 2.1. create the component

//inject secret to component

//inject component

// 2.2. init

// 2.3. put it into the runtime component pool

// TODO: implement initInputBinding
func (m *MosnRuntime) initInputBinding(factorys ...*mbindings.InputBindingFactory) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MosnRuntime) initSecretStores(factorys ...*msecretstores.Factory) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. register all factory methods.

// 2. loop initializing

// 2.1. create the component

//inject component

// 2.2. init

// 2.3. save runtime related configs

func (m *MosnRuntime) AppendInitRuntimeStage(f initRuntimeStage) { _ = "STUB: not implemented"; return }

func (m *MosnRuntime) initRuntime(r *runtimeOptions) error {
	_ = "STUB: not implemented"

	// register pluggable component
	return nil
}

// check default handler

// do initialization

func (m *MosnRuntime) registerPluggableComponent() { _ = "STUB: not implemented"; return }

// todo custom

func (m *MosnRuntime) SetCustomComponent(kind string, name string, component custom.Component) {
	_ = "STUB: not implemented"
	return
}

func (m *MosnRuntime) initCustomComponents(kind2factorys map[string][]*custom.Factory) error {
	_ = "STUB: not implemented"
	return nil
}

// loop all configured custom components.

// register all the factorys

// loop initializing component instances

// create the component

//inject secret to component

//inject component

// init

// initialization finish

func (m *MosnRuntime) initComponentInject(comp interface{}, config *refconfig.ComponentRefConfig) error {
	_ = "STUB: not implemented"
	return nil
}
