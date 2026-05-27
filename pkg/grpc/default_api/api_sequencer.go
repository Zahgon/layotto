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

	"mosn.io/layotto/components/sequencer"
	runtimev1pb "mosn.io/layotto/spec/proto/runtime/v1"
)

func (a *api) GetNextId(ctx context.Context, req *runtimev1pb.GetNextIdRequest) (*runtimev1pb.GetNextIdResponse, error) {
	_ = "STUB: not implemented"
	// 1. validate
	return nil, nil
}

// 2. convert

// modify key

// 3. find store component

// 4. invoke component

// WEAK

// STRONG

// 5. convert response

func (a *api) getNextIdWithWeakAutoIncrement(ctx context.Context, store sequencer.Store, compReq *sequencer.GetNextIdRequest) (int64, error) {
	_ = "STUB: not implemented"
	// 1. try to get from cache
	return 0, nil
}

// 2. get from component

func (a *api) getNextIdFromComponent(ctx context.Context, store sequencer.Store, compReq *sequencer.GetNextIdRequest) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetNextIdRequest2ComponentRequest(req *runtimev1pb.GetNextIdRequest) (*sequencer.GetNextIdRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
