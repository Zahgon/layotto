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

	"google.golang.org/protobuf/types/known/emptypb"

	runtimev1pb "mosn.io/layotto/spec/proto/runtime/v1"
)

// GetConfiguration gets configuration from configuration store.
func (a *api) GetConfiguration(ctx context.Context, req *runtimev1pb.GetConfigurationRequest) (*runtimev1pb.GetConfigurationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check store type supported or not

//here protect user use space for sting, eg: " ", "de fault"

// SaveConfiguration saves configuration into configuration store.
func (a *api) SaveConfiguration(ctx context.Context, req *runtimev1pb.SaveConfigurationRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteConfiguration deletes configuration from configuration store.
func (a *api) DeleteConfiguration(ctx context.Context, req *runtimev1pb.DeleteConfigurationRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SubscribeConfiguration gets configuration from configuration store and subscribe the updates.
func (a *api) SubscribeConfiguration(sub runtimev1pb.Runtime_SubscribeConfigurationServer) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO currently this goroutine model is error-prone,and it should be refactored after new version of configuration API being accepted
// 1. start a reader goroutine

// 1.1. read stream

// 1.2. if an error happens,stop all the subscribers

// stop all the subscribers

// TODO this method will stop subscribers created by other connections.Should be refactored

// stop writer goroutine

// 1.3. else find the component and delegate to it

// 1.3.1. stop if StoreName is not supported

// stop all the subscribers

// stop writer goroutine

// 1.3.2. use default settings if blank

// 1.3.3. delegate to the component

// 2. start a writer goroutine

// read response from components

// write to response stream

//	read exit signal
