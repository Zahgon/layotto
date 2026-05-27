// Copyright 2021 Layotto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hello

import (
	"context"

	"mosn.io/layotto/components/pluggable"
	helloproto "mosn.io/layotto/spec/proto/pluggable/v1/hello"
)

func init() {
	// spec.proto.pluggable.v1.Hello
	pluggable.AddServiceDiscoveryCallback(helloproto.Hello_ServiceDesc.ServiceName, func(compType string, dialer pluggable.GRPCConnectionDialer) pluggable.Component {
		return NewHelloFactory(compType, func() HelloService {
			return NewGRPCHello(dialer)
		})
	})
}

type grpcHello struct {
	dialer pluggable.GRPCConnectionDialer
	client helloproto.HelloClient
}

func NewGRPCHello(dialer pluggable.GRPCConnectionDialer) HelloService {
	_ = "STUB: not implemented"
	return *new(HelloService)
}

// todo 优雅关闭时关闭 conn

func (g *grpcHello) Init(config *HelloConfig) error {
	_ = "STUB: not implemented"
	// 1.dial grpc server
	return nil
}

// 2.init pluggable component

func (g *grpcHello) Hello(ctx context.Context, request *HelloRequest) (*HelloResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
