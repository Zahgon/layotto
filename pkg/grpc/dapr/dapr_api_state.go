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

	"github.com/dapr/components-contrib/state"
	"google.golang.org/protobuf/types/known/emptypb"

	dapr_common_v1pb "mosn.io/layotto/pkg/grpc/dapr/proto/common/v1"
	dapr_v1pb "mosn.io/layotto/pkg/grpc/dapr/proto/runtime/v1"
)

func (d *daprGrpcAPI) SaveState(ctx context.Context, in *dapr_v1pb.SaveStateRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	// 1. get store
	return nil, nil
}

// 2. convert requests

// 3. query

// 4. check result

// GetState obtains the state for a specific key.
func (d *daprGrpcAPI) GetState(ctx context.Context, request *dapr_v1pb.GetStateRequest) (*dapr_v1pb.GetStateResponse, error) {
	_ = "STUB: not implemented"
	// 1. get store
	return nil, nil
}

// 2. generate the actual key

// 3. query

// 4. check result

func (d *daprGrpcAPI) GetBulkState(ctx context.Context, request *dapr_v1pb.GetBulkStateRequest) (*dapr_v1pb.GetBulkStateResponse, error) {
	_ = "STUB: not implemented"
	// 1. get store
	return nil, nil
}

// 2. store.BulkGet
// 2.1. convert reqs

// 2.2. query

// 2.3. parse and return result if store supports this method

// 3. Simulate the method if the store doesn't support it

func (d *daprGrpcAPI) QueryStateAlpha1(ctx context.Context, request *dapr_v1pb.QueryStateRequest) (*dapr_v1pb.QueryStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. get state store component

// 2. check if this store has the query feature

// 3. Unmarshal query dsl

// 4. delegate to the store

// 5. convert response

func (d *daprGrpcAPI) DeleteState(ctx context.Context, request *dapr_v1pb.DeleteStateRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	// 1. get store
	return nil, nil
}

// 2. generate the actual key

// 3. convert and send request

// 4. check result

func (d *daprGrpcAPI) DeleteBulkState(ctx context.Context, request *dapr_v1pb.DeleteBulkStateRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	// 1. get store
	return nil, nil
}

// 2. convert request

// 3. send request

// 4. check result

func (d *daprGrpcAPI) ExecuteStateTransaction(ctx context.Context, request *dapr_v1pb.ExecuteStateTransactionRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	// 1. check params
	return nil, nil
}

// 2. find store

// 3. convert request

// 3.1. extract and validate fields

// tolerant npe

// 3.2. prepare TransactionalStateOperation struct according to the operation type

// 4. submit transactional request

// 5. check result

func (d *daprGrpcAPI) getStateStore(name string) (state.Store, error) {
	_ = "STUB: not implemented"
	// check if the stateStores exists
	return *new(state.Store), nil
}

// check name

func StateItem2SetRequest(grpcReq *dapr_common_v1pb.StateItem, key string) *state.SetRequest {
	_ = "STUB: not implemented"
	// Set the key for the request
	return nil
}

// check if the grpcReq exists

// Assign the value of grpcReq property to req

// Check grpcReq.Etag

// Check grpcReq.Options

func GetResponse2GetStateResponse(compResp *state.GetResponse) *dapr_v1pb.GetStateResponse {
	_ = "STUB: not implemented"
	// Initialize an element of type GetStateResponse
	return nil
}

// check if the compResp exists

func StateConsistencyToString(c dapr_common_v1pb.StateOptions_StateConsistency) string {
	_ = "STUB: not implemented"
	// check
	return ""
}

func StateConcurrencyToString(c dapr_common_v1pb.StateOptions_StateConcurrency) string {
	_ = "STUB: not implemented"
	// check the StateOptions of StateOptions_StateConcurrency
	return ""
}

// wrapDaprComponentError parse and wrap error from dapr component
func (d *daprGrpcAPI) wrapDaprComponentError(err error, format string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// check the Kind of error

func generateGetStateTask(store state.Store, req *state.GetRequest, resultCh chan *dapr_v1pb.BulkStateItem) func() {
	_ = "STUB: not implemented"

	// get
	return nil
}

// convert

// collect result

//never happen

// converting from BulkGetResponse to BulkStateItem
func BulkGetResponse2BulkStateItem(compResp *state.BulkGetResponse) *dapr_v1pb.BulkStateItem {
	_ = "STUB: not implemented"
	return nil
}

// converting from GetResponse to BulkStateItem
func GetResponse2BulkStateItem(compResp *state.GetResponse, key string) *dapr_v1pb.BulkStateItem {
	_ = "STUB: not implemented"
	// convert
	return nil
}

// converting from DeleteStateRequest to DeleteRequest
func DeleteStateRequest2DeleteRequest(grpcReq *dapr_v1pb.DeleteStateRequest, key string) *state.DeleteRequest {
	_ = "STUB: not implemented"
	// convert
	return nil
}

// converting from StateItem to DeleteRequest
func StateItem2DeleteRequest(grpcReq *dapr_common_v1pb.StateItem, key string) *state.DeleteRequest {
	_ = "STUB: not implemented"
	//convert
	return nil
}
