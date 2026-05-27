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
	"google.golang.org/grpc"

	"mosn.io/layotto/components/pluggable"
)

const (
	// the default folder to store pluggable component socket files.
	defaultSocketFolder = "/tmp/runtime/component-sockets"

	// SocketFolderEnvVar uses to set the path of folder to store pluggable component socket files replacing defaultSocketFolder
	SocketFolderEnvVar = "LAYOTTO_COMPONENTS_SOCKETS_FOLDER"
)

// GetSocketFolderPath gets the path of folder storing pluggable component socket files.
func GetSocketFolderPath() string { _ = "STUB: not implemented"; return "" }

// Discover discovers pluggable component.
// At present, layotto only support register component from unix domain socket connection,
// and not compatible with windows.
func Discover() ([]pluggable.Component, error) {
	_ = "STUB: not implemented"
	// 1. discover pluggable component
	return nil, nil
}

// 2. callback to register factory into MosnRuntime

// get service form socket files.
type reflectServiceClient interface {
	ListServices() ([]string, error)
	Reset()
}

type grpcService struct {
	// protoRef is the proto service name
	protoRef string
	// componentName is the component name that implements such service.
	componentName string
	// dialer is the used grpc connectiondialer.
	dialer pluggable.GRPCConnectionDialer
}

type grpcConnectionCloser interface {
	grpc.ClientConnInterface
	Close() error
}

// discover use grpc reflect to get services' information.
func discover() ([]grpcService, error) {
	_ = "STUB: not implemented"
	// set grpc connection timeout to prevent block
	return nil, nil
}

// reflectServiceConnectionCloser is used for cleanup the stream created to be used for the reflection service.
func reflectServiceConnectionCloser(conn grpcConnectionCloser, client reflectServiceClient) func() {
	_ = "STUB: not implemented"
	return nil
}

// discover service socket files and get service information from factory factor function.
func serviceDiscovery(reflectClientFactory func(socket string) (client reflectServiceClient, cleanup func(), err error)) ([]grpcService, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// 1. get socket folder files
}

// 2. read socket files

// skip dirs

// check is socket files

// 3. using reflectClientFactory gets service information

// callback use callback function to register pluggable component factories into MosnRuntime
func callback(services []grpcService) []pluggable.Component { _ = "STUB: not implemented"; return nil }
