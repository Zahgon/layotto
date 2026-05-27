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

package default_api

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"

	dapr_common_v1pb "mosn.io/layotto/pkg/grpc/dapr/proto/common/v1"
	dapr_v1pb "mosn.io/layotto/pkg/grpc/dapr/proto/runtime/v1"
	runtimev1pb "mosn.io/layotto/spec/proto/runtime/v1"
)

// GetState obtains the state for a specific key.
func (a *api) GetState(ctx context.Context, in *runtimev1pb.GetStateRequest) (*runtimev1pb.GetStateResponse, error) {
	_ = "STUB: not implemented"
	// Check if the StateRequest is exists
	return nil, nil
}

// convert request

// Generate response by request

func (a *api) SaveState(ctx context.Context, in *runtimev1pb.SaveStateRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	// Check if the request is nil
	return nil, nil
}

// convert request

// delegate to dapr api implementation

// GetBulkState gets a batch of state data
func (a *api) GetBulkState(ctx context.Context, in *runtimev1pb.GetBulkStateRequest) (*runtimev1pb.GetBulkStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert request

// Generate response by request

func (a *api) DeleteState(ctx context.Context, in *runtimev1pb.DeleteStateRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert request

func (a *api) DeleteBulkState(ctx context.Context, in *runtimev1pb.DeleteBulkStateRequest) (*empty.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert request

func (a *api) ExecuteStateTransaction(ctx context.Context, in *runtimev1pb.ExecuteStateTransactionRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert request

// some code for converting from runtimev1pb to dapr_common_v1pb

func convertEtagToDaprPB(etag *runtimev1pb.Etag) *dapr_common_v1pb.Etag {
	_ = "STUB: not implemented"
	return nil
}

func convertOptionsToDaprPB(op *runtimev1pb.StateOptions) *dapr_common_v1pb.StateOptions {
	_ = "STUB: not implemented"
	return nil
}

func convertStatesToDaprPB(states []*runtimev1pb.StateItem) []*dapr_common_v1pb.StateItem {
	_ = "STUB: not implemented"
	return nil
}

func convertTransactionalStateOperationToDaprPB(ops []*runtimev1pb.TransactionalStateOperation) []*dapr_v1pb.TransactionalStateOperation {
	_ = "STUB: not implemented"
	return nil
}
