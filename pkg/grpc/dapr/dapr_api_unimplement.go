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

	"google.golang.org/protobuf/types/known/emptypb"

	"mosn.io/layotto/pkg/grpc/dapr/proto/runtime/v1"
)

func (d *daprGrpcAPI) RegisterActorTimer(ctx context.Context, request *runtime.RegisterActorTimerRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *daprGrpcAPI) UnregisterActorTimer(ctx context.Context, request *runtime.UnregisterActorTimerRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *daprGrpcAPI) RegisterActorReminder(ctx context.Context, request *runtime.RegisterActorReminderRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *daprGrpcAPI) UnregisterActorReminder(ctx context.Context, request *runtime.UnregisterActorReminderRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *daprGrpcAPI) GetActorState(ctx context.Context, request *runtime.GetActorStateRequest) (*runtime.GetActorStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *daprGrpcAPI) ExecuteActorStateTransaction(ctx context.Context, request *runtime.ExecuteActorStateTransactionRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *daprGrpcAPI) InvokeActor(ctx context.Context, request *runtime.InvokeActorRequest) (*runtime.InvokeActorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *daprGrpcAPI) GetConfigurationAlpha1(ctx context.Context, request *runtime.GetConfigurationRequest) (*runtime.GetConfigurationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *daprGrpcAPI) SubscribeConfigurationAlpha1(request *runtime.SubscribeConfigurationRequest, server runtime.Dapr_SubscribeConfigurationAlpha1Server) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *daprGrpcAPI) GetMetadata(ctx context.Context, empty *emptypb.Empty) (*runtime.GetMetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *daprGrpcAPI) SetMetadata(ctx context.Context, request *runtime.SetMetadataRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *daprGrpcAPI) Shutdown(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
