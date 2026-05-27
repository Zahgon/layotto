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

package helloworld

import (
	"context"

	rawGRPC "google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"

	"mosn.io/layotto/cmd/layotto_multiple_api/helloworld/component"
	"mosn.io/layotto/components/lock"
	"mosn.io/layotto/pkg/grpc"
	grpc_api "mosn.io/layotto/pkg/grpc"
)

const kind = "helloworld"

// This demo will always use this component name.
const componentName = "demo"

func NewHelloWorldAPI(ac *grpc_api.ApplicationContext) grpc.GrpcAPI {
	_ = "STUB: not implemented"
	// 1. convert custom components
	return *new(grpc.GrpcAPI)
}

// we only care about those components of type "helloworld"

// convert them using type assertion

// 2. construct your API implementation

// Your API plugin can store and use all the components.
// For example,this demo set all the LockStore components here.

// Custom components of type "helloworld"

// server is used to implement helloworld.GreeterServer.
type server struct {
	appId string
	// custom components which implements the `HelloWorld` interface
	name2component map[string]component.HelloWorld
	// LockStore components. They are not used in this demo, we put them here as a demo.
	name2LockStore map[string]lock.LockStore
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer.SayHello
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) Init(conn *rawGRPC.ClientConn) error { _ = "STUB: not implemented"; return nil }

func (s *server) Register(rawGrpcServer *rawGRPC.Server) error {
	_ = "STUB: not implemented"
	return nil
}
