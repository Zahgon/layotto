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

	"mosn.io/layotto/components/lock"
	runtimev1pb "mosn.io/layotto/spec/proto/runtime/v1"
)

func (a *api) TryLock(ctx context.Context, req *runtimev1pb.TryLockRequest) (*runtimev1pb.TryLockResponse, error) {
	_ = "STUB: not implemented"
	// 1. validate
	return nil, nil
}

// 2. find store component

// 3. convert request

// modify key

// 4. delegate to the component

// 5. convert response

func (a *api) Unlock(ctx context.Context, req *runtimev1pb.UnlockRequest) (*runtimev1pb.UnlockResponse, error) {
	_ = "STUB: not implemented"
	// 1. validate
	return nil, nil
}

// 2. find store component

// 3. convert request

// modify key

// 4. delegate to the component

// 5. convert response

func (a *api) LockKeepAlive(ctx context.Context, request *runtimev1pb.LockKeepAliveRequest) (*runtimev1pb.LockKeepAliveResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newInternalErrorUnlockResponse() *runtimev1pb.UnlockResponse {
	_ = "STUB: not implemented"
	return nil
}

func TryLockRequest2ComponentRequest(req *runtimev1pb.TryLockRequest) *lock.TryLockRequest {
	_ = "STUB: not implemented"
	return nil
}

func TryLockResponse2GrpcResponse(compResponse *lock.TryLockResponse) *runtimev1pb.TryLockResponse {
	_ = "STUB: not implemented"
	return nil
}

func UnlockGrpc2ComponentRequest(req *runtimev1pb.UnlockRequest) *lock.UnlockRequest {
	_ = "STUB: not implemented"
	return nil
}

func UnlockComp2GrpcResponse(compResp *lock.UnlockResponse) *runtimev1pb.UnlockResponse {
	_ = "STUB: not implemented"
	return nil
}
