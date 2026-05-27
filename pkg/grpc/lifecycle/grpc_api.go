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

package lifecycle

import (
	"context"

	rawGRPC "google.golang.org/grpc"

	"mosn.io/layotto/components/pkg/common"
	"mosn.io/layotto/pkg/runtime/lifecycle"

	"mosn.io/layotto/pkg/grpc"
	grpc_api "mosn.io/layotto/pkg/grpc"
	runtimev1pb "mosn.io/layotto/spec/proto/runtime/v1"
)

func NewLifecycleAPI(ac *grpc_api.ApplicationContext) grpc.GrpcAPI {
	_ = "STUB: not implemented"
	return *new(grpc.GrpcAPI)
}

// server implements runtimev1pb.LifecycleServer
type server struct {
	components map[lifecycle.ComponentKey]common.DynamicComponent
}

func (s *server) ApplyConfiguration(ctx context.Context, in *runtimev1pb.DynamicConfiguration) (*runtimev1pb.ApplyConfigurationResponse, error) {
	_ = "STUB: not implemented"
	// 1. validate parameters
	return nil, nil
}

// 2. find the component

// 3. delegate to the components

func invalidArgumentError(format string, a ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) Init(conn *rawGRPC.ClientConn) error { _ = "STUB: not implemented"; return nil }

func (s *server) Register(rawGrpcServer *rawGRPC.Server) error {
	_ = "STUB: not implemented"
	return nil
}
