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

package pluggable

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

// TestServerFor returns a grpcServer factory that bootstraps a grpcserver backed by a buf connection (in memory), and returns the given clientFactory instance to communicate with it.
// it also provides cleanup function for close the grpcserver and client connection.
//
//	usage,
//		serverFactory := TestServerFor(proto.RegisterMyService, proto.NewMyServiceClient)
//	 	client, cleanup, err := serverFactory(&your_service{})
//		require.NoError(t, err)
//		defer cleanup()
func TestServerFor[TServer any, TClient any](registersvc func(grpc.ServiceRegistrar, TServer), clientFactory func(grpc.ClientConnInterface) TClient) func(svc TServer) (client TClient, cleanup func(), err error) {
	_ = "STUB: not implemented"
	return nil
}

// TestSocketServerFor returns a grpcServer factory that bootstraps a grpcserver backed by grpc socket, and returns the given clientFactory instance to communicate with it.
// it also provides cleanup function for close the grpcserver and client connection. The closure of dial needs to be handled by the user themselves
//
//	usage,
//		serverFactory := TestSocketServerFor(proto.RegisterMyService, NewGRPCHello) // (func NewGRPCHello(dialer pluggable.GRPCConnectionDialer) *grpcHello)
//	 	client, cleanup, err := serverFactory(&your_service{})
//		require.NoError(t, err)
//		defer cleanup()
func TestSocketServerFor[TServer any, TClient any](registersvc func(grpc.ServiceRegistrar, TServer), clientFactory func(GRPCConnectionDialer) TClient) func(svc TServer) (client TClient, cleanup func(), err error) {
	_ = "STUB: not implemented"
	return nil
}

func dial(lis *bufconn.Listener) func(ctx context.Context, s string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}
