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

package dapr

import (
	"context"

	"github.com/dapr/components-contrib/bindings"
	"github.com/dapr/components-contrib/pubsub"
	"github.com/dapr/components-contrib/secretstores"
	"github.com/dapr/components-contrib/state"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"

	"mosn.io/layotto/components/configstores"
	"mosn.io/layotto/components/file"
	"mosn.io/layotto/components/hello"
	"mosn.io/layotto/components/lock"
	"mosn.io/layotto/components/rpc"
	"mosn.io/layotto/components/sequencer"
	grpc_api "mosn.io/layotto/pkg/grpc"
	dapr_common_v1pb "mosn.io/layotto/pkg/grpc/dapr/proto/common/v1"
	dapr_v1pb "mosn.io/layotto/pkg/grpc/dapr/proto/runtime/v1"
)

type DaprGrpcAPI interface {
	dapr_v1pb.DaprServer
	grpc_api.GrpcAPI
}

type daprGrpcAPI struct {
	appId                    string
	hellos                   map[string]hello.HelloService
	configStores             map[string]configstores.Store
	rpcs                     map[string]rpc.Invoker
	pubSubs                  map[string]pubsub.PubSub
	stateStores              map[string]state.Store
	transactionalStateStores map[string]state.TransactionalStore
	fileOps                  map[string]file.File
	lockStores               map[string]lock.LockStore
	sequencers               map[string]sequencer.Store
	sendToOutputBindingFn    func(name string, req *bindings.InvokeRequest) (*bindings.InvokeResponse, error)
	secretStores             map[string]secretstores.SecretStore
	// app callback
	AppCallbackConn   *grpc.ClientConn
	topicPerComponent map[string]TopicSubscriptions
	// json
	json jsoniter.API
}

func (d *daprGrpcAPI) Init(conn *grpc.ClientConn) error {
	_ = "STUB: not implemented"
	// 1. set connection
	return nil
}

func (d *daprGrpcAPI) Register(rawGrpcServer *grpc.Server) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *daprGrpcAPI) InvokeService(ctx context.Context, in *dapr_v1pb.InvokeServiceRequest) (*dapr_common_v1pb.InvokeResponse, error) {
	_ = "STUB: not implemented"
	// 1. convert request to RPCRequest,which is the parameter for RPC components
	return nil, nil
}

// 2. route to the specific rpc.Invoker component.
// Only support mosn component now.

// 3. delegate to the rpc.Invoker component

// 4. convert result

// 5. convert result

// fix https://github.com/mosn/layotto/issues/285

func (d *daprGrpcAPI) InvokeBinding(ctx context.Context, in *dapr_v1pb.InvokeBindingRequest) (*dapr_v1pb.InvokeBindingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *daprGrpcAPI) isSecretAllowed(storeName string, key string) bool {
	_ = "STUB: not implemented"
	// TODO: add permission control
	return false
}

// NewDaprAPI_Alpha construct a grpc_api.GrpcAPI which implements DaprServer.
// Currently it only support Dapr's InvokeService and InvokeBinding API.
// Note: this feature is still in Alpha state and we don't recommend that you use it in your production environment.
func NewDaprAPI_Alpha(ac *grpc_api.ApplicationContext) grpc_api.GrpcAPI {
	_ = "STUB: not implemented"
	// filter out transactionalStateStores
	return *new(grpc_api.GrpcAPI)
}

func NewDaprServer(
	appId string,
	hellos map[string]hello.HelloService,
	configStores map[string]configstores.Store,
	rpcs map[string]rpc.Invoker,
	pubSubs map[string]pubsub.PubSub,
	stateStores map[string]state.Store,
	transactionalStateStores map[string]state.TransactionalStore,
	files map[string]file.File,
	lockStores map[string]lock.LockStore,
	sequencers map[string]sequencer.Store,
	sendToOutputBindingFn func(name string, req *bindings.InvokeRequest) (*bindings.InvokeResponse, error),
	secretStores map[string]secretstores.SecretStore,
) DaprGrpcAPI {
	_ = "STUB: not implemented"
	// construct
	return *new(DaprGrpcAPI)
}
