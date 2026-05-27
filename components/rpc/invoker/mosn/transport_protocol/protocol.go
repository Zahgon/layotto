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

package transport_protocol

import (
	"mosn.io/api"

	"mosn.io/layotto/components/rpc"
)

// protocolRegistry is storage protocol
var protocolRegistry = map[string]TransportProtocol{}

// transport protocol support by mosn(bolt/boltv2...)
type TransportProtocol interface {
	Init(map[string]interface{}) error
	api.Encoder
	api.Decoder
	ToFrame(*rpc.RPCRequest) api.XFrame
	FromFrame(api.XRespFrame) (*rpc.RPCResponse, error)
}

// GetProtocol is get TransportProtocol
func GetProtocol(protocol string) TransportProtocol {
	_ = "STUB: not implemented"
	return *new(TransportProtocol)
}

// RegistProtocol is regist protocol
func RegistProtocol(protocol string, proto TransportProtocol) { _ = "STUB: not implemented"; return }

type fromFrame struct{}

// FromFrame is XRespFrame transform RPCResponse
func (f *fromFrame) FromFrame(resp api.XRespFrame) (*rpc.RPCResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
