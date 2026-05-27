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

	"mosn.io/layotto/pkg/grpc/dapr/proto/runtime/v1"
)

func (d *daprGrpcAPI) GetSecret(ctx context.Context, request *runtime.GetSecretRequest) (*runtime.GetSecretResponse, error) {
	_ = "STUB: not implemented"
	// 1. check parameters
	return nil, nil
}

// 2. TODO permission control

// 3. delegate to components

// 4. parse result

func (d *daprGrpcAPI) GetBulkSecret(ctx context.Context, in *runtime.GetBulkSecretRequest) (*runtime.GetBulkSecretResponse, error) {
	_ = "STUB: not implemented"
	// 1. check parameters
	return nil, nil
}

// 2. delegate to components

// 3. parse result

// 4. filter result

// TODO: permission control
